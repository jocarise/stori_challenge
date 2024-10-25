package services

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"

	"newsletter-service/pkg/files"
	"newsletter-service/pkg/utils"
	"newsletter-service/src/models"
	"newsletter-service/src/repositories"
)

type NewsletterService struct {
	repo repositories.NewsletterRepository
}

func NewNewsletterService(repo repositories.NewsletterRepository) *NewsletterService {
	return &NewsletterService{repo: repo}
}

func (s *NewsletterService) GetCategories() ([]models.Category, error) {
	return s.repo.GetCategories()
}

func (s *NewsletterService) AddNewsletterWithFile(dto *models.NewsletterDTO, fileData io.Reader, header *multipart.FileHeader) (*models.Newsletter, error) {
	if !files.IsValidFileType(header) {
		return nil, errors.New("invalid file type")
	}

	filePath, err := files.SaveFile(fileData, header, uuid.New().String(), os.Getenv("FILES_PATH"))
	if err != nil {
		log.Printf("error saving file: %v", err)
		return nil, errors.New("invalid file type")
	}

	if err := dto.Validate(); err != nil {
		return nil, err
	}

	category, err := s.repo.FindCategory(dto.CategoryId)
	if err != nil {
		log.Printf("error getting category: %v", err)
		return nil, errors.New("invalid category")
	}

	newsletter := models.Newsletter{
		ID:            dto.ID,
		Title:         dto.Title,
		Html:          bluemonday.UGCPolicy().Sanitize(dto.Html),
		Attachment:    filePath,
		CategoryID:    &category.ID,
		ScheduledDate: dto.ScheduledDate,
		Scheduled:     dto.Scheduled,
	}

	if err := s.repo.SaveNewsletter(&newsletter); err != nil {
		return nil, err
	}

	baseURL := os.Getenv("BASE_URL") + os.Getenv("API_VERSION")

	emails := dto.Recipients

	var existingRecipients []models.Recipient
	if err := s.repo.FindRecipientsByEmails(emails, &existingRecipients); err != nil {
		return nil, err
	}

	existingRecipientMap := make(map[string]models.Recipient)
	for _, rec := range existingRecipients {
		existingRecipientMap[rec.Email] = rec
	}

	for _, email := range emails {
		if existingRecipient, found := existingRecipientMap[email]; found {
			if err := s.repo.Associate(&newsletter, &existingRecipient); err != nil {
				return nil, err
			}
		} else {
			id := uuid.New().String()
			newRecipient := models.Recipient{
				ID:            id,
				Email:         email,
				UnsuscribeUrl: utils.GenerateUnsubscribeURLFromRequest(baseURL, id),
			}
			if err := s.repo.SaveRecipient(&newRecipient); err != nil {
				return nil, err
			}

			if err := s.repo.Associate(&newsletter, &newRecipient); err != nil {
				return nil, err
			}
		}
	}

	return &newsletter, nil
}

func (s *NewsletterService) GetNewsletters() ([]models.Newsletter, error) {
	return s.repo.GetNewsletters()
}

func (s *NewsletterService) GetNewsletterById(id string) (*models.Newsletter, error) {
	return s.repo.GetNewsletterByID(id)
}

func (s *NewsletterService) UnsubscribeFromNewsletters(recipientId string) error {
	return s.repo.RemoveRecipientFromNewsletters(recipientId)
}

func (s *NewsletterService) UnsubscribeFromSpecificCategory(categoryId uint, recipientId string) error {
	return s.repo.RemoveRecipientFromNewsletterByCategory(categoryId, recipientId)
}

func (s *NewsletterService) CreateRecipient(newsletterId, email string) (*models.Recipient, error) {
	baseURL := os.Getenv("BASE_URL") + os.Getenv("API_VERSION")

	if !utils.ValidateEmailRequirements(email) {
		return nil, errors.New("invalid email format")
	}

	existingRecipient, err, exist := s.repo.FindRecipientByEmail(email)
	if err != nil {
		log.Printf("Error checking existing recipient: %v", err)
	}

	newsletter, err := s.repo.GetNewsletterByID(newsletterId)
	if err != nil {
		log.Printf("Error getting newsletter: %s %v", newsletterId, err)
		return nil, errors.New("internal server error")
	}

	if exist {
		err = s.repo.Associate(newsletter, existingRecipient)
		if err != nil {
			log.Printf("Error associating recipient with newsletter")
			return nil, errors.New("internal server error")
		}

		return existingRecipient, nil
	}

	if !exist {
		id := uuid.New().String()
		newRecipient := models.Recipient{
			ID:            id,
			Email:         email,
			UnsuscribeUrl: utils.GenerateUnsubscribeURLFromRequest(baseURL, id),
		}

		err = s.repo.SaveRecipient(&newRecipient)
		if err != nil {
			log.Printf("Error saving recipient: %s %v", id, err)
			return nil, errors.New("internal server error")
		}

		err = s.repo.Associate(newsletter, &newRecipient)
		if err != nil {
			log.Printf("Error associating recipient: %s with newsletter: %s  %v", id, newsletter.ID, err)
			return nil, errors.New("internal server error")
		}

		return &newRecipient, nil
	}

	return nil, errors.New("internal server error")
}

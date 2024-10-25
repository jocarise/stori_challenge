package repositories

import (
	"errors"
	"log"
	"newsletter-service/src/models"

	"gorm.io/gorm"
)

type NewsletterRepository interface {
	GetCategories() ([]models.Category, error)
	GetNewsletters() ([]models.Newsletter, error)
	GetNewsletterByID(string) (*models.Newsletter, error)
	FindRecipientByEmail(string) (*models.Recipient, error, bool)
	SaveNewsletter(*models.Newsletter) error
	SaveRecipient(*models.Recipient) error
	Associate(*models.Newsletter, *models.Recipient) error
	FindCategory(uint) (*models.Category, error)
	FindRecipientsByEmails([]string, *[]models.Recipient) error
	RemoveRecipientFromNewsletters(string) error
	RemoveRecipientFromNewsletterByCategory(uint, string) error
}

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	return &GORMRepository{db: db}
}

func (r *GORMRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GORMRepository) SaveNewsletter(newsletter *models.Newsletter) error {
	return r.db.Create(newsletter).Error
}

func (r *GORMRepository) SaveRecipient(rec *models.Recipient) error {
	return r.db.Create(rec).Error
}

func (r *GORMRepository) Associate(newsletter *models.Newsletter, rec *models.Recipient) error {
	return r.db.Model(newsletter).Association("Recipients").Append(rec)
}

func (r *GORMRepository) FindCategory(id uint) (*models.Category, error) {
	var category models.Category

	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, nil
}

func (r *GORMRepository) FindRecipientsByEmails(emails []string, recipients *[]models.Recipient) error {
	return r.db.Where("email IN ?", emails).Find(recipients).Error
}

func (r *GORMRepository) GetNewsletters() ([]models.Newsletter, error) {
	var newsletters []models.Newsletter

	err := r.db.Preload("Category").
		Preload("Recipients").
		Order("created_at DESC").
		Find(&newsletters).Error

	if err != nil {
		return nil, err
	}

	return newsletters, nil
}

func (r *GORMRepository) GetNewsletterByID(id string) (*models.Newsletter, error) {
	var newsletter models.Newsletter

	err := r.db.Preload("Category").
		Preload("Recipients").
		First(&newsletter, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &newsletter, nil
}

func (r *GORMRepository) RemoveRecipientFromNewsletters(recipientId string) error {
	var recipient models.Recipient

	if err := r.db.First(&recipient, "id = ?", recipientId).Error; err != nil {
		return err
	}

	if err := r.db.Model(&recipient).Association("Newsletters").Clear(); err != nil {
		return err
	}

	log.Printf("all newsletters removed for recipient %v", recipientId)
	return nil
}

func (r *GORMRepository) RemoveRecipientFromNewsletterByCategory(categoryId uint, recipientId string) error {
	var recipient models.Recipient

	if err := r.db.First(&recipient, "id = ?", recipientId).Error; err != nil {
		return err
	}

	var newsletters []models.Newsletter
	if err := r.db.Model(&recipient).Association("Newsletters").Find(&newsletters); err != nil {
		return err
	}

	for _, newsletter := range newsletters {
		if newsletter.CategoryID != nil && *newsletter.CategoryID == categoryId {
			if err := r.db.Model(&newsletter).Association("Recipients").Delete(&recipient); err != nil {
				return err
			}
			log.Printf("removed recipient %v from newsletter %v in category %v", recipientId, newsletter.ID, categoryId)
		}
	}

	return nil
}

func (r *GORMRepository) FindRecipientByEmail(email string) (*models.Recipient, error, bool) {
	var recipient models.Recipient

	if err := r.db.Where("email = ?", email).First(&recipient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, false
		}
		return nil, nil, false
	}

	return &recipient, nil, true
}

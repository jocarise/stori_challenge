package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"newsletter-service/pkg/email"
	"newsletter-service/pkg/utils"
	"newsletter-service/src/models"
	"newsletter-service/src/services"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type NewsletterHandler struct {
	service *services.NewsletterService
}

func NewNewsletterHandler(service *services.NewsletterService) *NewsletterHandler {
	return &NewsletterHandler{service: service}
}

func (h *NewsletterHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	categories, err := h.service.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Categories []models.Category `json:"categories"`
	}{
		Categories: categories,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *NewsletterHandler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseMultipartForm(5 << 20) // Limit to 5 MB
	if err != nil {
		log.Printf("error parsing form: %v", err)
		http.Error(w, "unable to process file", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("unable to save the file: %v", err)
		http.Error(w, "unable to process file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	newsletterDTO := &models.NewsletterDTO{}

	newsletterDTO.ID = r.FormValue("id")
	newsletterDTO.Title = r.FormValue("title")
	newsletterDTO.Html = r.FormValue("html")

	newsletterDTO.Recipients = r.Form["recipients[]"]

	categoryId, err := utils.ParseStringToUInt(r.FormValue("categoryId"))
	if err != nil {
		http.Error(w, "invalid category id format", http.StatusBadRequest)
		return
	}
	newsletterDTO.CategoryId = categoryId

	newsletterDTO.Scheduled = r.FormValue("scheduled") == "true"
	if newsletterDTO.Scheduled {
		scheduledDate, err := utils.ParseScheduledDate(r.FormValue("scheduledDate"))
		if err != nil {
			http.Error(w, "invalid date format", http.StatusBadRequest)
			return
		}

		if !time.Now().After(scheduledDate) {
			newsletterDTO.ScheduledDate = &scheduledDate
		} else {
			http.Error(w, "invalid date", http.StatusBadRequest)
			return
		}
	} else {
		newsletterDTO.ScheduledDate = nil
	}

	newsletter, err := h.service.AddNewsletterWithFile(newsletterDTO, file, header)
	if err != nil {
		log.Printf("Error saving Newsletter or file: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newsletter); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *NewsletterHandler) GetNewsletters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newsletters, err := h.service.GetNewsletters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Newsletters []models.Newsletter `json:"newsletters"`
	}{
		Newsletters: newsletters,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *NewsletterHandler) SendNewsletters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	smtpPort, err := strconv.Atoi(os.Getenv("GMAIL_PORT"))
	if err != nil {
		log.Printf("Invalid port number: %v", err)
		http.Error(w, "Invalid smtpPort config:", http.StatusInternalServerError)
		return
	}

	server := gomail.NewDialer(os.Getenv("GMAIL_HOST"), smtpPort, os.Getenv("GMAIL_EMAIL"), os.Getenv("GMAIL_KEY"))
	if _, err := server.Dial(); err != nil {
		http.Error(w, "Failed to connect to the email server", http.StatusInternalServerError)
		return
	}

	id := chi.URLParam(r, "newsletterId")
	newsletter, err := h.service.GetNewsletterById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newsletter.Scheduled { //Early return
		log.Printf("newsletter is already scheduled: %v", newsletter.ID)
		w.WriteHeader(http.StatusConflict)
		return
	}

	var wg sync.WaitGroup
	for _, recipient := range newsletter.Recipients {
		wg.Add(1)
		go func(recipient models.Recipient) {
			defer wg.Done()
			if err := email.SendEmail(
				server,
				newsletter.Title,
				newsletter.Attachment,
				newsletter.Html,
				recipient.Email,
				recipient.UnsuscribeUrl,
				utils.GenerateUnsubscribeURLByCategoryFromRequest(recipient.UnsuscribeUrl, newsletter.Category.ID),
			); err != nil {
				log.Printf("Failed to send email to %s: %v", recipient.Email, err)
			}
		}(recipient)
	}
	wg.Wait()

	w.WriteHeader(http.StatusOK)
}

func (h *NewsletterHandler) UnsuscribeRecipient(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	rId := query.Get("recipientId")
	cId := query.Get("categoryId")

	if rId == "" {
		http.Error(w, "recipientId is required", http.StatusBadRequest)
		return
	}

	if err := uuid.Validate(rId); err != nil {
		http.Error(w, "invalid recipientId", http.StatusBadRequest)
		return
	}

	if cId == "" {
		if err := h.service.UnsubscribeFromNewsletters(rId); err != nil {
			log.Printf("error unsubscribing recipient: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Print("executed unsubscribe for recipient:", rId)
	} else {
		categoryId, err := utils.ParseStringToUInt(cId)
		if err != nil {
			http.Error(w, "invalid category id format", http.StatusBadRequest)
			return
		}

		err = h.service.UnsubscribeFromSpecificCategory(categoryId, rId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("executed unsubscribe from newsletter: categoryId=%d, recipientId=%s", categoryId, rId)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully unsubscribed"))
}

func (h *NewsletterHandler) AddRecipient(w http.ResponseWriter, r *http.Request) {
	newsletterId := chi.URLParam(r, "newsletterId")

	var body struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	recipient, err := h.service.CreateRecipient(newsletterId, body.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(recipient); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

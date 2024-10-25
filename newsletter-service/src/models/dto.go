package models

import (
	"errors"
	"newsletter-service/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type NewsletterDTO struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Html          string     `json:"html"`
	Recipients    []string   `json:"recipients"`
	CategoryId    uint       `json:"categoryId"`
	Scheduled     bool       `json:"scheduled"`
	ScheduledDate *time.Time `json:"scheduledDate"`
}

func (n *NewsletterDTO) Validate() error {
	if n.ID == "" || uuid.Validate(n.ID) != nil {
		return errors.New("uuid error")
	}

	if len(n.Recipients) == 0 {
		return errors.New("recipients cannot be empty")
	}

	emailMap := make(map[string]struct{})

	for _, recipient := range n.Recipients {
		if !utils.ValidateEmailRequirements(recipient) {
			return errors.New("invalid email format: " + recipient)
		}

		if _, exists := emailMap[recipient]; exists {
			return errors.New("duplicate email found: " + recipient)
		}
		emailMap[recipient] = struct{}{}
	}

	if n.Scheduled && n.ScheduledDate.IsZero() {
		return errors.New("ScheduledDate must be set when IsScheduled is true")
	}

	return nil
}

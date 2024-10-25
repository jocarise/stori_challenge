package utils

import (
	"fmt"
	"newsletter-service/pkg/regex"
	"regexp"
	"strconv"
	"time"
)

func ValidateEmailRequirements(email string) bool {
	a := regexp.MustCompile(regex.EmailRegex)
	if !a.MatchString(email) {
		return false
	}

	b := regexp.MustCompile(`\.\.+`)

	return !b.MatchString(email)
}

func GenerateUnsubscribeURLFromRequest(baseURL, recipientId string) string {
	return fmt.Sprintf("%s/newsletters/unsuscribe?recipientId=%s", baseURL, recipientId)
}

func GenerateUnsubscribeURLByCategoryFromRequest(url string, categoryId uint) string {
	return fmt.Sprintf(url+"&categoryId=%v", categoryId)
}

func ParseStringToUInt(input string) (uint, error) {
	value, err := strconv.ParseUint(input, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func ParseScheduledDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, fmt.Errorf("scheduled date string is empty")
	}

	scheduledDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid scheduled date format: %w", err)
	}

	return scheduledDate, nil
}

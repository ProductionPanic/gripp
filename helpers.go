package gripp

import (
	"errors"
	"time"
)

func GetToday() string {
	now := time.Now()
	return GetDateInFormat(now)
}

func EnsureDateFormat(date string) (string, error) {
	// List the formats you want to support
	formats := []string{"02-01-2006", "2006-01-02"}

	for _, f := range formats {
		t, err := time.Parse(f, date)
		if err == nil {
			// Once parsed successfully, return it in your target format
			return t.Format("2006-01-02"), nil
		}
	}

	return "", errors.New("invalid date: must be DD-MM-YYYY or YYYY-MM-DD")
}

func GetDateInFormat(date time.Time) string {
	return date.Format("2006-01-02")
}

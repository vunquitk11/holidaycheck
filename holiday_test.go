package holidaycheck

import (
	"testing"
	"time"
)

func Test_IsTodayPublicHoliday(t *testing.T) {
	// Define test cases
	testCases := []struct {
		countryCode string
		date        string
		expected    bool
		name        string
	}{
		{"SG", "2025-01-01", true, "New Year's Day"},
		{"SG", "2025-01-29", true, "Chinese New Year"},
		{"SG", "2025-01-30", true, "Chinese New Year"},
		{"SG", "2025-04-18", true, "Good Friday"},
		{"SG", "2025-05-01", true, "Labour Day"},
		{"SG", "2025-03-31", true, "Hari Raya Puasa"},
		{"SG", "2025-05-12", true, "Vesak Day"},
		{"SG", "2025-06-07", true, "Hari Raya Haji"},
		{"SG", "2025-08-09", true, "National Day"},
		{"SG", "2025-10-29", true, "Deepavali"},
		{"SG", "2025-12-25", true, "Christmas Day"},
		{"SG", "2025-07-04", false, "Not a holiday"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			date, _ := time.Parse("2006-01-02", tc.date)
			isHoliday, err := IsTodayPublicHoliday(tc.countryCode, date)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if isHoliday != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, isHoliday)
			}
		})
	}
}

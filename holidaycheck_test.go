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
		// --- Singapore 2025 ---
		{"SG", "2025-01-01", true, "New Year's Day"},
		{"SG", "2025-01-29", true, "Chinese New Year"},
		{"SG", "2025-01-30", true, "Chinese New Year"},
		{"SG", "2025-04-18", true, "Good Friday"},
		{"SG", "2025-05-01", true, "Labour Day"},
		{"SG", "2025-03-31", true, "Hari Raya Puasa"},
		{"SG", "2025-05-12", true, "Vesak Day"},
		{"SG", "2025-06-06", true, "Hari Raya Haji"},
		{"SG", "2025-08-09", true, "National Day"},
		{"SG", "2025-10-29", true, "Deepavali"},
		{"SG", "2025-12-25", true, "Christmas Day"},
		{"SG", "2025-07-04", false, "Not a holiday"},
		// --- Singapore 2026 ---
		{"SG", "2026-01-01", true, "New Year's Day"},
		{"SG", "2026-02-17", true, "Chinese New Year"},
		{"SG", "2026-02-18", true, "Chinese New Year Holiday"},
		{"SG", "2026-03-20", true, "Hari Raya Puasa"},
		{"SG", "2026-04-03", true, "Good Friday"},
		{"SG", "2026-05-01", true, "Labour Day"},
		{"SG", "2026-05-27", true, "Hari Raya Haji"},
		{"SG", "2026-05-31", true, "Vesak Day"},
		{"SG", "2026-06-01", true, "Vesak Day Holiday (observed)"},
		{"SG", "2026-08-09", true, "National Day"},
		{"SG", "2026-08-10", true, "National Day Holiday (observed)"},
		{"SG", "2026-11-08", true, "Deepavali"},
		{"SG", "2026-11-09", true, "Deepavali Holiday (observed)"},
		{"SG", "2026-12-25", true, "Christmas Day"},
		// --- Vietnam 2025 ---
		{"VN", "2025-01-01", true, "New Year's Day"},
		{"VN", "2025-01-29", true, "Lunar New Year's Eve"},
		{"VN", "2025-01-30", true, "Lunar New Year"},
		{"VN", "2025-01-31", true, "Tet Holiday"},
		{"VN", "2025-02-01", true, "Tet Holiday"},
		{"VN", "2025-02-02", true, "Tet Holiday"},
		{"VN", "2025-02-03", true, "Tet Holiday"},
		{"VN", "2025-04-11", true, "Hung Kings Commemoration Day"},
		{"VN", "2025-04-30", true, "Reunification Day"},
		{"VN", "2025-05-01", true, "International Workers' Day"},
		{"VN", "2025-09-02", true, "National Day"},
		{"VN", "2025-06-01", false, "Not a holiday"},
		// --- Vietnam 2026 ---
		{"VN", "2026-01-01", true, "New Year's Day"},
		{"VN", "2026-02-16", true, "Lunar New Year's Eve"},
		{"VN", "2026-02-17", true, "Lunar New Year"},
		{"VN", "2026-02-18", true, "Tet Holiday"},
		{"VN", "2026-02-19", true, "Tet Holiday"},
		{"VN", "2026-02-20", true, "Tet Holiday"},
		{"VN", "2026-02-21", true, "Tet Holiday"},
		{"VN", "2026-03-29", true, "Hung Kings Commemoration Day"},
		{"VN", "2026-04-30", true, "Reunification Day"},
		{"VN", "2026-05-01", true, "International Workers' Day"},
		{"VN", "2026-09-02", true, "National Day"},
		{"VN", "2026-03-01", false, "Not a holiday"},
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

func Test_GetHolidayName(t *testing.T) {
	// Define test cases
	testCases := []struct {
		countryCode string
		date        string
		expected    string
		name        string
		expectError bool
	}{
		// Singapore test cases
		{"SG", "2025-01-01", "New Year's Day", "SG New Year's Day", false},
		{"SG", "2025-08-09", "National Day", "SG National Day", false},
		// Vietnam test cases
		{"VN", "2025-01-01", "New Year's Day", "VN New Year's Day", false},
		{"VN", "2025-09-02", "National Day", "VN National Day", false},
		// Indonesia test cases
		{"ID", "2025-01-01", "New Year's Day", "ID New Year's Day", false},
		{"ID", "2025-08-17", "Independence Day", "ID Independence Day", false},
		// Malaysia test cases
		{"MY", "2025-01-01", "New Year's Day", "MY New Year's Day", false},
		{"MY", "2025-08-31", "National Day (Merdeka)", "MY National Day", false},
		// Error test cases
		{"INVALID", "2025-01-01", "", "Invalid country code", true},
		// Empty result test cases
		{"SG", "2025-07-04", "", "SG Not a holiday", false},
		{"VN", "2025-06-15", "", "VN Not a holiday", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			date, _ := time.Parse("2006-01-02", tc.date)
			holidayName, err := GetHolidayName(tc.countryCode, date)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if holidayName != tc.expected {
					t.Errorf("expected %v, got %v", tc.expected, holidayName)
				}
			}
		})
	}
}

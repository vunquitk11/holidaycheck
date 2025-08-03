package main

import (
	"fmt"
	"time"

	"github.com/vunquitk11/holidaycheck/data"
	"github.com/vunquitk11/holidaycheck/models"
)

// IsTodayPublicHoliday checks if the given date is a public holiday
// for the specified country code.
func IsTodayPublicHoliday(countryCode string, todayUTC time.Time) (bool, error) {

	enumCountryCode := models.CountryCode(countryCode)
	// Check if the country code is valid
	if !enumCountryCode.IsValid() {
		return false, fmt.Errorf("invalid country code: %v", countryCode)
	}

	// Convert UTC to local time and get year and month-day
	year, monthDate, err := enumCountryCode.ToLocalYearAndDate(todayUTC)
	if err != nil {
		return false, fmt.Errorf("error converting to local time: %v", err)
	}

	// Check if the holiday exists
	if holidayName, exists := data.Holidays[year][enumCountryCode][monthDate]; exists {
		fmt.Println("holidayName: ", holidayName)
		return true, nil
	}
	return false, nil
}

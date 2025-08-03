package data

import (
	"github.com/vunquitk11/holidaycheck/models"
)

// Holidays define the total map for holidays across different years
var Holidays = map[int]map[models.CountryCode]map[string]string{
	2025: holidays2025,
}

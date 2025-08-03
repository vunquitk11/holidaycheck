package data

import (
	"github.com/vunquitk11/holidaycheck/models"
)

// Holidays define the total map for holidays across different years
var Holidays = map[int]map[models.CountryCode]map[string]string{
	2025: holidays2025,
	2026: holidays2026,
	2027: holidays2027,
	2028: holidays2028,
}

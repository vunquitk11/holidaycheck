package models

import (
	"time"
)

const (
	monthDateTimeFormat = "01-02"
)

// CountryCode define enum for country codes
type CountryCode string

const (
	SG CountryCode = "SG"
	VN CountryCode = "VN"
	ID CountryCode = "ID"
	MY CountryCode = "MY"
)

// IsValid checks if the CountryCode is valid
func (cc CountryCode) IsValid() bool {
	switch cc {
	case SG, VN, ID, MY:
		return true
	default:
		return false
	}
}

// TimeZoneStr returns the LocalTimeZone enum corresponding as string to the CountryCode
func (cc CountryCode) TimeZoneStr() string {
	switch cc {
	case SG:
		return AsiaSingapore.String()
	case VN:
		return AsiaHoChiMinh.String()
	case ID:
		return AsiaJakarta.String()
	case MY:
		return AsiaKualaLumpur.String()
	default:
		return ""
	}
}

// LocalTimeZone define enum for local time zones
type LocalTimeZone string

const (
	AsiaSingapore   LocalTimeZone = "Asia/Singapore"
	AsiaHoChiMinh   LocalTimeZone = "Asia/Ho_Chi_Minh"
	AsiaJakarta     LocalTimeZone = "Asia/Jakarta"
	AsiaKualaLumpur LocalTimeZone = "Asia/Kuala_Lumpur"
)

// String returns the string representation of the LocalTimeZone
func (ltz LocalTimeZone) String() string {
	switch ltz {
	case AsiaSingapore:
		return "Asia/Singapore"
	case AsiaHoChiMinh:
		return "Asia/Ho_Chi_Minh"
	case AsiaJakarta:
		return "Asia/Jakarta"
	case AsiaKualaLumpur:
		return "Asia/Kuala_Lumpur"
	default:
		return "Unknown"
	}
}

// ToLocalYearAndDate converts a UTC time to the local time of the given country code
// and returns the year as an int, the month-day as a string 'MM-DD', and an error if any.
func (cc CountryCode) ToLocalYearAndDate(utcTime time.Time) (int, string, error) {
	loc, err := time.LoadLocation(cc.TimeZoneStr())
	if err != nil {
		return 0, "", err
	}
	localTime := utcTime.In(loc)
	return localTime.Year(), localTime.Format(monthDateTimeFormat), nil
}

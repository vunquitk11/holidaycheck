package data

import (
	"github.com/vunquitk11/holidaycheck/models"
)

// holidays2025 define the map for holidays in 2025
var holidays2025 = map[models.CountryCode]map[string]string{
	models.ID: {
		"01-01": "New Year's Day",
		"01-29": "Chinese New Year",
		"02-27": "Isra Mi'raj",
		"03-31": "Hari Raya Idul Fitri",
		"04-01": "Hari Raya Idul Fitri Holiday",
		"05-01": "Labour Day",
		"05-12": "Vesak Day",
		"05-29": "Ascension of Jesus Christ",
		"06-06": "Ascension of Prophet Muhammad",
		"08-17": "Independence Day",
		"10-05": "Prophet’s Birthday (Maulid)",
		"12-25": "Christmas Day",
	},
	models.MY: {
		"01-01": "New Year's Day",
		"01-29": "Chinese New Year",
		"01-30": "Chinese New Year Holiday",
		"02-01": "Federal Territory Day",
		"03-31": "Hari Raya Aidilfitri",
		"04-01": "Hari Raya Aidilfitri Holiday",
		"05-01": "Labour Day",
		"05-12": "Wesak Day",
		"06-02": "Agong’s Birthday",
		"06-06": "Hari Raya Haji",
		"06-27": "Islamic New Year",
		"08-31": "National Day (Merdeka)",
		"09-16": "Malaysia Day",
		"10-05": "Prophet’s Birthday (Maulidur Rasul)",
		"10-20": "Deepavali",
		"12-25": "Christmas Day",
	},
	models.SG: {
		"01-01": "New Year's Day",
		"01-29": "Chinese New Year",
		"01-30": "Chinese New Year",
		"03-31": "Hari Raya Puasa",
		"04-18": "Good Friday",
		"05-01": "Labour Day",
		"05-12": "Vesak Day",
		"06-07": "Hari Raya Haji",
		"08-09": "National Day",
		"10-29": "Deepavali",
		"12-25": "Christmas Day",
	},
	models.VN: {
		"01-01": "New Year's Day",
		"01-29": "Lunar New Year's Eve",
		"01-30": "Lunar New Year",
		"01-31": "Tet Holiday",
		"02-01": "Tet Holiday",
		"02-02": "Tet Holiday",
		"02-03": "Tet Holiday",
		"04-11": "Hung Kings Commemoration Day",
		"04-30": "Reunification Day",
		"05-01": "International Workers' Day",
		"09-02": "National Day",
	},
}

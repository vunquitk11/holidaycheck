package data

import "github.com/vunquitk11/holidaycheck/models"

// holidays2025 contains public holidays for Singapore, Vietnam, Indonesia, Malaysia in 2025.
// Movable holidays are based on official and projected calendars, may be subject to change.
var holidays2025 = map[models.CountryCode]map[string]string{
	models.SG: {
		"01-01": "New Year's Day",           // ✅ Fixed
		"01-29": "Chinese New Year",         // 📅 Movable (Lunar)
		"01-30": "Chinese New Year Holiday", // 📅 Movable (Lunar)
		"03-31": "Hari Raya Puasa",          // 📅 Movable (Islamic)
		"04-18": "Good Friday",              // 📅 Movable (Christian)
		"05-01": "Labour Day",               // ✅ Fixed
		"05-12": "Vesak Day",                // 📅 Movable (Buddhist)
		"06-06": "Hari Raya Haji",           // 📅 Movable (Islamic)
		"08-09": "National Day",             // ✅ Fixed
		"10-29": "Deepavali",                // 📅 Movable (Hindu)
		"12-25": "Christmas Day",            // ✅ Fixed
	},
	models.VN: {
		"01-01": "New Year's Day",               // ✅ Fixed
		"01-29": "Lunar New Year's Eve",         // 📅 Movable (Lunar)
		"01-30": "Lunar New Year",               // 📅 Movable (Lunar)
		"01-31": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-01": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-02": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-03": "Tet Holiday",                  // 📅 Movable (Lunar)
		"04-11": "Hung Kings Commemoration Day", // 📅 Movable (Lunar)
		"04-30": "Reunification Day",            // ✅ Fixed
		"05-01": "International Workers' Day",   // ✅ Fixed
		"09-02": "National Day",                 // ✅ Fixed
	},
	models.ID: {
		"01-01": "New Year's Day",                // ✅ Fixed
		"01-29": "Chinese New Year",              // 📅 Movable (Lunar)
		"02-27": "Isra Mi'raj",                   // 📅 Movable (Islamic)
		"03-31": "Hari Raya Idul Fitri",          // 📅 Movable (Islamic)
		"04-01": "Hari Raya Idul Fitri Holiday",  // 📅 Movable (Islamic)
		"05-01": "Labour Day",                    // ✅ Fixed
		"05-12": "Vesak Day",                     // 📅 Movable (Buddhist)
		"05-29": "Ascension of Jesus Christ",     // 📅 Movable (Christian)
		"06-06": "Ascension of Prophet Muhammad", // 📅 Movable (Islamic)
		"08-17": "Independence Day",              // ✅ Fixed
		"10-05": "Prophet’s Birthday (Maulid)",   // 📅 Movable (Islamic)
		"12-25": "Christmas Day",                 // ✅ Fixed
	},
	models.MY: {
		"01-01": "New Year's Day",                      // ✅ Fixed
		"01-29": "Chinese New Year",                    // 📅 Movable (Lunar)
		"01-30": "Chinese New Year Holiday",            // 📅 Movable (Lunar)
		"02-01": "Federal Territory Day",               // ✅ Fixed (in select territories)
		"03-31": "Hari Raya Aidilfitri",                // 📅 Movable (Islamic)
		"04-01": "Hari Raya Aidilfitri Holiday",        // 📅 Movable (Islamic)
		"05-01": "Labour Day",                          // ✅ Fixed
		"05-12": "Wesak Day",                           // 📅 Movable (Buddhist)
		"06-02": "Agong’s Birthday",                    // ✅ Fixed (usually 1st Monday of June)
		"06-06": "Hari Raya Haji",                      // 📅 Movable (Islamic) — corrected from 06-07
		"06-27": "Islamic New Year",                    // 📅 Movable (Islamic)
		"08-31": "National Day (Merdeka)",              // ✅ Fixed
		"09-16": "Malaysia Day",                        // ✅ Fixed
		"10-05": "Prophet’s Birthday (Maulidur Rasul)", // 📅 Movable (Islamic)
		"10-20": "Deepavali",                           // 📅 Movable (Hindu)
		"12-25": "Christmas Day",                       // ✅ Fixed
	},
}

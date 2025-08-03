package data

import "github.com/vunquitk11/holidaycheck/models"

// holidays2026 contains public holidays for Singapore, Vietnam, Indonesia, Malaysia in 2026.
// Movable holidays are based on official and projected calendars, may be subject to change.
var holidays2026 = map[models.CountryCode]map[string]string{
	models.SG: {
		"01-01": "New Year's Day",                            // ✅ Fixed
		"02-17": "Chinese New Year",                          // 📅 Movable (Lunar calendar)
		"02-18": "Chinese New Year Holiday (2nd day)",        // 📅 Movable
		"03-20": "Hari Raya Puasa",                           // 📅 Movable (Islamic calendar)
		"04-03": "Good Friday",                               // 📅 Movable (Christian calendar)
		"05-01": "Labour Day",                                // ✅ Fixed
		"05-27": "Hari Raya Haji",                            // 📅 Movable (Islamic calendar)
		"05-31": "Vesak Day (falls on Sunday)",               // 📅 Movable (Buddhist calendar)
		"06-01": "Vesak Day Holiday (observed on Monday)",    // 🔁 Observed for Vesak (05-31)
		"08-09": "National Day (falls on Sunday)",            // ✅ Fixed, 🔁 Observed
		"08-10": "National Day Holiday (observed on Monday)", // 🔁 Observed for 08-09
		"11-08": "Deepavali (falls on Sunday)",               // 📅 Movable (Hindu), 🔁 Observed
		"11-09": "Deepavali Holiday (observed on Monday)",    // 🔁 Observed for 11-08
		"12-25": "Christmas Day",                             // ✅ Fixed
	},
	models.VN: {
		"01-01": "New Year's Day",               // ✅ Fixed
		"02-16": "Lunar New Year's Eve",         // 📅 Movable (Lunar)
		"02-17": "Lunar New Year",               // 📅 Movable (Lunar)
		"02-18": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-19": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-20": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-21": "Tet Holiday",                  // 📅 Movable (Lunar)
		"03-29": "Hung Kings Commemoration Day", // 📅 Movable (Lunar)
		"04-30": "Reunification Day",            // ✅ Fixed
		"05-01": "International Workers' Day",   // ✅ Fixed
		"09-02": "National Day",                 // ✅ Fixed
	},
	models.ID: {
		"01-01": "New Year's Day",                // ✅ Fixed
		"01-26": "Chinese New Year",              // 📅 Movable (Lunar)
		"02-16": "Isra Mi'raj",                   // 📅 Movable (Islamic)
		"02-17": "Hari Raya Idul Fitri",          // 📅 Movable (Islamic)
		"02-18": "Hari Raya Idul Fitri Holiday",  // 📅 Movable (Islamic)
		"05-01": "Labour Day",                    // ✅ Fixed
		"05-14": "Ascension of Jesus Christ",     // 📅 Movable (Christian)
		"05-31": "Waisak Day",                    // 📅 Movable (Buddhist)
		"06-27": "Ascension of Prophet Muhammad", // 📅 Movable (Islamic)
		"08-17": "Independence Day",              // ✅ Fixed
		"10-26": "Prophet’s Birthday (Maulid)",   // 📅 Movable (Islamic)
		"12-25": "Christmas Day",                 // ✅ Fixed
	},
	models.MY: {
		"01-01": "New Year's Day",                                 // ✅ Fixed
		"02-01": "Federal Territory Day",                          // ✅ Fixed (in select territories)
		"02-17": "Chinese New Year and Hari Raya Aidilfitri",      // 📅 Movable (Lunar & Islamic) — overlap
		"02-18": "Chinese New Year Holiday or Aidilfitri Holiday", // 📅 Movable — likely observed for second day
		"05-01": "Labour Day",                                     // ✅ Fixed
		"05-31": "Wesak Day",                                      // 📅 Movable (Buddhist)
		"06-06": "Agong’s Birthday",                               // ✅ Fixed (usually first Monday of June)
		"06-27": "Hari Raya Haji",                                 // 📅 Movable (Islamic)
		"07-27": "Islamic New Year",                               // 📅 Movable (Islamic)
		"08-31": "National Day (Merdeka)",                         // ✅ Fixed
		"09-16": "Malaysia Day",                                   // ✅ Fixed
		"10-26": "Prophet’s Birthday (Maulidur Rasul)",            // 📅 Movable (Islamic)
		"11-06": "Deepavali",                                      // 📅 Movable (Hindu)
		"12-25": "Christmas Day",                                  // ✅ Fixed
	},
}

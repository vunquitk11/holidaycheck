package data

import "github.com/vunquitk11/holidaycheck/models"

// holidays2028 contains public holidays for Singapore, Vietnam, Indonesia, Malaysia in 2028.
// Movable holidays are based on official and projected calendars, may be subject to change.
var holidays2028 = map[models.CountryCode]map[string]string{
	models.SG: {
		"01-01": "New Year's Day",                 // ✅ Fixed
		"01-26": "Chinese New Year",               // 📅 Movable (Lunar)
		"01-27": "Chinese New Year Holiday",       // 📅 Movable (Lunar)
		"02-26": "Hari Raya Puasa",                // 📅 Movable (Islamic)
		"04-14": "Good Friday",                    // 📅 Movable (Christian)
		"05-01": "Labour Day",                     // ✅ Fixed
		"05-05": "Hari Raya Haji",                 // 📅 Movable (Islamic)
		"05-31": "Vesak Day (falls on Wednesday)", // 📅 Movable (Buddhist)
		"08-09": "National Day",                   // ✅ Fixed
		"11-15": "Deepavali",                      // 📅 Movable (Hindu)
		"12-25": "Christmas Day",                  // ✅ Fixed
	},
	models.VN: {
		"01-01": "New Year's Day",                       // ✅ Fixed
		"01-25": "Lunar New Year's Eve",                 // 📅 Movable (Lunar)
		"01-26": "Lunar New Year",                       // 📅 Movable (Lunar)
		"01-27": "Tet Holiday",                          // 📅 Movable (Lunar)
		"01-28": "Tet Holiday",                          // 📅 Movable (Lunar)
		"01-29": "Tet Holiday",                          // 📅 Movable (Lunar)
		"01-30": "Tet Holiday",                          // 📅 Movable (Lunar)
		"04-04": "Hung Kings Commemoration Day",         // 📅 Movable (Lunar)
		"04-30": "Reunification Day (falls on Sunday)",  // ✅ Fixed, 🔁 Observed
		"05-01": "International Workers' Day",           // ✅ Fixed
		"05-02": "Day off in lieu of Reunification Day", // 🔁 Observed for 04-30
		"09-02": "National Day (falls on Saturday)",     // ✅ Fixed, 🔁 Observed
		"09-04": "Day off in lieu of National Day",      // 🔁 Observed for 09-02
	},
	models.MY: {
		"01-01": "New Year's Day",                      // ✅ Fixed
		"01-26": "Chinese New Year",                    // 📅 Movable (Lunar)
		"01-27": "Chinese New Year Holiday",            // 📅 Movable (Lunar)
		"02-01": "Federal Territory Day",               // ✅ Fixed
		"02-26": "Hari Raya Aidilfitri",                // 📅 Movable (Islamic)
		"02-27": "Hari Raya Aidilfitri Holiday",        // 📅 Movable (Islamic)
		"05-01": "Labour Day",                          // ✅ Fixed
		"05-05": "Hari Raya Haji",                      // 📅 Movable (Islamic)
		"05-31": "Wesak Day",                           // 📅 Movable (Buddhist)
		"06-05": "Agong’s Birthday",                    // ✅ Fixed (likely 1st Mon of June)
		"07-05": "Islamic New Year",                    // 📅 Movable (Islamic)
		"08-31": "National Day (Merdeka)",              // ✅ Fixed
		"09-16": "Malaysia Day",                        // ✅ Fixed
		"10-04": "Prophet’s Birthday (Maulidur Rasul)", // 📅 Movable (Islamic)
		"11-15": "Deepavali",                           // 📅 Movable (Hindu)
		"12-25": "Christmas Day",                       // ✅ Fixed
	},
	models.ID: {
		"01-01": "New Year's Day",               // ✅ Fixed
		"01-26": "Chinese New Year",             // 📅 Movable (Lunar)
		"02-26": "Hari Raya Idul Fitri",         // 📅 Movable (Islamic)
		"02-27": "Hari Raya Idul Fitri Holiday", // 📅 Movable (Islamic)
		"03-10": "Isra Mi'raj",                  // 📅 Movable (Islamic)
		"03-30": "Good Friday",                  // 📅 Movable (Christian)
		"05-01": "Labour Day",                   // ✅ Fixed
		"05-05": "Hari Raya Haji",               // 📅 Movable (Islamic)
		"05-31": "Waisak Day",                   // 📅 Movable (Buddhist)
		"06-05": "Pancasila Day",                // ✅ Fixed
		"07-05": "Islamic New Year",             // 📅 Movable (Islamic)
		"08-17": "Independence Day",             // ✅ Fixed
		"10-04": "Prophet’s Birthday (Maulid)",  // 📅 Movable (Islamic)
		"12-25": "Christmas Day",                // ✅ Fixed
	},
}

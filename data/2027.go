package data

import "github.com/vunquitk11/holidaycheck/models"

// holidays2027 contains public holidays for Singapore, Vietnam, Indonesia, Malaysia in 2027.
// Movable holidays are based on official and projected calendars, may be subject to change.
var holidays2027 = map[models.CountryCode]map[string]string{
	models.SG: {
		"01-01": "New Year's Day",           // ✅ Fixed
		"02-06": "Chinese New Year",         // 📅 Movable (Lunar)
		"02-07": "Chinese New Year Holiday", // 📅 Movable (Lunar)
		"03-26": "Good Friday",              // 📅 Movable (Christian)
		"04-10": "Hari Raya Puasa",          // 📅 Movable (Islamic)
		"05-01": "Labour Day",               // ✅ Fixed
		"05-20": "Vesak Day",                // 📅 Movable (Buddhist)
		"06-07": "Hari Raya Haji",           // 📅 Movable (Islamic)
		"08-09": "National Day",             // ✅ Fixed
		"10-27": "Deepavali",                // 📅 Movable (Hindu)
		"12-25": "Christmas Day",            // ✅ Fixed
	},
	models.VN: {
		"01-01": "New Year's Day",               // ✅ Fixed
		"02-05": "Lunar New Year's Eve",         // 📅 Movable (Lunar)
		"02-06": "Lunar New Year",               // 📅 Movable (Lunar)
		"02-07": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-08": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-09": "Tet Holiday",                  // 📅 Movable (Lunar)
		"02-10": "Tet Holiday",                  // 📅 Movable (Lunar)
		"04-08": "Hung Kings Commemoration Day", // 📅 Movable (Lunar)
		"04-30": "Reunification Day",            // ✅ Fixed
		"05-01": "International Workers' Day",   // ✅ Fixed
		"09-02": "National Day",                 // ✅ Fixed
	},
	models.ID: {
		"01-01": "New Year's Day",                // ✅ Fixed
		"02-06": "Chinese New Year",              // 📅 Movable (Lunar)
		"03-17": "Isra Mi'raj",                   // 📅 Movable (Islamic)
		"04-10": "Hari Raya Idul Fitri",          // 📅 Movable (Islamic)
		"04-11": "Hari Raya Idul Fitri Holiday",  // 📅 Movable (Islamic)
		"05-01": "Labour Day",                    // ✅ Fixed
		"05-20": "Ascension of Jesus Christ",     // 📅 Movable (Christian)
		"05-31": "Waisak Day",                    // 📅 Movable (Buddhist)
		"06-07": "Ascension of Prophet Muhammad", // 📅 Movable (Islamic)
		"08-17": "Independence Day",              // ✅ Fixed
		"10-17": "Prophet’s Birthday (Maulid)",   // 📅 Movable (Islamic)
		"12-25": "Christmas Day",                 // ✅ Fixed
	},
	models.MY: {
		"01-01": "New Year's Day",                      // ✅ Fixed
		"02-01": "Federal Territory Day",               // ✅ Fixed
		"02-06": "Chinese New Year",                    // 📅 Movable (Lunar)
		"02-07": "Chinese New Year Holiday",            // 📅 Movable (Lunar)
		"04-10": "Hari Raya Aidilfitri",                // 📅 Movable (Islamic)
		"04-11": "Hari Raya Aidilfitri Holiday",        // 📅 Movable (Islamic)
		"05-01": "Labour Day",                          // ✅ Fixed
		"05-31": "Wesak Day",                           // 📅 Movable (Buddhist)
		"06-07": "Agong’s Birthday or Hari Raya Haji",  // 📅 Movable (Islamic) or ✅ Fixed (overlap)
		"07-17": "Islamic New Year",                    // 📅 Movable (Islamic)
		"08-31": "National Day (Merdeka)",              // ✅ Fixed
		"09-16": "Malaysia Day",                        // ✅ Fixed
		"10-17": "Prophet’s Birthday (Maulidur Rasul)", // 📅 Movable (Islamic)
		"10-27": "Deepavali",                           // 📅 Movable (Hindu)
		"12-25": "Christmas Day",                       // ✅ Fixed
	},
}

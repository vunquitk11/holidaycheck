# HolidayCheck

HolidayCheck is a Go library for checking public holidays across multiple countries, including Singapore, Vietnam, Indonesia, and Malaysia. Easily determine if a given date is a public holiday.

## Features
- **Check Public Holidays**: Determine if a given date is a public holiday for any supported country
- **Get Holiday Names**: Retrieve the specific name of a public holiday for any date
- **Multi-Country Support**: Currently supports Singapore (SG), Vietnam (VN), Indonesia (ID), and Malaysia (MY)
- **Timezone Handling**: Automatic conversion from UTC to local timezone for accurate holiday checking
- **Comprehensive Coverage**: Includes both fixed and movable holidays (Lunar New Year, Islamic holidays, etc.)
- **Easy Integration**: Simple API with clear error handling
- **Extensible Design**: Easy to add more countries and years

## Installation
```bash
go get github.com/vunquitk11/holidaycheck
```

## Usage
```go
import "github.com/vunquitk11/holidaycheck"

// Check if today is a public holiday
isHoliday, err := holidaycheck.IsTodayPublicHoliday("SG", time.Now())
if err != nil {
    // handle error
}
if isHoliday {
    fmt.Println("Today is a public holiday!")
} else {
    fmt.Println("Today is not a public holiday.")
}

// Get holiday name for a specific date
holidayName, err := holidaycheck.GetHolidayName("SG", time.Now())
if err != nil {
    // handle error
}
if holidayName != "" {
    fmt.Println("Today is:", holidayName)
} else {
    fmt.Println("Today is not a public holiday.")
}
```

## Requirements
- Go 1.16 or later

## License
This project is licensed under the MIT License.

## Contact
For issues or contributions, please create an issue on GitHub or contact us at support@example.com.

## Supported Enums and Countries

HolidayCheck supports the following country codes as enums:
- `SG`: Singapore
- `VN`: Vietnam
- `ID`: Indonesia
- `MY`: Malaysia

Each country code corresponds to a specific timezone:
- `SG`: Asia/Singapore
- `VN`: Asia/Ho_Chi_Minh
- `ID`: Asia/Jakarta
- `MY`: Asia/Kuala_Lumpur

## Input Time
- The input time should be in UTC to ensure accurate conversion to local time zones for holiday checking.

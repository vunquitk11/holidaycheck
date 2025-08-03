# HolidayCheck

HolidayCheck is a Go library for checking public holidays across multiple countries, including Singapore, Vietnam, Indonesia, and Malaysia. Easily determine if a given date is a public holiday.

## Features
- Check public holidays by country and year.
- Supports multiple countries and time zones.
- Easy to extend and customize.

## Installation
```bash
go get github.com/vunquitk11/holidaycheck
```

## Usage
```go
import "github.com/vunquitk11/holidaycheck"

isHoliday, err := holidaycheck.IsTodayPublicHoliday("SG", time.Now())
if err != nil {
    // handle error
}
if isHoliday {
    fmt.Println("Today is a public holiday!")
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

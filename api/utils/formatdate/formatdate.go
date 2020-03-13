package formatdate

import (
	"time"
)

func FormatDate(date string) {
	newDate, err := time.Parse("2006-01-02", date)

	// return newDate, err
}

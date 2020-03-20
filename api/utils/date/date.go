package date

import (
	"regexp"
)

func ValiDate(date string) bool {
	re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")

	return re.MatchString(date)
}

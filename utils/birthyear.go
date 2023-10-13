package utils

import (
	"fmt"
	"strconv"
	"time"
)

func FindBirthYear(fnr string) string {
	thisYear := time.Now().Year()
	birthYear, _ := strconv.Atoi("20" + fnr[4:6])

	// check if person is born in the 1900 or 2000
	// limit: persons over 100 years old wil be 1 year old with this function
	if birthYear >= thisYear {
		return "19" + fnr[4:6]
	} else {
		return fmt.Sprint(birthYear)
	}
}

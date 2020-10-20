package ValidationHelper

import (
	"fmt"
	"strings"
)

func GetValidationMsg(property string, msg string) string {
	splitProperty := strings.Split(property, ".")
	return fmt.Sprintf(
		"Key: '%v' Error:Field validation for '%v' %v",
		property,
		splitProperty[1],
		msg,
	)
}

package ValidationHelper

import (
	"fmt"
	"strings"
)

func GetValidationMsg(property string, msg string) string {
	splitProperty := strings.Split(property, ".")
	lastProperty := splitProperty[len(splitProperty)-1]
	return fmt.Sprintf(
		"Key: '%v' Error:Field validation for '%v' %v",
		property,
		lastProperty,
		msg,
	)
}

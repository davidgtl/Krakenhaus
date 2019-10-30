package validators

import (
	. "github.com/tobyhede/go-underscore"
	"regexp"
)

func NonEmptyStr(fields...string) bool {
	return Every(func(v string) bool {
		return v != ""
	}, fields)
}

func ValidEmail(fields...string) bool {
	ex := regexp.MustCompile(`.*@.*\..*`)
	return Every(fields, func(v string) bool {
		return ex.MatchString(v)
	})
}

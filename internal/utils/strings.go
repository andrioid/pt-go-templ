package utils

import "regexp"

var stripNonAlphaNumericRegex *regexp.Regexp

func StripNonAlphaNumeric(in string) string {
	return stripNonAlphaNumericRegex.ReplaceAllString(in, "")
}

func init() {
	stripNonAlphaNumericRegex = regexp.MustCompile("[^a-zA-Z0-9-/]+")
}

package text

import (
	"encoding/base64"
	"regexp"
	"strings"
)

var (
	reAlphaNum8    *regexp.Regexp
	reDigits20     *regexp.Regexp
	reDigits4To15  *regexp.Regexp
	reAlphaNum7    *regexp.Regexp
	reAlphaNum8To9 *regexp.Regexp
)

func init() {
	reAlphaNum8 = regexp.MustCompile("^[A-Z0-9]{8,8}$")
	reDigits20 = regexp.MustCompile("^[0-9]{20,20}$")
	reDigits4To15 = regexp.MustCompile("^[0-9]{4,15}$")
	reAlphaNum7 = regexp.MustCompile("^[A-Z0-9]{7,7}$")
	reAlphaNum8To9 = regexp.MustCompile("^[A-Z0-9]{8,9}$")
}

func ValidAlphaNum8(s string) bool {
	return reAlphaNum8.MatchString(s)
}

func ValidAlphaNum7(s string) bool {
	return reAlphaNum7.MatchString(s)
}

func ValidAlphaNum8To9(s string) bool {
	return reAlphaNum8To9.MatchString(s)
}

func ValidDigits20(s string) bool {
	return reDigits20.MatchString(s)
}

func ValidDigits4To15(s string) bool {
	return reDigits4To15.MatchString(s)
}

func EscapeLatex(s string) string {
	s2 := strings.Replace(s, "\\", "\\textbackslash", -1)
	s2 = strings.Replace(s2, "&", "\\&", -1)
	s2 = strings.Replace(s2, "%", "\\%", -1)
	s2 = strings.Replace(s2, "$", "\\$", -1)
	s2 = strings.Replace(s2, "#", "\\#", -1)
	s2 = strings.Replace(s2, "_", "\\_", -1)
	s2 = strings.Replace(s2, "{", "\\{", -1)
	s2 = strings.Replace(s2, "}", "\\}", -1)
	s2 = strings.Replace(s2, "~", "\\textasciitilde", -1)
	return strings.Replace(s2, "^", "\\textasciicircum", -1)

}

func CheckAddZero(s string) string {
	if s[0] == '.' {
		return "0" + s
	} else {
		return s
	}
}

func Base64Encode(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

func Base64Decode(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(decoded)

}

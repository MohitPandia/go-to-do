package constants

var Regex = struct {
	REGEXP_MOBILE_NUMBER string
	REGEXP_EMAIL         string
	REGEXP_ALPHANUMBERIC string
	REGEXP_NUMBERIC      string
}{
	REGEXP_MOBILE_NUMBER: "^[5-9]{1}[0-9]{9}$",
	REGEXP_EMAIL:         `[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}`,
	REGEXP_ALPHANUMBERIC: `^[a-zA-Z0-9]+$`,
	REGEXP_NUMBERIC:      `^[0-9]+$`,
}

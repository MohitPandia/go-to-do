package smerrors

var errorType = struct {
	validation         string
	server             string
	Unauthorized       string
	conflict           string
	ServiceUnavailable string
	Downstream         string
}{
	validation:         "validation",
	server:             "server",
	Unauthorized:       "unauthorized",
	conflict:           "conflict",
	ServiceUnavailable: "service unavailable",
	Downstream:         "downstream",
}

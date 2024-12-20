package constants

var Headers = struct {
	OrgHeader     string
	OrgAccHeader  string
	AUTHORIZATION string
	AceessSecret  string
	Sandbox       string
	TraceID       string
}{
	OrgHeader:     "x-zoop-org",
	OrgAccHeader:  "x-zoop-org-acc",
	AUTHORIZATION: "Authorization",
	AceessSecret:  "x-zoop-key",
	TraceID:       "x-zoop-trace-id",
	Sandbox:       "sandbox",
}

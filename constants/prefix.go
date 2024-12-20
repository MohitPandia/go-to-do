package constants

var Prefix = struct {
	TraceID     string
	USER         string
	OrgAccounts string

	FILE                string
	RequestResponseLogs string
	AccessKeyLive       string
	SecretKeyLive       string
	AccessKeyTest       string
	SecretKeyTest       string
	RULE                string
	KnowledgeBase       string
}{
	USER:                 "user",
	OrgAccounts:         "orgacc",
	TraceID:             "rule-engine-trace-ID-",
	FILE:                "file",
	RequestResponseLogs: "log",
	AccessKeyLive:       "ak_live",
	SecretKeyLive:       "sk_live",
	AccessKeyTest:       "ak_test",
	SecretKeyTest:       "sk_test",
	RULE:                "rule",
	KnowledgeBase:       "kbase",
}

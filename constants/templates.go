package constants

var Templates = struct {
	GMAIL_SYNC string
}{
	GMAIL_SYNC: "gmail_sync",
}

var TemplateKeys = struct {
	INPUT_KEY  string
	OUTPUT_KEY string
}{
	INPUT_KEY:  "InputData",
	OUTPUT_KEY: "OutputData",
}

var TemplatesEnum = map[string]bool{
	Templates.GMAIL_SYNC: true,
}

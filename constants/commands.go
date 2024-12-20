package constants

var Command = struct {
	SEED        string
	MIGRATE     string
	DROP_TABLES string
}{
	SEED:        "seed",
	MIGRATE:     "migrate",
	DROP_TABLES: "droptables",
}

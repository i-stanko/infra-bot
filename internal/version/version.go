package version

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func Info() string {
	return "infra-bot version: " + Version + " (commit: " + Commit + ", built: " + BuildDate + ")"
}

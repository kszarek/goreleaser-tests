package version

var (
	// version is set by the build process
	version = "dev"
	// commit is set by the build process
	commit = "none"
	// date is set by the build process
	date = "unknown"
)

// GetVersion returns the version string
func GetVersion() string {
	return version
}

// GetCommit returns the commit hash
func GetCommit() string {
	return commit
}

// GetDate returns the build date
func GetDate() string {
	return date
}

// GetFullVersion returns a formatted version string with all information
func GetFullVersion() string {
	return version + " (" + commit + ") built on " + date
}

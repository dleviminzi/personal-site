package data

// DatabaseInfo contains information required to establish connection to database
type DatabaseInfo struct {
	DBName   string /* either test or site */
	Host     string
	Port     string
	User     string
	Password string
}

package orm

var autoMigrate, logMode, seedDB bool
var hostDB, userDB, portDB, passwordDB, nameDB, sslDB, dialect string


func init() {
	dialect = "postgres"
	hostDB = "localhost"
	portDB = "5432"
	userDB = "postgres"
	passwordDB = "iü"
	nameDB = "go-rest-server"
	sslDB = "disable"
	seedDB = true
	logMode = true
	autoMigrate = true
}


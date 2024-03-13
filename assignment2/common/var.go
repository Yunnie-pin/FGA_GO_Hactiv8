package common

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "curut"
	dbname   = "orders_by"
)

func GetDBHost() string {
	return host
}

func GetDBPort() int {
	return port
}

func GetDBUsername() string {
	return user
}

func GetDBPassword() string {
	return password
}

func GetDBName() string {
	return dbname
}

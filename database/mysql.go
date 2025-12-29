package database

var connection string

func init() {
	connection = "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8"
}

func GetDatabase() string {
	return connection
}

package util

//DbProperties contains all the data necessary to connect to a database
type DbProperties struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

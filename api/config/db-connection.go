package config

import "fmt"

type MySQLConnection struct {
	Config *DBConfig
}

// Connect connects to a MySQL database and returns a dns string or error.
func MySQLConnectionString() (string, error) {
	// get variable data from the env
	newConfig := DBConfig{
		User:     "root",
		Password: "testDB",
		Host:     "127.0.0.1",
		Port:     "3306",
		DB:       "myDB",
		TimeZone: "Local",
	}
	//free sql db confiq
	// newConfig := DBConfig{
	// 	User:     "root",
	// 	Password: "testDB",
	// 	Host:     "sql12.freesqldatabase.com",
	// 	Port:     "3306",
	// 	DB:       "sql12715595",
	// 	TimeZone: "Local",
	// }
	dns := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s"
	return fmt.Sprintf(dns, newConfig.User, newConfig.Password, newConfig.Host, newConfig.Port, newConfig.DB, newConfig.TimeZone), nil
}

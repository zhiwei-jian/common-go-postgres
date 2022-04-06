package compostgres

import (
	"database/sql"
	"fmt"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type AppContext struct {
	db *sql.DB
}

// func main() {
// 	c, err := connectDB()
// 	defer c.db.Close()

// 	if err != "" {
// 		print(err)
// 	}
// 	c.Create()
// 	fmt.Println("add action done!")

// 	c.Read()
// 	fmt.Println("get action done!")

// 	c.Update()
// 	fmt.Println("update action done!")

// 	c.Delete()
// 	fmt.Println("delete action done!")

// }

func connectDB(config *PostgresConfig) (c *AppContext, errorMessage string) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Failed to connect Postgres " + err.Error())

	} else {

		fmt.Println("Successfully connected!")
	}

	err = db.Ping()
	if err != nil {

		fmt.Println("DBPing error ：" + err.Error())

	} else {
		fmt.Println("DBPingSuccess")

	}

	return &AppContext{db}, ""
}

package infrastructure

import (
	"fmt"
	"locator-api/utils"

	// go-sql-driver/mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupModels : initializing mysql database
func SetupModels() *gorm.DB {
	get := utils.GetEnvWithKey
	USER := get("DB_USER")
	PASS := get("DB_PASS")
	HOST := get("DB_HOST")
	PORT := get("DB_PORT")
	DBNAME := get("DB_NAME")
	TYPE := get("DB_TYPE")

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	if TYPE == "cloudsql" {
		createDBDsn = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			USER,
			PASS,
			HOST,
			DBNAME,
		)
	}

	database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})
	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	if err != nil {
		panic(err.Error())
	}

	return database
}

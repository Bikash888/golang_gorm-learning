package main

import (
	"locator-api/infrastructure"
	"locator-api/utils"
)

func main() {
	utils.LoadEnv()
	db := infrastructure.SetupModels()
	infrastructure.SetupRoutes(db)
}

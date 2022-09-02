package main

import (
	"fmt"
	"mini-project/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB = config.SetupConnectionDatabase()
)

func main() {

	defer config.CloseConnectionDatabase(db)
	r := gin.Default()

	fmt.Println("Hello test connection")

	r.Run()
}
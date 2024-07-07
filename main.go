package main

import (
	"a-tour-of-go/app/controllers"
	"a-tour-of-go/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}

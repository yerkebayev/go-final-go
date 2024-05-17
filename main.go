package main

import (
	"fmt"

	"github.com/yerkebayev/go-final-go/controllers"
	"github.com/yerkebayev/go-final-go/db"
)

var Port = ":50051"

func main() {
	DB := db.Init()
	fmt.Println("Listing to port : ", Port)
	controllers.Controller(Port, DB)
}

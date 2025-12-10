package main

import (
	"fmt"
	"kevinmajesta/testkemas/internal/builder"
	"kevinmajesta/testkemas/internal/entity"
	"kevinmajesta/testkemas/pkg/database"
	"kevinmajesta/testkemas/pkg/server"
	"os"
)

// func main() {
// 	_, err := configs.NewConfig(".env")
// 	checkError(err)

// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	db, err := database.OpenDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	fmt.Println("Database connection established:", db)
	db.AutoMigrate(&entity.Products{})

	publicRoutes := builder.BuildPublicRoutes(db)

	srv := server.NewServer("app", publicRoutes)
	srv.Run()
}

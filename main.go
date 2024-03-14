package main

import (
	"log"
	"net/http"
	"stakeholders/handler"
	"stakeholders/model"
	"stakeholders/repo"
	"stakeholders/routing"
	"stakeholders/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connection_url := "user=postgres password=super dbname=SOA port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	database.Exec("DROP SCHEMA IF EXISTS users CASCADE")
	database.Exec("DROP SCHEMA IF EXISTS people CASCADE")
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Person{})
	database.Exec("INSERT INTO users (id, username, password, email, role, is_active) VALUES (1, 'turista1', 'turista1', 'turista@gmail.com', 2, true) ON CONFLICT DO NOTHING")
	database.Exec("INSERT INTO users (id, username, password, email, role, is_active) VALUES (2, 'autor', 'autor', 'autor@gmail.com', 1, true) ON CONFLICT DO NOTHING")
	database.Exec("INSERT INTO users (id, username, password, email, role, is_active) VALUES (3, 'admin', 'admin', 'admin@gmail.com', 0, true) ON CONFLICT DO NOTHING")

	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	//router := routing.SetupRoutes(userHandler)

	personRepo := &repo.PersonRepository{DatabaseConnection: database}
	personService := &service.PersonService{PersonRepo: personRepo}
	personHandler := &handler.PersonHandler{PersonService: personService}

	router := routing.SetupRoutes(userHandler, personHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8081", router))
}

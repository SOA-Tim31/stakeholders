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

	database.Exec("DROP TABLE IF EXISTS users CASCADE")
	database.Exec("DROP TABLE IF EXISTS people CASCADE")
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Person{})
	database.Exec("INSERT INTO users (id, username, password, role, is_active, verification_token) VALUES (1, 'turista1', 'turista1', 2, true, 'aea71b9a6ca84d75dcbe8a78f8f6a1f3cde0f7e8569ba0b03946b57580379189')")
	database.Exec("INSERT INTO users (id, username, password, role, is_active, verification_token) VALUES (2, 'autor', 'autor', 1, true, 'bea71b9a6ca84d75dcbe8a78f8f6a1f3cde0f7e8569ba0b03946b57580379189') ")
	database.Exec("INSERT INTO users (id, username, password, role, is_active, verification_token) VALUES (3, 'admin', 'admin', 0, true, 'cea71b9a6ca84d75dcbe8a78f8f6a1f3cde0f7e8569ba0b03946b57580379189')")
	database.Exec("INSERT INTO people (id, user_id, name, surname, email, profile_image, bio, quote) VALUES (1, 1, 'John', 'Doe', 'john.doe@example.com', 'profile.jpg', 'Software Engineer', 'Live life to the fullest.')")
	database.Exec("INSERT INTO people (id, user_id, name, surname, email) VALUES (2, 2, 'Jane', 'Smith', 'jane.smith@example.com')")
	database.Exec("INSERT INTO people (id, user_id, name, surname, email) VALUES (3, 3, 'Alice', 'Johnson', 'alice.johnson@example.com') ")

	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	personRepo := &repo.PersonRepository{DatabaseConnection: database}
	personService := &service.PersonService{PersonRepo: personRepo}
	personHandler := &handler.PersonHandler{PersonService: personService}

	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepo: userRepo, PersonService: personService}
	userHandler := &handler.UserHandler{UserService: userService}

	//router := routing.SetupRoutes(userHandler)

	router := routing.SetupRoutes(userHandler, personHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8081", router))
}

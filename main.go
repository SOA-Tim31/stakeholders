package main

import (
	"log"
	"net/http"
	"os"
	"stakeholders/handler"
	"stakeholders/model"
	"stakeholders/repo"
	"stakeholders/routing"
	"stakeholders/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	//connection_url := "user=postgres password=super dbname=explorer port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	database.Exec("DROP TABLE IF EXISTS users CASCADE")
	database.Exec("DROP TABLE IF EXISTS people CASCADE")
	database.Exec("DROP TABLE IF EXISTS app_ratings CASCADE")
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Person{})
	database.AutoMigrate(&model.AppRating{})
	database.Exec("INSERT INTO users (username, password, role, is_active, verification_token) VALUES ('turista1', 'turista1', 2, true, 'aea71b9a6ca84d75dcbe8a78f8f6a1f3cde0f7e8569ba0b03946b57580379189')")
	database.Exec("INSERT INTO users (username, password, role, is_active, verification_token) VALUES ('autor', 'autor', 1, true, 'bea71b9a6ca84d75dcbe8a78f8f6a1f3cde0f7e8569ba0b03946b57580379189') ")
	database.Exec("INSERT INTO users (username, password, role, is_active, verification_token) VALUES ('admin', 'admin', 0, true, 'cea71b9a6ca84d75dcbe8a78f8f6a1f3cde0f7e8569ba0b03946b57580379189')")
	database.Exec("INSERT INTO people (user_id, name, surname, email, profile_image, bio, quote) VALUES (1, 'John', 'Doe', 'john.doe@example.com', 'https://upload.wikimedia.org/wikipedia/commons/4/41/Profile-720.png', 'Software Engineer', 'Live life to the fullest.')")
	database.Exec("INSERT INTO people (user_id, name, surname, email) VALUES (2, 'Jane', 'Smith', 'jane.smith@example.com')")
	database.Exec("INSERT INTO people (user_id, name, surname, email) VALUES (3, 'Alice', 'Johnson', 'alice.johnson@example.com') ")
	database.Exec("INSERT INTO app_ratings (user_id, rating, description, date_created) VALUES (2, 9, 'Odlicna aplikacija, svaka cast', '2024-03-16 21:15:41.765+01')")
	database.Exec("INSERT INTO app_ratings (user_id, rating, description, date_created) VALUES (3, 2, 'Uzas aplikacija, nepregledno', '2024-03-11 20:15:41.765+01')")

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

	appRatingRepo := &repo.AppRatingRepository{DatabaseConnection: database}
	appRatingService := &service.AppRatingService{AppRatingRepository: appRatingRepo}
	appRatingHandler := &handler.AppRatingHandler{AppRatingService: appRatingService}

	//router := routing.SetupRoutes(userHandler)

	router := routing.SetupRoutes(userHandler, personHandler, appRatingHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8082", router))
}

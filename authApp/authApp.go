package authApp

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/izaakdale/auth/domain"
	"github.com/izaakdale/auth/service"
	"github.com/jmoiron/sqlx"
)

func Start() {

	sanityCheck()
	router := mux.NewRouter()
	dbClient := getDbClient()

	authRepoDb := domain.NewAuthRepoDb(dbClient)
	authHandler := AuthHandler{
		service.NewAuthRepoDb(authRepoDb),
	}

	router.HandleFunc("auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("auth/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("auth/verify", authHandler.Verify).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Println("Starting the auth server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Missing env variable ")
	}
}

func getDbClient() *sqlx.DB {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	schema := os.Getenv("DB_SCHEMA")

	client, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, pass, schema))

	if err != nil {
		fmt.Println("Unable to connect to DB")
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

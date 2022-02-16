package main

import (
	studentstruct "goelster/StudentStruct"
	"goelster/controllers"
	"goelster/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main(){
	IntDb()
	log.Println("Starting The Http Server on port 8090")
	router:=mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	
	router.HandleFunc("/get/{id}",controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/get",controllers.GetPersonData).Methods("GET")
	router.HandleFunc("/newget/{id}",controllers.GetByID).Methods("GET")
	router.HandleFunc("/update",controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}",controllers.DeletePersonById).Methods("DELETE")

	c:=cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})

	handler:=c.Handler(router)
	log.Fatal(http.ListenAndServe(":8090",handler))

	
}



func IntDb(){
	config:=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "Thilina1999@",
			DB:         "test5",
		}
		connectionString :=database.GetConnectionString(config)
		err:=database.Connect(connectionString)
		if err !=nil{
			panic(err.Error())
		}
		database.Migrate(&studentstruct.Person{})
}
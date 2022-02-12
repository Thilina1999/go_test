package controllers

import (
	"encoding/json"
	studentstruct "goelster/StudentStruct"
	"goelster/database"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
) 

func CreatePerson(w http.ResponseWriter,r *http.Request){
	requestBody, _:=ioutil.ReadAll((r.Body))
	var person studentstruct.Person
	json.Unmarshal(requestBody,&person)

	database.Connector.Create(person)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}
func GetPersonData(w http.ResponseWriter,r *http.Request){
	person:= []studentstruct.Person{}
	database.Connector.Find(&person)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(person)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	key:=vars["id"]
  	var person studentstruct.Person
	
	database.Connector.Find(&person,key)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(person)
}

func GetByID(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	key:=vars["id"]
	person:= []studentstruct.Person{}

	database.Connector.Find(&person)

	for _, persons:= range person {
			s,err:= strconv.Atoi(key)
			if err==nil{
				if persons.ID == s{
					w.Header().Set("Content-Type","application/json")
					json.NewEncoder(w).Encode(persons)
					return
				}
			}
	}
}
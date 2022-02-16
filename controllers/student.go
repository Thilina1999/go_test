package controllers

import (
	"encoding/json"
	"fmt"
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
	
	persons:= []studentstruct.Person{}
	
	database.Connector.Find(&persons)
	 total :=0
	for _, person := range persons{
		fmt.Println(person.FirstName)
		total+=person.Age
		
	} 
	var totals studentstruct.Total
	totals.Total=total
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(totals)
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

func UpdatePersonByID(w http.ResponseWriter, r *http.Request){
	
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person studentstruct.Person
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	var person studentstruct.Person
	id, _ :=strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)

}
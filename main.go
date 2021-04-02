package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {

	dbConfig := mysql.NewConfig()
	dbConfig.User = "root"
	dbConfig.Passwd = "root"
	dbConfig.Addr = "godockerDB"
	dbConfig.DBName = "onepuchman"
	dbConfig.Net = "tcp"

	//"root:root@tcp(godockerDB)/onepuchman"
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		db := dbConn()
		defer db.Close()
		selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := Employee{}
		for selDB.Next() {
			var id int
			var name, city string
			err = selDB.Scan(&id, &name, &city)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.City = city
		}
		fmt.Println(emp)
		// Handle the request
	}).Methods(http.MethodGet, http.MethodPut, http.MethodPatch)

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	// selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// emp := Employee{}
	// for selDB.Next() {
	// 	fmt.Println("entrou aqui")
	// 	var id int
	// 	var name, city string
	// 	err = selDB.Scan(&id, &name, &city)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	emp.Id = id
	// 	emp.Name = name
	// 	emp.City = city

	//

}

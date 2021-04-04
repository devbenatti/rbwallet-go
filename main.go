package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devbenatti/rbwallet-go/driven/database"
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/account"
	action "github.com/devbenatti/rbwallet-go/driver/webApi/action/account"

	"github.com/devbenatti/rbwallet-go/driven/uuid"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	service "github.com/devbenatti/rbwallet-go/service/account"
	"github.com/gorilla/mux"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		db := database.GetConnection()

		defer db.Close()

		uuidGenerator := uuid.NewUuidGenerator()
		id := uuidGenerator.Generate()
		accountDTO := dto.NewAccountDTO(id.Val(), "08749067940")
		accDAO := dao.NewAccountDAO(db)
		accService := service.NewAccountService(accDAO)
		accService.Create(accountDTO)
		id = uuidGenerator.Generate()
		result, err := accService.FindOne(id)
		if err != nil {
			fmt.Println("Nenhuma conta encontrada para esse ID")
		}
		fmt.Println(result)
		// selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")

		// if err != nil {
		// 	panic(err.Error())
		// }

		// var emp = Employee{}

		// for selDB.Next() {
		// 	var id int
		// 	var name, city string
		// 	err = selDB.Scan(&id, &name, &city)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// 	emp.Id = id
		// 	emp.Name = name
		// 	emp.City = city
		// }
		// fmt.Println(emp)
		// Handle the request
	}).Methods(http.MethodGet, http.MethodPut, http.MethodPatch)

	uuidGen := uuid.NewUuidGenerator()

	db := database.GetConnection()

	accDAO := dao.NewAccountDAO(db)

	accService := service.NewAccountService(accDAO)

	createAccount := action.NewCreateAccount(uuidGen, accService)

	r.HandleFunc("/teste", createAccount.Handle()).Methods(http.MethodPost)

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
}

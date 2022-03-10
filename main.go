package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"go-hexagonal-architecture/repository"
	"go-hexagonal-architecture/service"
)

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		"root",
		"P@ssw0rd",
		"13.76.163.73",
		3306,
		"banking",
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	_ = customerService
}

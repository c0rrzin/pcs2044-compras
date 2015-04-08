package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/c0rrzin/router"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	DefRoutes()
	router.RouteAll()
	// Setuping Database
	db, err := gorm.Open("sqlite3", "/tmp/compras.db")
	if err != nil {
		panic(err)
	}

	db.DB()
	db.DropTable(&OrdemDeCompra{})
	db.DropTable(&Item{})
	db.CreateTable(&OrdemDeCompra{})
	db.CreateTable(&Item{})
	db.AutoMigrate(&OrdemDeCompra{}, &Item{})
	// Then you could invoke `*sql.DB`'s functions with it
	// db.DB().Ping()
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
	log.Fatal(s.ListenAndServe())
	fmt.Println("Starting server at port" + s.Addr)
}

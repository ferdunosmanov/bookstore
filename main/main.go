package main

import (
	"log"
	"net/http"

	"github.com/ferdunosmanov/go-bookstore/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("172.17.0.2:3306", r))

}

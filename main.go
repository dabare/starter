//go:generate goversioninfo -icon=icon.ico
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	go checkGUI()
	err := http.ListenAndServe(":9090",nil)
	if err!= nil {
		log.Fatal(err)
	}
}

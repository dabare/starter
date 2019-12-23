//go:generate goversioninfo -icon=icon.ico
package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	checkGUI()
}

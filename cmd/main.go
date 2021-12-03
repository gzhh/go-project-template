package main

import (
	"demo/cmd/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.Execute()
}

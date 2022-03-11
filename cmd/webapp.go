package main

import (
	"demo/cmd/webapp"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	webapp.Execute()
}

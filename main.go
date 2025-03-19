package main

import (
	"fmt"

	"github.com/vnFuhung2903/postgresql/config"
)

func main() {
	db := config.ConnectPostgresDb()
	fmt.Println(db)
}

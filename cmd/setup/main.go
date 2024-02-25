package main

import (
	"fmt"

	"todox/internal"
	"todox/internal/migrations"

	"github.com/leapkit/core/db"
	"github.com/paganotoni/tailo"
)

func main() {
	err := tailo.Setup()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("✅ Tailwind CSS setup successfully")
	conn, err := internal.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Migrations ran successfully")
}

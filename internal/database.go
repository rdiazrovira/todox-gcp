package internal

import (
	"fmt"

	"github.com/leapkit/core/db"
	"github.com/leapkit/core/envor"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL    = envor.Get("DATABASE_URL", "./todox.db")
	DatabaseDriver = envor.Get("DATABASE_DRIVER", "sqlite3")

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(DatabaseURL, db.WithDriver(DatabaseDriver))
)

func init() {
	fmt.Println("This is the database URL I am reading: ", DatabaseURL)
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"tudai.seminario.golang.practica/internal/config"
	"tudai.seminario.golang.practica/internal/database"
	"tudai.seminario.golang.practica/internal/service/bookstore"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		panic(err)
	}

	service, err := bookstore.New(db, cfg)

	httpService := bookstore.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}

func createSchema(db *sqlx.DB) error {

	schema := `CREATE TABLE IF NOT EXISTS books (
								id        integer     NOT NULL PRIMARY KEY autoincrement,
								name      varchar(50) NOT NULL,
								language  varchar(50) NOT NULL,
								status    varchar(15) NOT NULL,
								genre     varchar(35) NOT NULL,
								editorial varchar(50) NOT NULL,
								author    varchar(30) NOT NULL,
								publicado varchar(10) NOT NULL,
								price     float       NOT NULL
							);`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

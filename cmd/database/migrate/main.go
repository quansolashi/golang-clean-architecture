package main

import (
	database "clean-architecture/internal/infrastructure/database/gorm"
	"clean-architecture/internal/infrastructure/database/migrate"
	"context"
	"flag"
	"fmt"
	"time"
)

var (
	dbsocket   string
	dbhost     string
	dbport     string
	dbname     string
	dbusername string
	dbpassword string
)

func main() {
	startedAt := time.Now()
	if err := run(); err != nil {
		panic(err)
	}
	const format = "2006-01-02 15:04:05"
	fmt.Printf("Done: %s -> %s\n", startedAt.Format(format), time.Now().Format(format))
}

func run() error {
	flag.StringVar(&dbsocket, "db-socket", "tcp", "mysql server protocol")
	flag.StringVar(&dbhost, "db-host", "127.0.0.1", "mysql server host")
	flag.StringVar(&dbport, "db-port", "3306", "mysql server port")
	flag.StringVar(&dbname, "db-name", "clean-architecture", "mysql database name")
	flag.StringVar(&dbusername, "db-username", "root", "mysql auth username")
	flag.StringVar(&dbpassword, "db-password", "root", "mysql auth password")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	db, err := newDBClient(dbsocket, dbhost, dbport, dbname, dbusername, dbpassword)
	if err != nil {
		return err
	}

	fmt.Println("database migration will begin")
	if err := migrate.Run(ctx, db); err != nil {
		return err
	}
	fmt.Println("database migration has been completed")
	return nil
}

func newDBClient(
	socket, host, port, db, username, password string,
) (*database.Client, error) {
	params := &database.Params{
		Socket:   socket,
		Host:     host,
		Port:     port,
		Database: db,
		Username: username,
		Password: password,
	}
	return database.NewDatabaseClient(params)
}

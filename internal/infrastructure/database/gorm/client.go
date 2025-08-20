package database

import (
	"fmt"

	dmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Client struct {
	DB *gorm.DB
}

type Params struct {
	Socket   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func NewDatabaseClient(params *Params) (*Client, error) {
	db, err := NewDatabase(params)
	if err != nil {
		return nil, err
	}
	client := &Client{
		DB: db,
	}
	return client, nil
}

func NewDatabase(params *Params) (*gorm.DB, error) {
	conf := &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var addr string
	switch params.Socket {
	case "tcp":
		addr = fmt.Sprintf("%s:%s", params.Host, params.Port)
	case "unix":
		addr = params.Host
	}

	dsn := &dmysql.Config{
		User:                 params.Username,
		Passwd:               params.Password,
		Net:                  params.Socket,
		Addr:                 addr,
		DBName:               params.Database,
		ParseTime:            true,
		Collation:            "utf8mb4_general_ci",
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		Params:               map[string]string{"charset": "utf8mb4"},
		MaxAllowedPacket:     4194304,
	}
	return gorm.Open(mysql.Open(dsn.FormatDSN()), conf)
}

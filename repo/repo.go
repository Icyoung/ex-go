package repo

import (
	"context"
	"ex/model"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
)

func ConnectMysql() *gorm.DB {
	dial, dialErr := ssh.Dial("tcp", "192.168.1.6:22", &ssh.ClientConfig{
		User: "icy",
		Auth: []ssh.AuthMethod{
			ssh.Password("510510"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if dialErr != nil {
		panic(dialErr)
	}

	mysql.RegisterDialContext("ssh+tcp", func(_ context.Context, addr string) (net.Conn, error) {
		return dial.Dial("tcp", addr)
	})

	dsn := "root:Abcd1234.@ssh+tcp(localhost:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local"
	db, mysqlErr := gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if mysqlErr != nil {
		panic(mysqlErr)
	}

	dbErr := db.AutoMigrate(&model.User{})
	if dbErr != nil {
		panic(dbErr)
	}
	return db
}

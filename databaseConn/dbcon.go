package databaseConn

import (
	"fmt"
	"os"

	"github.com/RamazanZholdas/APIWithGin/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err        error
	Db         *gorm.DB
	user       = os.Getenv("USER")
	password   = os.Getenv("PASSWORD")
	dbEndpoint = os.Getenv("DB_ENDPOINT")
	dbName     = os.Getenv("DB_NAME")
)

func ConnectToDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, dbEndpoint, dbName)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func SyncDB() {
	Db.AutoMigrate(&structs.Song{})
}

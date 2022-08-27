package databaseConn

import (
	"github.com/RamazanZholdas/APIWithGin/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	Db  *gorm.DB
)

func ConnectToDB() {
	dns := "root:fsdfas1341231@tcp(127.0.0.1:3306)/restapiproject?parseTime=true"
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func SyncDB() {
	Db.AutoMigrate(&structs.Song{})
}

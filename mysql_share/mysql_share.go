package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Shared struct {
	Key string `gorm:"primary_key"`
	Val int
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	// CREATE DATABASE concurrent CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
	_ = mysql.MySQLError{} // for auto import
	dataSource := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123qwe", "127.0.0.1", "3306", "concurrent")
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	dbRet := db.AutoMigrate(&Shared{})
	if dbRet.Error != nil {
		log.Fatal(dbRet.Error)
	}
	dbRet = db.Save(&Shared{Key: "a", Val: 0})
	if dbRet.Error != nil {
		log.Fatal(dbRet.Error)
	}

	//
	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			a := Shared{Key: "a"}
			dbRet := db.Find(&a)
			if dbRet.Error != nil {
				log.Println(dbRet.Error)
			}
			a.Val += 1
			dbRet = db.Save(&a)
			if dbRet.Error != nil {
				log.Println(dbRet.Error)
			}
		}()
	}
	wg.Wait()

	//
	a := Shared{Key: "a"}
	_ = db.Find(&a)
	if a.Val != 30 {
		log.Fatalf("expected: 30, actual: %v\n", a.Val)
	}
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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
	db.DB().SetMaxOpenConns(50) // MySQL default max connections = 150

	// init a shared row
	migrateRet := db.AutoMigrate(&Shared{})
	if migrateRet.Error != nil {
		log.Fatal(migrateRet.Error)
	}
	initRet := db.Save(&Shared{Key: SharedKey0, Val: 0})
	if initRet.Error != nil {
		log.Fatal(initRet.Error)
	}

	// concurrently update the row
	bTime := time.Now()
	n := 1000
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			updated, err := Incr(db)
			log.Printf("incr: err: %v, updatedValue: %v\n", err, updated)
		}()
	}
	wg.Wait()
	log.Println("updates duration:", time.Since(bTime))

	// check the row after updates
	a := Shared{Key: SharedKey0}
	_ = db.Find(&a)
	if a.Val != n {
		log.Fatalf("expected: %v, actual: %v\n", n, a.Val)
	} else {
		log.Printf("expected: %v, actual: %v\n", n, a.Val)
		log.Println("ngon")
	}
}

const SharedKey0 = "SharedKey0"

type Shared struct {
	Key string `gorm:"primary_key"`
	Val int
}

func Incr(db *gorm.DB) (updatedValue int, err error) {
	tx := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if tx.Error != nil {
		err := fmt.Errorf("error when create tx: %v", tx.Error)
		return 0, err
	}
	a := Shared{Key: SharedKey0}
	findRet := tx.Set("gorm:query_option", "FOR UPDATE").Find(&a)
	if findRet.Error != nil {
		err := fmt.Errorf("error when select: %v", findRet.Error)
		tx.Rollback()
		return 0, err
	}
	a.Val += 1
	SaveRet := tx.Save(&a)
	if SaveRet.Error != nil {
		err := fmt.Errorf("error when update: %v", SaveRet.Error)
		tx.Rollback()
		return 0, err
	}
	commitRet := tx.Commit()
	if commitRet.Error != nil {
		err := fmt.Errorf("error when commit: %v", commitRet.Error)
		return 0, err
	}
	updatedValue = a.Val
	return updatedValue, err
}

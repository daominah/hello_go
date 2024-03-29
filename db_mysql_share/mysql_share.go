package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	// CREATE DATABASE test_concurrent CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci

	//mysqlHosts := []string{"10.100.50.101", "10.100.50.102"} // group replication
	mysqlHosts := []string{"127.0.0.1"}
	runningHosts := make([]string, 0)

	dbs := make(map[string]*gorm.DB)
	for _, nodeHost := range mysqlHosts {
		dataSource := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			"root", "123qwe", nodeHost, "3306", "test_concurrent")
		db, err := gorm.Open(mysql.Open(dataSource),
			&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			log.Println(err)
			continue
		}
		runningHosts = append(runningHosts, nodeHost)

		// init a shared row
		migrateRet := db.AutoMigrate(&Shared{})
		if migrateRet != nil {
			log.Fatalf("error create table `shareds`: %v", migrateRet)
		}
		initRet := db.Save(&Shared{Key: SharedKey0, Val: 0})
		if initRet.Error != nil {
			log.Fatal(initRet.Error)
		}
		dbs[nodeHost] = db
	}
	if len(runningHosts) == 0 {
		log.Fatal("no running host")
	}

	// concurrently update the row
	bTime := time.Now()
	n := 500
	nErrs := 0
	nTries := 0
	mutex := &sync.Mutex{}
	wg := sync.WaitGroup{}
	limitGoroutines := make(chan bool, 40)
	for i := 0; i < n; i++ {
		wg.Add(1)
		limitGoroutines <- true
		go func() {
			defer func() {
				wg.Add(-1)
				<-limitGoroutines
			}()
			host := runningHosts[rand.Intn(len(runningHosts))]
			db, found := dbs[host]
			if !found {
				return
			}
			updated := 0
			job := func() error {
				var err error
				updated, err = Incr(db)
				return err
			}
			nTriesLocal, err := retry(50*time.Millisecond, 10, job)
			log.Printf("host %v incr: err: %v, updatedValue: %v\n",
				host, err, updated)

			mutex.Lock()
			nTries += nTriesLocal
			if err != nil {
				nErrs += 1
			}
			mutex.Unlock()
		}()
	}
	wg.Wait()
	log.Println("updates duration:", time.Since(bTime))

	// check the row after updates
	a := Shared{Key: SharedKey0}
	_ = dbs[runningHosts[0]].Find(&a)
	log.Printf("expected: %v, actual: %v, nErrs: %v, nTries: %v\n",
		n, a.Val, nErrs, nTries)
}

const SharedKey0 = "SharedKey0"

type Shared struct {
	Key       string `gorm:"primary_key"`
	Val       int
	UpdatedAt time.Time
}

func Incr(db *gorm.DB) (updatedValue int, err error) {
	tx := db.Begin(&sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if tx.Error != nil {
		err := fmt.Errorf("error when create tx: %v", tx.Error)
		return 0, err
	}
	a := Shared{Key: SharedKey0}
	findRet := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&a)
	if findRet.Error != nil {
		err := fmt.Errorf("error when select: %v", findRet.Error)
		tx.Rollback()
		return 0, err
	}
	a.Val += 1
	a.UpdatedAt = time.Now()
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
	return updatedValue, nil
}

func Incr2(db *gorm.DB) (updatedValue int, err error) {
	a := Shared{Key: SharedKey0}
	findRet := db.Find(&a)
	if findRet.Error != nil {
		err := fmt.Errorf("error when select: %v", findRet.Error)
		return 0, err
	}
	a.Val += 1
	a.UpdatedAt = time.Now()
	SaveRet := db.Save(&a)
	if SaveRet.Error != nil {
		err := fmt.Errorf("error when update: %v", SaveRet.Error)
		return 0, err
	}
	updatedValue = a.Val
	return updatedValue, nil
}

// https://dev.mysql.com/doc/refman/8.0/en/group-replication-limitations.html
func retry(baseBackoff time.Duration, maxRetries int, f func() error) (
	nTries int, retErr error) {
	backoff := baseBackoff
	for i := 1; i <= maxRetries; i++ {
		err := f()
		if err == nil {
			return i, nil
		}
		retErr = err
		approximateBackoff := backoff * time.Duration(80+rand.Intn(40)) / 100
		time.Sleep(approximateBackoff)
		backoff = time.Duration(float64(backoff) * 1.5)
	}
	return maxRetries, retErr
}

// Before you execute the program, Launch `cqlsh` and execute:
//	create keyspace dbname0 with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
//	create table dbname0.tweet(timeline text, id UUID, text text, PRIMARY KEY(timeline, id));
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "dbname0"
	cluster.Consistency = gocql.Quorum // majority of nodes
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("error CreateSession: %v", err)
	}
	defer session.Close()

	// insert a tweet
	err = session.Query(
		`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec()
	if err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	if err := session.Query(
		`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`, "me").
		Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("read Tweet:", id, text)

	// list all tweets
	iter := session.Query(
		`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	for iter.Scan(&id, &text) {
		fmt.Println("all Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

package cassandra

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Cassandra session extablished...")
}

func GetSession() *gocql.Session {
	return session
}

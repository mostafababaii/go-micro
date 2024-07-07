package config

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func NewDBConnection() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "auth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		return nil, err
	}
	return session, nil
}

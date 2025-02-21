package database

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"gopkg.in/mgo.v2"
)

// DBCollection encapsulates mgo collection
type DBCollection struct {
	*mgo.Collection
}

// Close safely close a collection and associate session
func (collection DBCollection) Close() {
	if collection.Database == nil {
		return
	}
	if collection.Database.Session == nil {
		return
	}
	collection.Database.Session.Close()
}

// DBConfig contains configurations for database
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// DBSession encapsulates mgo session
type DBSession struct {
	*mgo.Session
}

// Collection copy existing session and return collection by name
func (session *DBSession) Collection(name string) DBCollection {
	return DBCollection{session.Copy().DB("").C(name)}
}

// New returns new data store instance
func New() (*DBSession, error) {
	var config DBConfig
	viper.UnmarshalKey("database", &config)
	url := fmt.Sprintf("%s:%d", config.Host, config.Port)
	info := &mgo.DialInfo{
		Addrs:    []string{url},
		Username: config.Username,
		Password: config.Password,
		Database: config.DBName,
		Timeout:  60 * time.Second,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}
	return &DBSession{session}, nil
}

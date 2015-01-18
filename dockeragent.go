package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"log"
	"github.com/spf13/viper"
)

var (
	mongoDBHosts string
	verbose = false
)

const (
	Database       = "test"
	InfoCollection = "people"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	setupDockeragent()
	parseConfig()
	/*
		session, err := GetDBConnection("localhost")

		c := session.DB(Database).C(InfoCollection)

		c.Insert(&Person{"Ale", "+555381169639"}, &Person{"Cla", "+555384333233"})

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)

		if err == nil {
			fmt.Println("Phone:", result.Phone)
		}
	*/
}

func setupDockeragent() {
	viper.SetDefault("dockeragent.verbose", false)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yml")
	viper.AddConfigPath("/etc/dockeragent/")
	viper.AddConfigPath("$HOME/.dockeragent")
	viper.ReadInConfig()
}

func parseConfig() {
	mongoDBHosts = viper.GetString("db.hostname")
	dockeragent := viper.Get("dockeragent").(map[string]interface {})
	verbose = bool(dockeragent["verbose"].(bool))

	if verbose {
		log.Println("MongoDB Hostname", mongoDBHosts)
	}
}

func usage() {
	fmt.Println("usage: dockeragent -h <DBHostname>")
}

func GetDBConnection(url string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(url)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	return
}

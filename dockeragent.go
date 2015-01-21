package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"log"
	"flag"
	"labix.org/v2/mgo/bson"
)

var (
	dockerHost      = *flag.String("docker", "localhost:2376", "Dockerhost to poll")
	mongoDBHost     = *flag.String("mongo", "localhost:27017", "Host running mongodb to poll")
	mongoCollection = *flag.String("collection", "dockeragent", "Collection in mongoDB")
	mongoDB         = *flag.String("db", "dockeragent", "Collection in mongoDB")
	verbose         = *flag.Bool("v", false, "Be verbose or not")
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	flag.Parse()

	parseConfig()
	session, err := GetDBConnection(mongoDBHost)
	defer session.Close()

	c := session.DB(mongoDBHost).C(mongoCollection)

	c.Insert(&Person{"Ale", "+555381169639"}, &Person{"Cla", "+555384333233"})

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)

	if err == nil {
		fmt.Println("Phone:", result.Phone)
	}
}

func setupDockeragent() {

}

func parseConfig() {
	if verbose == true {
		fmt.Println("DockerHost", dockerHost)
		fmt.Println("MongoHost", mongoDBHost)
		fmt.Println("MongoDB", mongoDB)
		fmt.Println("MongoCollection", mongoCollection)
		fmt.Println("Verbose", verbose)
	}
}

func GetDBConnection(url string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(url)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return
}

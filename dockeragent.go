package main

import (
	"flag"
	"labix.org/v2/mgo/bson"
	"fmt"
	"labix.org/v2/mgo"
	"log"
)

type Config struct {
	DockerHost      string
	MongoDBHost     string
	MongoCollection string
	MongoDB         string
	Verbose         bool
}

var config = Config{*flag.String("docker", "localhost:2376", "Dockerhost to poll"),
	*flag.String("mongo", "localhost:27017", "Host running mongodb to poll"),
	*flag.String("collection", "dockeragent", "Collection in mongoDB"),
	*flag.String("db", "dockeragent", "Collection in mongoDB"),
	*flag.Bool("v", false, "Be verbose or not"),
}


const (
	ImagesCollection = "images"
	ContainersCollection = "containers"
	HostsCollection = "hosts"
)

type Person struct {
	Name  string
	Phone string
}

type Image struct {
	Created uint64 `json:"Created" bson:"Created,omitempty"`
	Id string `json:"Id" bson:"_id,omitempty"`
	ParentId string `json:"ParentId" bson:"ParentId"`
	RepoTags map[int]string `json:"RepoTags" bson:"RepoTags"`
	Size uint64 `json:"Size" bson:"Size"`
	VirtualSize uint64 `json:"VirtualSize" bson:"VirtualSize"`
}

func main() {
	flag.Parse()

	parseConfig()

	session, err := GetDBConnection(config.MongoDBHost)

	defer session.Close()

	c := session.DB(config.MongoDB).C(config.MongoCollection)

	c.Insert(&Person{"Ale", "+555381169639"}, &Person{"Cla", "+555384333233"})

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)

	if err == nil {
		fmt.Println("Phone:", result.Phone)
	}
}

func parseConfig() {
	fmt.Println(config.Verbose)
	if config.Verbose == true {
		fmt.Println("DockerHost", config.DockerHost)
		fmt.Println("MongoHost", config.MongoDBHost)
		fmt.Println("MongoDB", config.MongoDB)
		fmt.Println("MongoCollection", config.MongoCollection)
		fmt.Println("Verbose", config.Verbose)
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

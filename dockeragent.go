package main

import (
	"flag"
	"fmt"
	"labix.org/v2/mgo"
	"log"
	"github.com/jmcvetta/napping"
	"time"
)

type Config struct {
	DockerHost            string
	MongoDBHost           string
	MongoDB               string
	Interval              int
	Verbose               bool
}

var config = Config{*flag.String("docker", "localhost:2376", "Dockerhost to poll"),
	*flag.String("mongo", "localhost:27017", "Host running mongodb to poll"),
	*flag.String("db", "dockeragent", "Name of mongoDB"),
	*flag.Int("t", 30, "Poll interval in seconds"),
	*flag.Bool("v", false, "Be verbose or not"),
}

const (
	ImagesCollection     = "images"
	ContainersCollection = "containers"
	HostsCollection      = "hosts"
)

type Person struct {
	Name  string
	Phone string
}

type Image struct {
	Created     uint64 `json:"Created" bson:"Created,omitempty"`
	Id          string `json:"Id" bson:"_id,omitempty"`
	ParentId    string `json:"ParentId" bson:"ParentId"`
	RepoTags    []string `json:"RepoTags" bson:"RepoTags"`
	Size        uint64 `json:"Size" bson:"Size"`
	VirtualSize uint64 `json:"VirtualSize" bson:"VirtualSize"`
}

func main() {
	flag.Parse()

	parseConfig()

	s := napping.Session{}
	url := "http://188.166.55.153:2376/images/json"
	session, err := GetDBConnection(config.MongoDBHost)

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	c := session.DB(config.MongoDB).C(ImagesCollection)

	for ; ; {
		PollDockerDaemon(c, s, url)
		time.Sleep(time.Duration(config.Interval * 1000 * 1000 * 1000))
	}
}

func PollDockerDaemon(c *mgo.Collection, s napping.Session, url string) {
	var images []Image
	resp, err := s.Get(url, nil, &images, nil)

	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() == 200 {
		if config.Verbose {
			fmt.Println("Got Images from server")
		}
		writeImagesToDB(c, images)
	}
}

func writeImagesToDB(c *mgo.Collection, images []Image) {
	for index := range images {
		err := c.Insert(images[index])

		if err != nil {
			fmt.Println(err)
		}

		if err == nil && config.Verbose {
			fmt.Println("Written image to database")
		}
	}
}


func parseConfig() {
	fmt.Println(config.Verbose)
	if config.Verbose == true {
		fmt.Println("DockerHost", config.DockerHost)
		fmt.Println("MongoHost", config.MongoDBHost)
		fmt.Println("MongoDB", config.MongoDB)
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

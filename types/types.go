package types

type DockerHost struct {
	Id   string `json:"Id" bson:"_id,omitempty"`
	Name string `json:"Name" bson:"Name,omitempty"`
	Url  string `json:"Url" bson:"Url,omitempty"`
}

type Image struct {
	Created     uint64   `json:"Created" bson:"Created,omitempty"`
	Id          string   `json:"Id" bson:"_id,omitempty"`
	ParentId    string   `json:"ParentId" bson:"ParentId"`
	RepoTags    []string `json:"RepoTags" bson:"RepoTags"`
	Size        uint64   `json:"Size" bson:"Size"`
	VirtualSize uint64   `json:"VirtualSize" bson:"VirtualSize"`
}

type Port struct {
	IP          string `json:"IP" bson:"IP,omitempty"`
	PrivatePort int    `json:"PrivatePort" bson:"PrivatePort,omitempty"`
	PublicPort  int    `json:"PublicPort" bson:"PublicPort,omitempty"`
	Type        string `json:"Type" bson:"Type"`
}

type Container struct {
	Command string   `json:"Command bson:"Command,omitempty"`
	Created uint64   `json:"Created" bson:"Created,omitempty"`
	Image   string   `json:"Image" bson:"Image,omitempty"`
	Id      string   `json:"Id" bson:"_id,omitempty"`
	Names   []string `json:"Names" bson:"Names"`
	Ports   []Port   `json:"Ports" bson:"Ports"`
	Status  string   `json:"Status" bson:"Status"`
}

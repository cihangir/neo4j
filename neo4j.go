package neo4j

import (
	"net/http"
	"strconv"
)

type Neo4j struct {
	Client          *http.Client
	BaseUrl         string
	NodeUrl         string
	BatchUrl        string
	RelationshipUrl string
	IndexNodeUrl    string
}

func Connect(host string, port int) *Neo4j {
	if host == "" {
		host = "http://127.0.0.1"
	}

	if port == 0 {
		port = 7474
	}

	url := host + ":" + strconv.Itoa(port)

	baseUrl := url + "/db/data"
	return &Neo4j{
		Client:          http.DefaultClient,
		BaseUrl:         baseUrl,
		NodeUrl:         baseUrl + "/node",
		BatchUrl:        baseUrl + "/batch",
		IndexNodeUrl:    baseUrl + "/index/node",
		RelationshipUrl: baseUrl + "/relationship",
	}
}

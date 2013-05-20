package neo4j

import (
	"net/http"
	"strconv"
)

type Noe4j struct {
	Client          *http.Client
	BaseUrl         string
	NodeUrl         string
	RelationshipUrl string
	IndexNodeUrl    string
}

func Connect(host string, port int) *Noe4j {
	if host == "" {
		host = "http://127.0.0.1"
	}

	if port == 0 {
		port = 7474
	}

	url = host + ":" + strconv.Itoa(port)

	return &Noe4j{
		Client:          http.DefaultClient,
		BaseUrl:         url + "/db/data",
		NodeUrl:         BaseUrl + "/node",
		IndexNodeUrl:    BaseUrl + "/index/node",
		RelationshipUrl: BaseUrl + "/relationship",
	}
}

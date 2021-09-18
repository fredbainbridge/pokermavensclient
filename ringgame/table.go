package ringgame

import (
	"encoding/json"
	"log"

	"github.com/fredbainbridge/pokerapimodels"
	"github.com/fredbainbridge/pokermavensclient"
)

func Tables(url string, password string) ([]pokerapimodels.Table, error) {

	var params = make(map[string]string)
	params["Command"] = "RingGamesList"
	params["Password"] = password
	params["JSON"] = "Yes"
	params["Fields"] = "Name,Game"
	response, err := pokermavensclient.EncodedHttpClientRequest("POST", url, params)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var mavResp mavenTableResponse
	json.Unmarshal([]byte(response), &mavResp)
	var tables []pokerapimodels.Table
	for i := 0; i < mavResp.RingGames; i++ {
		tables = append(tables, pokerapimodels.Table{Game: mavResp.Game[i], Name: mavResp.Name[i]})
	}
	return tables, nil
}

type mavenTableResponse struct {
	Result    string
	RingGames int
	Name      []string
	Game      []string
}

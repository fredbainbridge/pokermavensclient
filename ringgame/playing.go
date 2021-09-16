package ringgame

import (
	"encoding/json"
	"log"

	"github.com/fredbainbridge/pokerapimodels"
	"github.com/fredbainbridge/pokermavensclient"
)

func Playing(url string, password string, tableNames ...string) ([]pokerapimodels.SeatedPlayer, error) {
	var seatedPlayers []pokerapimodels.SeatedPlayer
	var params = make(map[string]string)

	params["Command"] = "RingGamesPlaying"
	params["Password"] = password
	params["JSON"] = "Yes"
	if len(tableNames) == 0 {
		tables, err := Tables(url, password)
		if err != nil {
			log.Fatalf("Cannot find tables. %q", err)
			return nil, err
		}
		for i := range tables {
			params["Name"] = tables[i].Name
			seatedPlayers = append(seatedPlayers, getSeatedPlayers(tables[i].Name, params, url, password)...)
		}
	} else {
		for i := range tableNames {
			params["Name"] = tableNames[i]
			seatedPlayers = append(seatedPlayers, getSeatedPlayers(tableNames[i], params, url, password)...)
		}
	}
	return seatedPlayers, nil
}

func getSeatedPlayers(tableName string, params map[string]string, url string, password string) []pokerapimodels.SeatedPlayer {
	var seatedPlayers []pokerapimodels.SeatedPlayer
	var mavResp mavenGamePlayingResponse
	params["Name"] = tableName
	response, err := pokermavensclient.EncodedHttpClientRequest("POST", url, params)
	if err != nil {
		log.Fatalf("Failed when getting %q", tableName)
	}
	json.Unmarshal([]byte(response), &mavResp)
	for i := 0; i < mavResp.Count; i++ {
		seatedPlayers = append(seatedPlayers, pokerapimodels.SeatedPlayer{Name: mavResp.Player[i], TableName: tableName})
	}
	return seatedPlayers
}

type mavenGamePlayingResponse struct {
	Result string
	Count  int
	Player []string
	Seat   []int
	Chips  []int
	Net    []int
	Away   []int
}

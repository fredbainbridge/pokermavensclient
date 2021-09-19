package tournament

import (
	"log"

	"github.com/fredbainbridge/pokermavensclient"
)

func Offline(tableName string, now string, url string, password string) error {
	var params = make(map[string]string)
	params["Command"] = "TournamentsOffline"
	params["Password"] = password
	params["JSON"] = "Yes"
	params["Name"] = tableName
	params["Now"] = now
	_, err := pokermavensclient.EncodedHttpClientRequest("POST", url, params)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

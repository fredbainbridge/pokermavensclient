package ringgame

import (
	"github.com/fredbainbridge/pokermavensclient"
)

func AdminMessage(tableName string, message string, url string, password string) (string, error) {
	var params = make(map[string]string)
	params["Command"] = "RingGamesMessage"
	params["Name"] = tableName
	params["Message"] = message
	params["Password"] = password
	params["JSON"] = "Yes"
	return pokermavensclient.EncodedHttpClientRequest("POST", url, params)
}

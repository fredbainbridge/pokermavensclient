package tournament_test

import (
	"testing"

	"github.com/fredbainbridge/pokermavensclient/tournament"
)

const url = "http://192.168.0.100:8087/api"
const password = "password"

func TestTournamentTablesExist(t *testing.T) {
	results, err := tournament.Tables(url, password)
	if err != nil {
		t.Error(err)
	}
	if len(results) == 0 {
		t.Error("No tables returned.")
	}
}

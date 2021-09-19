package tournament_test

import (
	"testing"

	"github.com/fredbainbridge/pokermavensclient/tournament"
)

func TestOffline(t *testing.T) {
	tables, _ := tournament.Tables(url, password)

	err := tournament.Offline(tables[0].Name, "Yes", url, password)
	if err != nil {
		t.Errorf("Failed with error: %q", err)
	}
}

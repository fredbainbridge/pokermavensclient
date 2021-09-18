package ringgame_test

import (
	"testing"

	"github.com/fredbainbridge/pokermavensclient/ringgame"
)

func TestOffline(t *testing.T) {
	tables, _ := ringgame.Tables(url, password)

	err := ringgame.Offline(tables[0].Name, "Yes", url, password)
	if err != nil {
		t.Errorf("Failed with error: %q", err)
	}
}

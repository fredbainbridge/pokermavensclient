package ringgame_test

import (
	"testing"

	"github.com/fredbainbridge/pokermavensclient/ringgame"
)

func TestTablesExist(t *testing.T) {
	results, err := ringgame.Tables(url, password)
	if err != nil {
		t.Error(err)
	}
	if len(results) == 0 {
		t.Error("No tables returned.")
	}
}

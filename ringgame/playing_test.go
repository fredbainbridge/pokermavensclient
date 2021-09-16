package ringgame_test

import (
	"testing"

	"github.com/fredbainbridge/pokermavensclient/ringgame"
)

const url = "http://192.168.0.100:8087/api"
const password = "password"

func TestPlaying(t *testing.T) {
	players, err := ringgame.Playing(url, password)
	if err != nil {
		t.Errorf("Failed with error: %q", err)
	}
	for i := range players {
		if len(players[i].Name) == 0 {
			t.Errorf("Empty player name.")
		}
	}
}

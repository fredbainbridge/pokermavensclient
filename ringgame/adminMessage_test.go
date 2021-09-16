package ringgame_test

import (
	"encoding/json"
	"testing"

	"github.com/fredbainbridge/pokermavensclient/ringgame"
)

type mavenResponse struct {
	Result string
	Error  string
}

func TestAdminMessageTableDoesNotExist(t *testing.T) {
	results, err := ringgame.AdminMessage("test table", "test message", url, password)
	if err != nil {
		t.Errorf("Error: %q", err)
	}

	var response mavenResponse
	marshalErr := json.Unmarshal([]byte(results), &response)
	if marshalErr != nil {
		t.Errorf("Unexpected response %q", results)
	}
	expected := "Ring game not found"
	if response.Error != expected {
		t.Errorf("Looking for %q, found %q", expected, results)
	}
}

func TestAdminMessageTableDoesExist(t *testing.T) {
	results, err := ringgame.AdminMessage("test table", "test message", url, password)
	if err != nil {
		t.Errorf("Error: %q", err)
	}

	var response mavenResponse
	marshalErr := json.Unmarshal([]byte(results), &response)
	if marshalErr != nil {
		t.Errorf("Unexpected response %q", results)
	}
	expected := "Ring game not found"
	if response.Error != expected {
		t.Errorf("Looking for %q, found %q", expected, results)
	}
}

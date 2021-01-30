package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

var testGames = []struct {
	in  string
	out string
}{
	{"easy", "Easy"},
	{"empty", "Empty"},
	{"basic", "Basic"},
	{"medium", "Medium"},
	{"hard", "Hard"},
	{"master", "Master"},
}

func TestNewGame(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Log("no config file")
		t.Fail()
	}
	for _, tt := range testGames {
		level := tt.in
		url := fmt.Sprintf("%s/api/game?level=%s", config.GetURL(), level)
		resp, err := http.Get(url)

		if err != nil {
			t.Fail()
		}

		body, err := ioutil.ReadAll(resp.Body)
		var game map[string]interface{}
		err = json.Unmarshal(body, &game)
		if err != nil {
			t.Fail()
		}

		if l, ok := game["level"]; !ok || l != tt.out {
			t.Fail()
		}

		bodyJSON := gojsonschema.NewStringLoader(string(body))
		schema := gojsonschema.NewReferenceLoader("file://./schema/new-game-res.json")

		result, err := gojsonschema.Validate(schema, bodyJSON)

		if err != nil || !result.Valid() {
			for _, desc := range result.Errors() {
				t.Logf("- %s (%s)\n", desc, desc.Value())
			}
			t.Fail()
		}
	}
}

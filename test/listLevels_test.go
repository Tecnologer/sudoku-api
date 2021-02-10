package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func TestListLevels(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Log("no config file")
		t.Fail()
	}

	url := fmt.Sprintf("%s/api/game/levels", config.GetURL())
	resp, err := http.Get(url)

	if err != nil {
		t.Fail()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}
	bodyJSON := gojsonschema.NewStringLoader(string(body))
	schema := gojsonschema.NewReferenceLoader("file://./schema/list-levels.json")

	result, err := gojsonschema.Validate(schema, bodyJSON)

	if err != nil || !result.Valid() {
		for _, desc := range result.Errors() {
			t.Logf("- %s (%s)\n", desc, desc.Value())
		}
		t.Fail()
	}
}

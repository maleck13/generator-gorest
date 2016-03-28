package api_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"<%=basePackage %>/<%=baseName %>/api"
	"<%=basePackage %>/<%=baseName %>/config"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//todo add tests for other index methods

func TestIndexRoute(t *testing.T) {
	config.SetGlobalConfig("../config/config.json")
	server := httptest.NewServer(api.NewRouter())
	defer server.Close()
	resp, err := http.Get(server.URL + "/api/")
	assert.NoError(t, err, "did not expect an error")
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "did not expect an error reading body")
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	assert.NoError(t, err, "did not expect an error reading body")
	if v, ok := data["example"]; ok {
		assert.Equal(t, "value", v, "expected values to match")

	} else {
		assert.Fail(t, "expected returned json to have example key")
	}

}

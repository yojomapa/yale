package cli

import (
	"testing"
	"reflect"
	"os"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func TestConstructorError(t *testing.T) {
	content, _ := ioutil.ReadFile("../test/resources/marathon_tasks_response.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
        defer ts.Close()

	os.Args = append(os.Args, "--endpoint="+ts.URL, "list", "--image-filter=nginx", "--tag-filter=latest")
	//os.Args = append(os.Args, "--smoke-request=/smoke")
	RunApp()
	v := reflect.ValueOf(stackManager).Elem()
	stacks := v.FieldByName("stacks")
	assert.Equal(t, 1, stacks.Len(), "Cli should instantiate at least one stack")
}

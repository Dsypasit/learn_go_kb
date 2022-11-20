package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "Ong", "info": "gopher"}`))
}

func TestMakeHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	fmt.Println(server.URL)

	want := &Response{
		ID:   1,
		Name: "Ong",
		Info: "gopher",
	}

	t.Run("Happy server response", func(t *testing.T) {
		resp, _ := MakeHTTPCall(server.URL)
		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expeted (%v), got (%v)", want, resp)
		}
	})
}

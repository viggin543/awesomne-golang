package __parameterized_test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{"https://dummyjson.com/products/1", http.StatusOK},
		{"https://dummyjson.com/products/foo", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")

	for _, u := range urls {
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
		resp, err := http.Get(u.url)
		if err != nil {
			t.Fatal("\t\tShould be able to Get the url.", ballotX, err)
		}
		defer resp.Body.Close()
		t.Log("\t\tShould be able to Get the url.", checkMark)

		if resp.StatusCode == u.statusCode {
			t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
		} else {
			t.Errorf("\t\tShould have a \"%d\" status. %v %v", u.statusCode, ballotX, resp.StatusCode)
		}

	}

}

package __basic_unit_test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// a vanilla test with no asserts lib

func Test_get_product_req_should_return_200(t *testing.T) {
	url := "https://dummyjson.com/products/1"
	statusCode := 200
	t.Log("Given the need to test downloading content.")
	t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
	resp, err := http.Get(url)
	if err != nil {
		//this fails the test
		t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
	}
	t.Log("\t\tShould be able to make the Get call.", checkMark)

	defer resp.Body.Close()

	if resp.StatusCode == statusCode {
		t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
	} else {
		t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
			statusCode, ballotX, resp.StatusCode)
	}

}

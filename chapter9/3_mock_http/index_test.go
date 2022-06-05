package __mock_http

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

func mockServer() *httptest.Server {
	// this starts a new server on a random port
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}))
}

// EVERY PIECE OF CODE HAS THREE CONSUMERS ( WRITER, READER, RUNNER )
// TWO OF WHICH ARE HUMAN

// TEST TRIPLE A FORMAT
// ARRANGE
// ACT
// ASSERT
func TestXmlResponse(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()
	t.Log("Given the need to test downloading content.")
	t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)
	//use white space to emphasize the ACT in a test

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
	}
	t.Log("\t\tShould be able to make the Get call.", checkMark)

	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
	}
	t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)

	var d Document
	if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
		t.Fatal("\t\tShould be able to unmarshal the response.", ballotX, err)
	}
	t.Log("\t\tShould be able to unmarshal the response.", checkMark)

	if len(d.Channel.Items) == 1 {
		t.Log("\t\tShould have \"1\" item in the feed.", checkMark)
	} else {
		t.Error("\t\tShould have \"1\" item in the feed.", ballotX, len(d.Channel.Items))
	}

}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

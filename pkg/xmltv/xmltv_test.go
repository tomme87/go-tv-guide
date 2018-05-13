package xmltv

import (
	"testing"
	"os"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"encoding/json"
)

func TestUnmarshalXML(t *testing.T) {
	f, err := os.Open("../../test/example.xml")
	if err != nil {
		t.Fatal(err)
	}

	var tv Tv

	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&tv)
	if err != nil {
		t.Fatal(err)
	}

	//spew.Dump(tv)

	_, err = json.Marshal(&tv)
	if err != nil {
		t.Fatal(err)
	}

	//fmt.Println(string(j))
}

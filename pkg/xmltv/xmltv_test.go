package xmltv

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestUnmarshalXML(t *testing.T) {
	f, err := os.Open("../../test/example.xml")
	if err != nil {
		t.Fatal(err)
	}

	var tv Tv

	err = tv.LoadXML(f)
	if err != nil {
		t.Fatal(err)
	}

	//spew.Dump(tv)

	j, err := json.Marshal(&tv)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}

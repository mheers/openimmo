
package openimmo


import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)


func Decoder() {
	dec := xml.NewDecoder(os.Stdin)
	var doc Openimmo
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
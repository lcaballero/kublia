package shared

import (
	"encoding/json"
	"fmt"
)

func JsonOrPanic(conf interface{}) []byte {
	bin, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}
	return bin
}

func MustShowJson(c []byte) {
	fmt.Println(string(c))
}

func ShowJsonOrPanic(c interface{}, err error) {
	if err != nil {
		panic(err)
	}
	MustShowJson(JsonOrPanic(c))
}

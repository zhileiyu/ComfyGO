package main

import (
	"encoding/json"
	"fmt"
)

// type Tdata [T]
// type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

func main() {
	testData := make(map[string]string)
	testData["0"] = "tt"
	testData["1"] = "ss"
	vd, _ := json.Marshal(testData)
	fmt.Println(string(vd))
	//ci := comfyGO.NewComfyClient()
	//println(ci.Ready())Ready
	//err := ci.PromptEnqueue()
	//if err != nil {
	//	println(err.Error())
	//}Error
}

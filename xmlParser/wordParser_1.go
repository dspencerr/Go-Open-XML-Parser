package xmlParser

import (
	"github.com/dspencerr/uscore"
)

var groups1 []interface {}
var tblResArray [][]string

func parseWord1(needle string, arrayNode string, sectionNode string) [][]string {

	tblResArray = [][]string{}
	var instance []interface {}
	groups1 = instance

	for _, group := range dataSelection {

		result := uscore.HasVal(group.(map[string]interface{}), needle)

		if !result {
			continue
		}

		dataTypeSwitch1(group.(map[string]interface {})[arrayNode], sectionNode)

		tblResArray, _ = uscore.HasKey(groups1, "p", "r", "t")
	}

	return tblResArray
}

func dataTypeSwitch1(data interface{}, key string){

	switch myMap := data.(type) {
		case string:
	// do nothing
		case map[string]interface{}:
		handleMaps1(myMap, key)
		case []interface{}:
		handleArray1(myMap, key)
	}
}

func handleMaps1(data map[string]interface{}, key string){
	for k, val := range data {
		if(k == key){
			groups1 = append(groups1, val)
		} else{
			dataTypeSwitch1(val, key)
		}
	}
}

func handleArray1(data []interface{}, key string){
	for _, val := range data {
		dataTypeSwitch1(val, key)
	}
}
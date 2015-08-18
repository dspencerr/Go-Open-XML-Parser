package xmlParser

import (
	"github.com/dspencerr/uscore"
)

var groups4 []interface {}
var tblResArray4 [][]string

func parseWord4(needle string, arrayNode string, sectionNode string) [][]string {

	tblResArray = [][]string{}
	var instance []interface {}
	groups4 = instance

	for _, group := range dataSelection {

		result := uscore.HasVal(group.(map[string]interface{}), needle)

		if !result {
			continue
		}

		dataTypeSwitch4(group.(map[string]interface {})[arrayNode], sectionNode)

		tblResArray4, _ = uscore.HasKey(groups1, "p", "r", "t")
	}

	return tblResArray4
}

func dataTypeSwitch4(data interface{}, key string){

	switch myMap := data.(type) {
		case string:
	// do nothing
		case map[string]interface{}:
		handleMaps4(myMap, key)
		case []interface{}:
		handleArray4(myMap, key)
	}
}

func handleMaps4(data map[string]interface{}, key string){
	for k, val := range data {
		if(k == key){
			groups1 = append(groups1, val)
		} else{
			dataTypeSwitch4(val, key)
		}
	}
}

func handleArray4(data []interface{}, key string){
	for _, val := range data {
		dataTypeSwitch4(val, key)
	}
}
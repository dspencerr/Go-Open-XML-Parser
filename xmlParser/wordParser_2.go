package xmlParser

import (
	"github.com/dspencerr/uscore"
)

var groups2 []interface {}
var tblResArray2 [][]string

func parseWord2(needle string, arrayNode string, sectionNode string) [][]string {

	tblResArray = [][]string{}
	var instance []interface {}
	groups2 = instance

	for _, group := range dataSelection {

		result := uscore.HasVal(group.(map[string]interface{}), needle)

		if !result {
			continue
		}

		dataTypeSwitch2(group.(map[string]interface {})[arrayNode], sectionNode)

		tblResArray2, _ = uscore.HasKey(groups1, "tc", "", "t")
	}

	return tblResArray2
}

func dataTypeSwitch2(data interface{}, key string){

	switch myMap := data.(type) {
		case string:
	// do nothing
		case map[string]interface{}:
		handleMaps2(myMap, key)
		case []interface{}:
		handleArray2(myMap, key)
	}
}

func handleMaps2(data map[string]interface{}, key string){
	for k, val := range data {
		if(k == key){
			groups1 = append(groups1, val)
		} else{
			dataTypeSwitch2(val, key)
		}
	}
}

func handleArray2(data []interface{}, key string){
	for _, val := range data {
		dataTypeSwitch2(val, key)
	}
}
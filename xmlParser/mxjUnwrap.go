package xmlParser
import (
	"github.com/clbanning/mxj"
)

var xmlMap mxj.Map
var dataSelection  []interface {}

func getParsedMap(xmlData []byte) {
	m, err := mxj.NewMapXml(xmlData)
	if err != nil {
		panic("error mapping file")
	}
	xmlMap = m
}

func selectXmlSection(dotLocation string) bool {
	selection, err := xmlMap.ValuesForPath(dotLocation)
	if err != nil {
		return false
	}
	dataSelection = selection
	return true
}
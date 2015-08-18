package xmlParser

import (
	"os"
	"io/ioutil"
)

func ParseXmlDocument(xmlPath string, xmlType string, targetDir string) [][]string {
	var data [][]string
	if xmlType == "docx" {
		 data = parseWordDocs(xmlPath)
	} else if xmlType == "xlsx" {
		 data = parseExcelSheets(xmlType)
	}
	return data
}

func readFileContents(xmlFile string) []byte {
	fileContents, err := os.Open(xmlFile)
	if err != nil {
		panic("could not open document.xml")
	}
	xmlData, err := ioutil.ReadAll(fileContents)
	if err != nil {
		panic("error reading contents");
	}
	return xmlData
}


























//func getHeaderCount(data [][]string) string int{
//	var headerCount int
//	for _, val := range data{
//		if val[0] != "Rev"{ continue }
//		headerCount = len(val)
//		d.Header = strings.Join(val, "\t")
//	}
//	return headerCount;
//}
//
//func getDataFromPattern(data [][]string, count int){
//	for _, val := range data{
//		// This should be a "Tolerance" variable
//		if len(val) >= count -2{
//			d.Result = strings.Join(val, "\t")
//		}
//	}
//}
package xmlParser

import (
	"github.com/dspencerr/docParser/fileWriter"
	"os"
	"strconv"
)

func parseWordDocs(xmlPath string) [][]string {
	var data [][]string
	var xmlFile string
	var xmlData []byte

	/**
	 * =======================
	 * header.xml
	 * =======================
	 */
	for x := 1; x < 5; x++{
		xmlFile = xmlPath + "/word/header"+strconv.Itoa(x)+".xml"
		if _, err := os.Stat(xmlFile); os.IsNotExist(err) {
			continue
		}
		if len(data) > 0 {
			continue
		}

		xmlData = readFileContents(xmlFile)
		getParsedMap(xmlData)

		/**
 		 * Complex header
 		 */
		selectXmlSection("hdr.p.r.pict.shape.textbox.txbxContent")
		data = parseWord1("Rev:", "p", "r")
		if check := hasData(data); check  {
			continue
		}

		/**
		 * Pulls Table
		 */
		selectXmlSection("hdr.tbl")
		data = parseWord1("Rev:", "tr", "tc")
		if check := hasData(data); check  {
			continue
		}
		data = parseWord1("Rev.", "tr", "tc")
		if check := hasData(data); check  {
			continue
		}

		/**
		 * Base header pullc
		 */
		selectXmlSection("hdr")
		data = parseWord1("Rev.", "p", "r")
		if check := hasData(data); check  {
			continue
		}
		data = parseWord1("Revision", "p", "r")
		if check := hasData(data); check  {
			continue
		}
		data = parseWord1("Rev", "p", "r")
		if check := hasData(data); check {
			continue
		}
	}
	if check := hasData(data); check {
		//fmt.Println(data)
		//fileWriter.WriteToFile(data, xmlPath)
		return data
	}

	/**
	 * =======================
	 * document.xml
	 * =======================
	 */

	xmlFile = xmlPath + "/word/document.xml"
	xmlData = readFileContents(xmlFile)
	getParsedMap(xmlData)
	selectXmlSection("document.body.tbl")
	data = parseWord3("Revision History", "tr", "tc")
	if check := hasData(data); check  {
		fileWriter.WriteToFile(data, xmlPath)
		return data
	}
	data = parseWord4("Rev:", "tr", "tc")
	if check := hasData(data); check  {
		fileWriter.WriteToFile(data, xmlPath)
		return data
	}
	data = parseWord4("Rev.", "tr", "tc")
	if check := hasData(data); check  {
			fileWriter.WriteToFile(data, xmlPath)
		return data
	}

//

	data = parseWord4("Revision", "p", "r")
	if check := hasData(data); check  {
		return data
	}
	data = parseWord4("Rev:", "p", "r")
	if check := hasData(data); check  {
		return data
	}
	data = parseWord4("Rev.", "p", "r")
	if check := hasData(data); check  {
		return data
	}
	data = parseWord4("Rev", "p", "r")
	if check := hasData(data); check  {
		return data
	}

	fileWriter.WriteLineToFile(xmlPath)

	return nil

}



func hasData(data [][]string) bool {
	found := false
	for _, d := range data {
		if(len(d) > 0){
			found = true
		}
	}
	return found
}
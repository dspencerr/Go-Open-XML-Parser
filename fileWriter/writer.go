package fileWriter

import (
	"io/ioutil"
	"os"
)

var TargetDir string
var MissedDocs int

func WriteToFile(dataToWrite [][]string, path string){
	return
	var stringData = path + "\n"
	for _, data := range dataToWrite {
		line := 0
		for _, d := range data{
			stringData += d
			if line < len(data) - 1 {
				stringData += " | "
			}
			line++
		}
		stringData += "\n"
	}

	stringData += "\n"
	file := TargetDir + "/data.txt"
	f, _ := os.OpenFile(file, os.O_APPEND, 0777)
	f.WriteString(stringData)
	f.Close()
}

func EmptyFile(){
	byteData := []byte("")
	file := TargetDir + "/data.txt"
	ioutil.WriteFile(file, byteData, 0777)
}

func WriteLineToFile(line string){
	MissedDocs++
	//return
	var stringData string
	stringData = "\n"
	stringData += line
	stringData += "\n"
	file := "C:\\Projects\\Go\\src\\github.com\\dspencerr\\docParser\\docTemp\\data.txt"
	//file := TargetDir + "/data.txt"
	f, _ := os.OpenFile(file, os.O_APPEND, 0777)
	f.WriteString(stringData)
	f.Close()

}
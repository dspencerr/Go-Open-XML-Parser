package fileMgr


import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"sync"
)

var wg sync.WaitGroup

func ParseFileByFile() {

	for x:=0; x<len(FilesToParse); x++{
		wg.Add(1)
		go ParseFileSafely(&FilesToParse[x])
	}
	wg.Wait();

	//writeDataToCSV()
	//DeleteTempDirectory()
}

func ParseFileSafely(file *DocFile) {
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			file.Status = "Go Routine Fail: "+file.Status
			fmt.Println("Go Routine Failed: ", err)
		}
	}()
	RenameAndUnzip(file)
}

func RenameAndUnzip(file *DocFile){

	file.Result = "empty"

	file.Status = "copying doc"
	file.copyDoc(xmlTempDir)

	file.Status = "making temp directory"
	file.makeDir(xmlTempDir)

	file.Status = "unziping folder"
	file.unzipToFolder()

	file.Status = "parsing XML"
	file.parseXmlDocument()

	file.Status = "deleting XML folder"
	//file.deleteXml()

	file.Status = "finished"
}

func writeDataToCSV(){
	var dataToWrite string
	for x, doc := range FilesToParse{
		if x == 0{
			dataToWrite += "DocName	DocPath	DocStatus	"+doc.Key+"\n"
		}
		name := strings.TrimSuffix(doc.Name, ".zip")
		dataToWrite += name+"\t"+doc.Path+"\t"+doc.Status+"\t"+doc.Result+"\n"
	}

	byteData := []byte(dataToWrite)
	//fmt.Println(byteData)

	file := targetDir + "/results.txt"
	//fmt.Println(file)

	ioutil.WriteFile(file, byteData, 0777)
}


func DeleteTempDirectory(){
	err := os.RemoveAll(xmlTempDir)
	if err != nil{
		fmt.Println(err);
	}
}
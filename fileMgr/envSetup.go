package fileMgr

import (
	"os"
	"os/user"
	"github.com/dspencerr/docParser/fileWriter"
)

var sourceDir string
var targetDir string
var xmlTempDir string

func SetupEnvironment(source_path string, target_path string) {
	if err := createSourcePath(source_path); len(err) > 0 {
		panic(err)
	}

	if err := createTargetPath(target_path); len(err) > 0 {
		panic(err)
	}

	if err := createTempXMLDir(); len(err) > 0{
		panic(err)
	}


	fileWriter.TargetDir = targetDir
	fileWriter.EmptyFile()
}

func createTempXMLDir() string {
	xmlTempDir = targetDir + "/xmlTemp"

	if info, err := os.Stat(xmlTempDir); err != nil || !info.IsDir() {
		if err := os.Mkdir(xmlTempDir, 0777); err != nil{
			return "Could not create the xml temp directory"
		}
		return ""
	}
	return ""
}

func createSourcePath(path string) string {
	if info, err := os.Stat(path); err != nil || !info.IsDir() {
		return "Source directory is not a valid directory"
	}
	sourceDir = string(path)
	return ""
}

func createTargetPath(path string) string {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		targetDir = string(path)
		return ""
	} else if usr, err  := user.Current(); err == nil{
		targetDir = string(usr.HomeDir) + "/Desktop"
		return ""
	} else {
		return "Taget directory is not a valid directory"
	}
}
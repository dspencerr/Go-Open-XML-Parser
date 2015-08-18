package fileMgr

import (
	"path/filepath"
	"strings"
	"os"
)

var FilesToParse = []DocFile{}

func WalkTheFileStructure() {

	FilesToParse = []DocFile{}

	err := filepath.Walk(sourceDir, walkFileWorker)
	if err != nil {
		panic("error traversing files")
	}
}

func walkFileWorker(path string, FileInfo os.FileInfo, err error) error {
	if strings.Contains(FileInfo.Name(), ".docx") || strings.Contains(FileInfo.Name(), ".xlsx") {
		file := DocFile{ Name:FileInfo.Name(), Path:path, Status:"unopened"}
		FilesToParse = append(FilesToParse, file)
	}
	return nil
}
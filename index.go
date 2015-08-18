package docParser

import (
	"github.com/dspencerr/docParser/fileMgr"
	"fmt"
	"github.com/dspencerr/docParser/fileWriter"
)


func ParseDocs(source string, target string) ([]fileMgr.DocFile, int) {

	source_path := source
	target_path := target

	fileMgr.SetupEnvironment(source_path, target_path)
	fileMgr.WalkTheFileStructure()
	fmt.Println(len(fileMgr.FilesToParse))
//	//fileMgr.UpdateProgress(progressBar)
	fileMgr.ParseFileByFile()
	fmt.Println("Missed Docs: ", fileWriter.MissedDocs)
	//fmt.Println(fileMgr.FilesToParse)
	return fileMgr.FilesToParse, len(fileMgr.FilesToParse)
}

package fileMgr

import (
	"os"
	"io"
	"strings"
	"github.com/dspencerr/docParser/xmlParser"
	"archive/zip"
)

type DocFile struct {
	Name, Path string
	Status, ZipPath, XmlPath, Type, Revision string
	Key, Result string
	AllData [][]string
}

func (d *DocFile) copyDoc(dir string){
	d.ZipPath = dir+"\\"+d.Name+".zip"
	if _, err := os.Stat(d.ZipPath); os.IsNotExist(err) {
		err := os.Link(d.Path, d.ZipPath)
		if err != nil{
			d.Status = "Broke on copy file"
			panic(err)
		}
	}
}

func (d *DocFile) makeDir(xmlTempDirectory string){
	if strings.Contains(d.Path, ".docx"){
		d.XmlPath = xmlTempDirectory +"/"+ strings.TrimSuffix(d.Name, ".docx")
		d.Type = "docx"
	}else if strings.Contains(d.Path, ".xlsx") {
		d.XmlPath = xmlTempDirectory +"/"+ strings.TrimSuffix(d.Name, ".xlsx")
		d.Type = "xlsx"
	}

	if len(d.XmlPath) > 0 {
		os.Mkdir(d.XmlPath, 0777)
		os.Mkdir(d.XmlPath+"/customXml", 0777)
		os.Mkdir(d.XmlPath+"/customXml/_rels", 0777)
		os.Mkdir(d.XmlPath+"/docProps", 0777)
		os.Mkdir(d.XmlPath+"/_rels", 0777)

		if strings.Contains(d.Path, ".docx"){
			os.Mkdir(d.XmlPath+"/word", 0777)
			os.Mkdir(d.XmlPath+"/word/_rels", 0777)
			os.Mkdir(d.XmlPath+"/word/theme", 0777)
		} else {
			os.Mkdir(d.XmlPath+"/xl", 0777)
			os.Mkdir(d.XmlPath+"/xl/_rels", 0777)
			os.Mkdir(d.XmlPath+"/xl/theme", 0777)
			os.Mkdir(d.XmlPath+"/xl/printerSettings", 0777)
			os.Mkdir(d.XmlPath+"/xl/worksheets", 0777)
			os.Mkdir(d.XmlPath+"/xl/worksheets/_rels", 0777)
		}
	}
}

func (d *DocFile) unzipToFolder() error {

	r, err := zip.OpenReader(d.ZipPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := d.XmlPath +"\\"+ f.Name

		if f.FileInfo().IsDir() {
			//os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DocFile) deleteXml(){
	os.Remove(d.ZipPath)
	os.RemoveAll(d.XmlPath)
}

func (d * DocFile) parseXmlDocument(){

	//fmt.Println(d)

	d.AllData = xmlParser.ParseXmlDocument(d.XmlPath, d.Type, targetDir)

	//d.Header, d.Result = xmlParser.ParseXmlDocument(d.XmlPath, targetDir)

//	d.AllData =
//	count := d.getHeaderCount(d.AllData)
//	d.getDataFromPattern(d.AllData, count)
}
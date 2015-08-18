package xmlParser

import (
	"github.com/dspencerr/uscore"
)

func parseWord3(needle string, arrayNode string, sectionNode string) [][]string {

	var tblResArray [][]string

	for _, tbl := range dataSelection {
		result := uscore.HasVal(tbl.(map[string]interface{}), needle)

		if !result {
			continue
		}

		rows := tbl.(map[string]interface {})["tr"].([]interface {})
		var cellsArray [][]string
		var ok bool

		for _, row := range rows {

			cellArr, arrOk := row.(map[string]interface {})["tc"].([]interface {})
			cellMap, mapOk := row.(map[string]interface {})["tc"].(map[string]interface {})

			if arrOk {
				cellsArray, ok = uscore.HasKey(cellArr, "", "", "t")
			}else if mapOk{
				cellsArray, ok = uscore.HasKey(cellMap, "", "", "t")
			} else{
				cellsArray, ok = [][]string{}, false
				continue;
			}

			if ok{
				for _, d := range cellsArray{
					tblResArray = append(tblResArray, d)
				}
				//tblResArray = cellsArray
			}
		}
	}
	return tblResArray
}
package tree

import (
	"github.com/bettersun/rain"
	"log"
	"testing"
)

func TestBuildTree(t *testing.T) {

	file1 := "data/data_1.json"
	file2 := "data/data_2.json"
	file3 := "data/data_3.json"
	file4 := "data/data_4.json"

	var data1 []Data
	err := rain.JsonFileToStruct(file1, &data1)
	if err != nil {
		log.Println(err)
	}

	var data2 []Data
	err = rain.JsonFileToStruct(file2, &data2)
	if err != nil {
		log.Println(err)
	}

	var data3 []Data
	err = rain.JsonFileToStruct(file3, &data3)
	if err != nil {
		log.Println(err)
	}

	var data4 []Data
	err = rain.JsonFileToStruct(file4, &data4)
	if err != nil {
		log.Println(err)
	}

	//log.Printf("%+v\\n", data1)
	//log.Printf("%+v\\n", data2)
	//log.Printf("%+v\\n", data3)
	//log.Printf("%+v\\n", data4)

	var allData [][]Data
	allData = append(allData, data1)
	allData = append(allData, data2)
	allData = append(allData, data3)
	allData = append(allData, data4)

	var root Node
	root.Id = "500"
	BuildTree(&root, allData, len(allData)-1)

	err = rain.OutJson("tree.json", root)
	if err != nil {
		log.Println(err)
	}
}

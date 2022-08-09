package src

import (
	// "fmt"
	"strings"
)

type DataStruct struct {
	Extension []string
	Regex     string
}

func getData() []DataStruct {
	var Data = []DataStruct{
		{
			Extension: []string{"py", "rb", "r", "sh"},
			// single line comment support for python,ruby,bash,r
			Regex: `#[^!].+`,
			// Regex:     `#[^!].+$|=begin[\s|\S]+=end$|<!--[\s|\S]-->$|:\'[\s|\S][^']+\'$`,
		},
		{
			Extension: []string{"dart", "c", "cpp", "h", "hpp", "java", "js", "ts", "kt", "go", "css"},
			// only single line comments
			Regex: `\/\/.+`,
			// Regex:     `\/\*[\s|\S]+\*\/|\/\/.+`,
		},
	}
	return Data
}

func GetRegex(filename string) (string,string) {

	fileNameList := strings.Split(filename,".")
	fileExtension:=fileNameList[len(fileNameList)-1]
	// fmt.Println(fileExtension)
	for _,v := range getData() {
		for _,val := range v.Extension{
			if val == fileExtension{
				return v.Regex,fileExtension
			}
		}
	}
	return "",fileExtension
}
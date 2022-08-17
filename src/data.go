package src

import (
	"strings"
)

type DataStruct struct {
	Extension []string
	Single     string
	MulB	   string
	MulE       string
}

func getData() []DataStruct {
	var Data = []DataStruct{
		{
			Extension: []string{"py", "r"},
			// Python and R support only single-line comments
			Single: `#[^!].+$|^'.+'$`,
			MulB: `\"\"\".*`,
			MulE: `.*\"\"\"`,
			// Regex:     `#[^!].+$|=begin[\s|\S]+=end$|<!--[\s|\S]-->$|:\'[\s|\S][^']+\'$`,
		},
		{
			Extension: []string{"rb","pl"},
			Single: `#[^!].+`,
			MulB: `=begin`,
			MulE: `=end|=cut`,
		},
		{
			Extension: []string{"dart", "c", "cpp", "h", "hpp", "java", "js", "ts", "kt", "go", "css"},
			// support for single-line as well as multi-line comments
			Single: `\/\/.+`,
			MulB: `\/\*.*`,
			MulE: `.*\*\/`,
		},
	}
	return Data
}

func GetRegex(filename string) (string,string,string) {

	fileNameList := strings.Split(filename,".")
	fileExtension:=fileNameList[len(fileNameList)-1]
	for _,v := range getData() {
		for _,val := range v.Extension{
			if val == fileExtension{
				return v.Single,v.MulB,v.MulE
			}
		}
	}
	return "","",""
}
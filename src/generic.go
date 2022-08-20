package src

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//checks error
func ErrorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//checks whether a list of strings contains a string
func Contains(a string, list []string)bool{
	for _,val := range list {
		if val == a {
			return true
		}
	}
	return false
}

//check file path is correct or not
func CheckFileExists(filename string) bool {
	// loc, err := os.Getwd()
	// ErrorCheck(err)
	// filePath := loc+"/"+filename
	
	_, error := os.Stat(filename)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func ProcessLineNotStartsWithCmt(str string,cmtChar string,regex string,f *os.File,addLineNumber bool,deleteCmtChars bool,lineNumber int) {
	if addLineNumber{
		f.WriteString(strconv.Itoa(lineNumber)+" ")
	}
	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(str)
	if deleteCmtChars{
		finalString := strings.ReplaceAll(matches[0],cmtChar,"")
		f.WriteString(finalString+"\n")
	} else {
		f.WriteString(matches[0]+"\n")
	}
}

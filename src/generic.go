package src

import (
	"errors"
	"log"
	"os"
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
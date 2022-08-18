package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "time"
	// "github.com/nathan-fiscaletti/consolesize-go"
)

func getFiles(files []string) []string {
	if (files[0] == "."){
		//TODO: functionality to get the files in current location
	}
	return files
}

func Run(filesFromUser []string,outfile string,extract bool,addLineNumber bool,deleteCmtChars bool){
	// files := []string{"main.dart","widgets.dart","themes.dart","services.dart","homescreen.dart","splashscreen.dart"}

	files := getFiles(filesFromUser)
	if extract{
		fmt.Println("Extracting ...")
	} else {
		fmt.Println("Processing ...")
	}

	if CheckFileExists("outfile.txt") && outfile != "outfile.txt" {
		e:= os.Remove("outfile.txt")
		ErrorCheck(e)
	}
	if CheckFileExists(outfile){
		err:= os.Remove(outfile)
		ErrorCheck(err)
	} 
	f,err :=os.Create(outfile)
	ErrorCheck(err)
	defer f.Close()

	// TODO: Test this line of code in production

	// fmt.Print("\033[1A\033[K")
	for _,file := range files{
		fmt.Print(file)
		fileProcessing(file,f,addLineNumber,deleteCmtChars)
		// This line erases the recently printed line from terminal
		fmt.Println(" ✓")
	}

	// n,_ := consolesize.GetConsoleSize()
	// fmt.Println(n)
	// str := "#"
	// fmt.Println("Processing...")
	// for i:=0;i<n;i++{
	// 	fmt.Println(str)
	// 	time.Sleep(300 * time.Millisecond)
	// 	fmt.Print("\033[1A\033[K")
	// 	str = str + "#"
	// }
	//this code is to clear the latest println
	// fmt.Print("\033[1A\033[K")

	fmt.Println("Process Completed..!")
	fmt.Print("You can find the extracted comments in the file ",outfile,"\n")
}



func fileProcessing(filename string,f *os.File,addLineNumber bool,deleteCmtChars bool){

	isMultilineComment := false

	single,mulb,mule:= GetRegex(filename)
	if (single == "") {
		fmt.Printf("file type of \"%v\" is not supported\n",filename)
		os.Exit(0)
	}
	cmtChars := GetCommentCharacters(filename)

	file,err :=os.Open(filename)
	ErrorCheck(err)
	defer file.Close()

	f.WriteString(filename+"\n\n")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lineNumber := 1
	for scanner.Scan(){
		matchSingleLine,err:=regexp.Match(single,[]byte(scanner.Text()))
		ErrorCheck(err)
		if matchSingleLine{
			if addLineNumber{
				f.WriteString(strconv.Itoa(lineNumber)+" ")
			}
			if deleteCmtChars{
				finalString := removeCommentCharacters(scanner.Text(),cmtChars)
				f.WriteString(finalString+"\n")
			} else {
				f.WriteString(scanner.Text()+"\n")
			}
			lineNumber += 1
			continue
		}
		matchMullineBegin,err:=regexp.Match(mulb,[]byte(scanner.Text()))
		ErrorCheck(err)
		if matchMullineBegin && isMultilineComment==false {
			if addLineNumber{
				f.WriteString(strconv.Itoa(lineNumber)+" ")
			}
			if deleteCmtChars {
				finalString := removeCommentCharacters(scanner.Text(),cmtChars)
				f.WriteString(finalString+"\n")
			} else {
				f.WriteString(scanner.Text()+"\n")
			}
			isMultilineComment = true
			lineNumber += 1
			continue
		}
		matchMullineEnd,err:=regexp.Match(mule,[]byte(scanner.Text()))
		ErrorCheck(err)
		if matchMullineEnd && isMultilineComment==true {
			if addLineNumber{
				f.WriteString(strconv.Itoa(lineNumber)+" ")
			}
			if deleteCmtChars {
				finalString := removeCommentCharacters(scanner.Text(),cmtChars)
				f.WriteString(finalString+"\n")
			} else {
				f.WriteString(scanner.Text()+"\n")
			}
			isMultilineComment = false
			lineNumber += 1
			continue
		}
		if isMultilineComment {
			if addLineNumber{
				f.WriteString(strconv.Itoa(lineNumber)+" ")
			}
			if deleteCmtChars {
				finalString := removeCommentCharacters(scanner.Text(),cmtChars)
				f.WriteString(finalString+"\n")
			} else {
				f.WriteString(scanner.Text()+"\n")
			}
			lineNumber += 1
			continue
		}
		lineNumber += 1
	}
	f.WriteString("\n")
}

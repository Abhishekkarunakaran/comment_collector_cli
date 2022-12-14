package src

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func getFiles(files []string) []string {
	var outFiles []string
	curDir, err := os.Getwd()
	ErrorCheck(err)

	//! if the user need a the files inside currentDirectory
	if (files[0] == "."){
		

		er := filepath.Walk(curDir,func(path string,info os.FileInfo,err error) error{
			if err != nil{
				fmt.Println(err)
				return err
			}
			ignore,errFromRegex:=regexp.Match(`.*.git.*`,[]byte(path))
			ErrorCheck(errFromRegex)


			if (!info.IsDir() && !ignore && SupportedFiles(path)){
				// fmt.Println(path)
				outFiles = append(outFiles,path)
				
			}
			return nil
		})
		ErrorCheck(er)
	
		return outFiles
	}

	for _,file := range(files){

		outFiles = append(outFiles,curDir+string(os.PathSeparator)+file)

	}
	fmt.Println(outFiles)
	return outFiles
}

func Run(filesFromUser []string,outfile string,extract bool,addLineNumber bool,deleteCmtChars bool){
	// files := []string{"main.dart","widgets.dart","themes.dart","services.dart","homescreen.dart","splashscreen.dart"}

	files := getFiles(filesFromUser)
	if extract{
		fmt.Println("Extracting ...")
	} else {
		fmt.Println("Processing ...")
	}

	if CheckFileExists("outfile.md") && outfile != "outfile.md" {
		e:= os.Remove("outfile.md")
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

	curDir, err := os.Getwd()
	ErrorCheck(err)
	currentDirectoryList := strings.Split(curDir,string(os.PathSeparator)) 
	currentDirectory := currentDirectoryList[len(currentDirectoryList)-1]
	for _,file := range files{
		foldFileList := strings.Split(file,currentDirectory+string(os.PathSeparator))
		foldFile := foldFileList[len(foldFileList)-1]
		fmt.Print(foldFile)
		fileProcessing(file,foldFile,f,addLineNumber,deleteCmtChars)
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



func fileProcessing(filename string,filenameshort string,f *os.File,addLineNumber bool,deleteCmtChars bool){

	isMultilineComment := false

	single,mulb,mule:= GetRegex(filename)
	if (single == "") {
		fmt.Printf("\nfile type of \"%v\" is not supported\n",filename)
		// os.Exit(0)
	}
	cmtChars := GetCommentCharacters(filename)

	file,err :=os.Open(filename)
	ErrorCheck(err)
	defer file.Close()



	f.WriteString("## "+filenameshort+"\n\n")
	f.WriteString("```\n")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lineNumber := 1
	for scanner.Scan(){
		matchSingleLine,err:=regexp.Match(single,[]byte(scanner.Text()))
		ErrorCheck(err)
				
		if matchSingleLine{
			firstChar := string(scanner.Text()[0])
			
			//? This could go wrong (works perfectly for python)
			if firstChar != string(string(cmtChars[0])[0]) && firstChar != `'`{
				ProcessLineNotStartsWithCmt(scanner.Text(),cmtChars[0],single,f,addLineNumber,deleteCmtChars,lineNumber)
				continue
			}
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
		if matchMullineBegin && !isMultilineComment {
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
		if matchMullineEnd && isMultilineComment {
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

	f.WriteString("```\n")
}

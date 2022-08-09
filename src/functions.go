package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// "time"
	// "github.com/nathan-fiscaletti/consolesize-go"
)

func Run(files []string,outfile string,extract bool){
	// files := []string{"main.dart","widgets.dart","themes.dart","services.dart","homescreen.dart","splashscreen.dart"}
	if extract{
		fmt.Println("Extracing...")
	} else{
		fmt.Println("Processing...")
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

	for _,file := range files{
		fmt.Println(file)
		fileProcessing(file,f)
		// time.Sleep(600 * time.Millisecond)
		fmt.Print("\033[1A\033[K")
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
	fmt.Print("\033[1A\033[K")
	fmt.Println("Process Completed..!")
	fmt.Print("You can find the extracted comments in the file ",outfile,"\n")
}



func fileProcessing(filename string,f *os.File){

	regex,ext:= GetRegex(filename)
	if (regex == "") {
		fmt.Printf("extension \"%v\" is not supported\n",ext)
		os.Exit(0)
	}
	// re,_ :=regexp.Compile(regex)


	file,err :=os.Open(filename)
	ErrorCheck(err)
	defer file.Close()

	f.WriteString(filename+"\n\n")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		match,err:=regexp.Match(regex,[]byte(scanner.Text()))
		ErrorCheck(err)
		if match{
			f.WriteString(scanner.Text()+"\n")
		}
	}
	f.WriteString("\n")
}

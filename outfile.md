## /cltr.go

```
29 //the arguments should be seperated by commas
29 //etc. cltr run -s main.c,file.cpp 
53 // Action upon the run command
```
## /sample/main.dart

```
1 // This is a single-line comment
2 /* This is a multiline comment
3 khkasdf */
5 /*
6 lkjlkdjlgjasdga
7 */
10 //The comment with the code
12 // This is the end of the line
```
## /sample/main.py

```
1 # This is a comment
2 """
3 This is a multiline documentation comment
4 """
5 'This is a single line docstring'
8 #The comment with the code
12 #end of the file
```
## /src/data.go

```
19 // Python and R support only single-line comments
24 // Regex:     `#[^!].+$|=begin[\s|\S]+=end$|<!--[\s|\S]-->$|:\'[\s|\S][^']+\'$`,
34 // support for single-line as well as multi-line comments
37 //","/*","*/"},
42 /* 
43 Multiline comment for test
44 */
```
## /src/functions.go

```
18 //! if the user need a the files inside currentDirectory
31 // fmt.Println(path)
49 // files := []string{"main.dart","widgets.dart","themes.dart","services.dart","homescreen.dart","splashscreen.dart"}
69 // TODO: Test this line of code in production
69 // fmt.Print("\033[1A\033[K")
79 // This line erases the recently printed line from terminal
82 // n,_ := consolesize.GetConsoleSize()
82 // fmt.Println(n)
82 // str := "#"
82 // fmt.Println("Processing...")
82 // for i:=0;i<n;i++{
82 // 	fmt.Println(str)
82 // 	time.Sleep(300 * time.Millisecond)
82 // 	fmt.Print("\033[1A\033[K")
82 // 	str = str + "#"
82 // }
82 //this code is to clear the latest println
82 // fmt.Print("\033[1A\033[K")
96 // os.Exit(0)
117 //? This could go wrong (works perfectly for python)
```
## /src/generic.go

```
12 //checks error
19 //checks whether a list of strings contains a string
29 //check file path is correct or not
31 // loc, err := os.Getwd()
31 // ErrorCheck(err)
31 // filePath := loc+"/"+filename
33 //return !os.IsNotExist(err)
```
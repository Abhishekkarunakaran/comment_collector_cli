# comment_collector_cli

## This cli collects the comments from your project files

_work in progress_ :construction:

### Aug 15 : New updation
> Added the support for single-line & multiline comments in following languages:</br>
> __C, C++, Java, Dart, JavaScript, TypeScript, Kotlin, Golang, CSS,
> Python, Ruby, R, Perl.__

### Aug 18:
> 1. Added feature for extracting the line number of the comments from the source flie(s) ( _using the flag -ln_ )
> 2. Added the support for single-line as well as multi-line documentation comments in python
> 3. Added the support for removing the comment characters from the outfile ( _using the flag -rc_ )

### Aug 21:
> Fixed forgotten case.
> Extraction of comment from a line of code
> eg:
> ```python
> def main(): #This is the main fuction
>   print(something)
> ```
> New version supports the extract of theline "`#This is the main function`"
> Issue reported by [Ananthu B](https://github.com/AnanthuB2001)

### Sept 7:
> The binary can run in any directory.</br>
>**cltr** finds the location of working directory and work with the sourcecode files in the subdirectories.
> #### bug:</br>
> regex also matches single-line comment character ( // ) in a string. The rest of the line including the comment character is extracted from the file.</br>
Go through the file: <u></br>*outfile.md*  `src/data.go - line.no.37`</br></u>
[check here](https://github.com/Abhishekkarunakaran/comment_collector_cli/tree/main/sample/outfile.txt)

### Sept 21:
> When using '.' in the command `cltr run -s .` to indicate the whole files in the directory, there was this issue of the path separator.The windows and Linux-Unix uses different file separators. The code that handles the whole files in the working directory was written using a Linux Machine (Fedora). When the code executed in Windows (10), It prints the whole file location.</br></br>
> It is corrected using passing `os.PathSeparator` instead of the hardcoding the path separator character for each operating systems.
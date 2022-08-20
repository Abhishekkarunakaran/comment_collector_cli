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
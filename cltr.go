package main

import (
	// "flag"
	// "cltr/src"
	// "fmt"
	"cltr/src"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func info() {
	app.Name = "Comment Collector CLI App"
	app.Usage = "To extract comments from programs to a new file"
	app.Authors = []*cli.Author{{Name:"Abhishek K K", Email: "me.abhishek.2k@gmail.com"}}
	
}

func commands(){
	app.Commands = []*cli.Command{
		{
			Name: "run",
			Aliases: []string{"r,exec"},
			Usage: "Extract the comments from the file(s)",
			Description: "Copy the comments from the file(s) to outfile.txt",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					//the arguments should be seperated by commas
					//etc. cltr run -s main.c,file.cpp 
					Required: true,
					Name: "source",
					Aliases: []string{"s"},
				},
				&cli.StringFlag{
					Name: "output",
					Aliases: []string{"o"},
				},
				&cli.BoolFlag{
					Name: "extract",
					Aliases: []string{"e"},
					Value: false,
				},
				&cli.BoolFlag{
					Name: "lineNumber",
					Aliases: []string{"ln"},
					Value: false,
				},
				&cli.BoolFlag{
					Name: "removeCommentCharacters",
					Aliases: []string{"rc"},
					Value: false,
				},
			},
			// Action upon the run command
			Action: func (c *cli.Context) error {

				outfile:="outfile.txt"
				
				if (len(c.StringSlice("source"))==0){
					fmt.Println("source file should be defined")
					os.Exit(0)
				}
				if (len(c.String("output"))!=0) {
					outfile = c.String("output")
				}
				 
				src.Run(c.StringSlice("source"),outfile,c.Bool("extract"),c.Bool("lineNumber"),c.Bool("removeCommentCharacters"))

				return nil				
			},
		},

	}
}


func main()  {

	info()
	commands()
	app.Run(os.Args)






	// var command string
	// commands :=[]string{"run","help","version"}

	// if len(os.Args) == 1{ 
	// 	//if the is no command
	// 	fmt.Println("help")
	// 	os.Exit(0)
	// } else{
	// 	//checking whether it is a valid command
	// 	if src.Contains(os.Args[1],commands){ 
	// 		command = os.Args[1]
	// 	} else {
	// 		fmt.Println("Wrong usage")
	// 		fmt.Println("help")
	// 	}
	// }
	
	// switch command{
	// case "run":
	// 	src.Run()
	// 	fmt.Println("run")
	// case "help":
	// 	src.PrintHelp()
	// 	fmt.Println("help")
	// }
}
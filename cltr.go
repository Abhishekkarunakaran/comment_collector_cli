package main

import (
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

}
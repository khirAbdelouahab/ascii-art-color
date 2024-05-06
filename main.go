package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	//"banner/colors"
	"banner/banner"
)

func ShowMessageError() {
	fmt.Printf("  Usage: go run . [OPTION] [STRING]\n\n")
	fmt.Println("  EX: go run . --color=<color> <letters to be colored> \"something\"")
}
func IsFile(fileName string) bool {
	if filepath.Ext(fileName) != ".txt" {
		fileName+= ".txt"
	}
	return fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt" 
}
func Ascii_Arts_Generator(arguments []string) {
	argument_Size := len(arguments)
	fileName := "standard"
	switch argument_Size {
		case 4:
			if len(arguments[0]) > 8 && arguments[0][:8] == "--color=" {
				fileName = arguments[3]
				Record := arguments[2]
				color := arguments[0][8:]
				newLineCounter := strings.Count(fileName, "\\n")
				words := ManipulateData(Record,fileName)
				banner.Result(words,newLineCounter,GetDataFromFile(fileName),color,arguments[1])
			} else {
				ShowMessageError()
			}	 
			break
		case 3 :
			if len(arguments[0]) > 8 && arguments[0][:8] == "--color=" {
				color := arguments[0][8:]
				word := ""
				letters := ""
				if IsFile(arguments[2]) {
					fileName = arguments[2]
					word = arguments[1]
				} else {
					word = arguments[2]
					letters = arguments[1]
				}
				banner.Result(ManipulateData(word,fileName),strings.Count(word, "\\n"),GetDataFromFile(fileName),color,letters)
			} else {ShowMessageError()}
		case 2 :
			if IsFile(arguments[1]) {
				banner.Result(ManipulateData(arguments[0],arguments[1]),strings.Count(arguments[0], "\\n"),GetDataFromFile(arguments[1]),"","")
			} else if arguments[0][:8] == "--color=" {
				banner.Result(ManipulateData(arguments[1],fileName),strings.Count(arguments[1], "\\n"),GetDataFromFile(fileName),arguments[0][8:],"")
			} else {
				ShowMessageError()
			}
		case 1 :
			banner.Result(ManipulateData(arguments[0],fileName),strings.Count(arguments[0], "\\n"),GetDataFromFile(fileName),"","")
		default :
			ShowMessageError()
	}
}
func GetDataFromFile(fileName string) map[int][]string{
	if filepath.Ext(fileName) !=".txt" {
		fileName+= ".txt"
	}
	return banner.ReadBannerFiles(fileName)
}
func ManipulateData(arg ,fileName string) [] string {
		words := strings.Split(arg, "\\n")
		if !banner.CheckIfAllCharInFile(words) {
			
			fmt.Println("You Have a Character Not found in the file >>", fileName)
			return nil
		}
		return words
}
func main() {
	arg := os.Args
	
	Ascii_Arts_Generator(arg[1:])
	// --color=red "hey" "hey how are you" shadow  4
	// --color=red "hello" tinkthory 3
	// --color=red "i" "i like you" 3
	// --color=red "hello world" 2
	// "hello" 1
	// "hello" shadow 2
	//
	
}

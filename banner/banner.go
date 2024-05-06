package banner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SetColor(colorName string) string {
	switch colorName {
		case "red":
			return "\033[31m"	
		case "green":
			return "\033[32m"	
		case "yellow":
			return "\033[33m"
		case "blue":
			return "\033[34m"
		case "magenda":
			return "\033[35m"
		case "cyan":
			return "\033[36m"
		case "gray":
			return "\033[37m"
		case "white":
			return "\033[97m"
		default :
			return "\033[0m"
	}
}
// Read banner files and store the data in a map
func ReadBannerFiles(txt string) map[int][]string {
	DATA := make(map[int][]string)
	file, err := os.Open(txt)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	read := bufio.NewScanner(file)
	ascciStart := 31
	for read.Scan() {
		if read.Text() == "" {
			ascciStart++
		} else {
			DATA[ascciStart] = append(DATA[ascciStart], read.Text())
		}
	}
	return DATA
}

// Check if all characters in the input are within printable ASCII range
func CheckIfAllCharInFile(words []string) bool {
	Temp := strings.Join(words, "")
	for _, char := range Temp {
		if char < ' ' || char > '~' {
			return false
		}
	}
	return true
}

func PointsToBeColored(word , LettersToBeColored string) [][]int {
	size_Word := len(word)
	Points := make([][]int,0) 
	sizeOfLetters := len(LettersToBeColored) // "hello world" "he"
	for i := 0; i <= size_Word - 1; i++ {
		if i <= size_Word - sizeOfLetters && word[i:sizeOfLetters + i] == LettersToBeColored {
			Points = append(Points, []int{i,i + sizeOfLetters - 1}) 
		}
	}
	return Points
}
func PrintChars(word , LettersColored , color string, banner map[int][]string) {
	Points := PointsToBeColored(word,LettersColored)
	Size_Points := len(Points)
	activeIndex := 0
	wordSize := len(word)
	for i := 0; i < 8; i++ {
		for j := 0 ; j <= wordSize - 1; j++ {
				if Size_Points > activeIndex && j >= Points[activeIndex][0] && j <= Points[activeIndex][1] {
					print(SetColor(color))
					if j == Points[activeIndex][1] {
						activeIndex++
					}
				} else {
					print(SetColor("reset"))
				}
			if LettersColored == "" {
				print(SetColor(color))
			}
			if j == len(word) - 1 {                   // [8,12] "ooo" "oooo" (0,2) (1,3) || (0,0) (1,1)
				fmt.Println(banner[int(word[j])][i]) // "you" "yow are\nyou you" s[i:len(word)+i]
				continue
			}
			fmt.Print(banner[int(word[j])][i])
		}
		activeIndex = 0
	}
}
// Generate and print the result
func Result(words []string, newLineCounter int, banner map[int][]string, color, Letters string) {
	counter := 1
	for _, word := range words {
		if word == "" && counter <= newLineCounter {
			fmt.Println()
			counter++
			continue
		}
		PrintChars(word,Letters,color,banner)
	}
	print(SetColor("reset"))

}

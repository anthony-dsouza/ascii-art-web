package main

import (
	"bufio"

	"fmt"
	"io"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// getting font
	font, err := os.Open("fonts/standard.txt")
	Check(err)

	// output := flag.String("output", "", "[STRING] [BANNER] [OPTION]")
	// flag.Parse()
	// scanning by lines
	NewData := bufio.NewScanner(io.Reader(font))
	// Appending lines to Arr
	var Arr []string
	for NewData.Scan() {
		Arr = append(Arr, NewData.Text())
	}

	// counts lines and resets after reaching 9
	count := 0
	key := make(map[int][]int)
	// contains key value which is the decimal value for the character
	// this value is incremented once we are finished appending line numbers
	// to key
	runeVal := 32

	for i := range Arr {
		if count%9 == 0 && count != 0 {
			for j := i - 8; j < i; j++ {
				key[runeVal] = append(key[runeVal], j)
			}
			runeVal++
		}

		count++

	}

	str := os.Args[1]

	// converted to list of strings to change individual values
	var list []string
	word := ""

	for i, char := range str {
		if char == '\\' && rune(str[i+1]) == 'n' {

		} else if char == 'n' && rune(str[i-1]) == '\\' {
			if word != "" { // stop unnessary "" being added - would create an additional newline
				list = append(list, word)
				word = ""
			}

		} else {
			word = word + string(char)
		}

		if i == len(str)-1 {
			list = append(list, word)
		}
	}

	var banner []string 

	for _, word:= range list {
		banner = word + list + ""
	}

	for _, word := range list {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				line := ""
				for _, r := range word {
					line = line + banner[r][i]
				}
				fmt.Println(line)
			}
		}

	}

	// fmt.Println(*output)
}

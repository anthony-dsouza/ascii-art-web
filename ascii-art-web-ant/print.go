package ascii

import "fmt"

func Print(str string, banner map[rune][]string) {

	list := SplitByNewLine(str)

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
}

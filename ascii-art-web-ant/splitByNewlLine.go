package ascii

func SplitByNewLine(str string) []string {
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

	return list
}

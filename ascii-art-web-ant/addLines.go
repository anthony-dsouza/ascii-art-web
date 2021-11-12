package ascii

func AddLines(l []string, b map[rune][]string) string {

	str := ""
	for _, word := range l {
		if word == "" {
			str = str + "\n"
		} else {
			for i := 0; i < 8; i++ {
				line := ""
				for _, r := range word {
					line = line + b[r][i]
				}
				str = str + line + "\n"
			}
		}
	}
	return str
}

package main

import (
	"ascii"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Body   []byte
	Banner []string
}

func asciiArt(input string, ban string) (str string, bMap map[rune][]string) {
	data := ascii.Banner(ban)
	output := input

	// outputFile := flag.String("output", os.Args[2][9:], "output into file")

	// flag.Parse()
	// fmt.Println(*outputFile)
	// fmt.Println(flag.Parsed())

	defer data.Close()

	ArrayOfLines := ascii.Array(data)

	bannerMap := ascii.Map(ArrayOfLines)

	return output, bannerMap
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	output, bannerMap := asciiArt(r.FormValue("input"), r.FormValue("banner"))
	list := ascii.SplitByNewLine(output)

	str := ""
	for _, word := range list {
		if word == "" {
			str = str + "\n"
		} else {
			for i := 0; i < 8; i++ {
				line := ""
				for _, r := range word {
					line = line + bannerMap[r][i]
				}
				str = str + line + "\n"
			}
		}

	}
	p1 := &Page{Body: []byte(str)}
	fonts, _ := os.ReadDir("fonts")
	for _, name := range fonts {
		p1.Banner = append(p1.Banner, name.Name())
	}

	// p1.save()
	// p := loadPage("ascii.txt")

	t, _ := template.ParseFiles("ascii.html")
	t.Execute(w, p1)

}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting..")
	http.ListenAndServe(":3000", nil)

}

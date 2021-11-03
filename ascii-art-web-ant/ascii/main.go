package main

import (
	"ascii"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Body []byte
}

func (p *Page) save() error {
	filename := "ascii.txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := "ascii.txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Body: body}
}

func asciiArt() (str string, bMap map[rune][]string) {
	data := ascii.Banner()
	output := os.Args[1]

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
	p := loadPage("ascii.txt")

	t, _ := template.ParseFiles("ascii.html")
	t.Execute(w, p)

}

func main() {
	output, bannerMap := asciiArt()
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
	p1.save()
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting..")
	http.ListenAndServe(":3000", nil)
}

package main

import (
	"ascii"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Body   []byte
	Banner []string
	Input  string
	Font   string
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

func handlerGet(w http.ResponseWriter, r *http.Request) {
	input := "Welcome"
	font := "standard.txt"
	p1 := &Page{Input: input}
	output, bannerMap := asciiArt(input, font)
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
	p1.Body = []byte(str)
	p1.Font = font
	fonts, _ := os.ReadDir("fonts")
	for _, name := range fonts {
		p1.Banner = append(p1.Banner, name.Name())
	}

	t, _ := template.ParseFiles("ascii.html")
	t.Execute(w, p1)

}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	font := r.FormValue("banner")
	output, bannerMap := asciiArt(input, font)
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
	p1.Input = input
	p1.Font = font
	fonts, _ := os.ReadDir("fonts")
	for _, name := range fonts {
		p1.Banner = append(p1.Banner, name.Name())
	}

	t, _ := template.ParseFiles("ascii.html")
	t.Execute(w, p1)

}

func main() {

	http.HandleFunc("/", handlerGet)
	http.HandleFunc("/ascii-art", handlerPost)
	fmt.Println("starting..")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

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

	defer data.Close()

	ArrayOfLines := ascii.Array(data)

	bannerMap := ascii.Map(ArrayOfLines)

	return output, bannerMap
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
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

	t, err := template.ParseFiles("ascii.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, p1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

const (
	StatusOK                  = 200 // RFC 7231, 6.3.1
	StatusBadRequest          = 400 // RFC 7231, 6.5.1
	StatusInternalServerError = 500 // RFC 7231, 6.6.1
	StatusNotFound            = 404 // RFC 7231, 6.5.4
)

var statusText = map[int]string{
	StatusOK:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	input := r.FormValue("input")

	font := r.FormValue("banner")

	output, bannerMap := asciiArt(input, font)

	// removing multilines and replacing with \n

	sOutput := ""

	for _, char := range output {
		if char == 13 {

		} else if char == 10 {
			sOutput = sOutput + "\\" + "n"
		} else {
			sOutput = sOutput + string(char)
		}
	}

	list := ascii.SplitByNewLine(sOutput)

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
func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

	return

}

func main() {
	// handler := http.StatusText(handleRequest)
	// http.Handle("/ascii", handler)

	http.HandleFunc("/", handlerGet)
	http.HandleFunc("/ascii-art", handlerPost)
	fs := http.FileServer(http.Dir("stylesheets/"))
	http.Handle("/stylesheets/",
		http.StripPrefix("/stylesheets/", fs))
	fmt.Println("starting..")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

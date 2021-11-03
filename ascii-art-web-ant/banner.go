package ascii

import "os"

func Banner() *os.File {

	b, err := os.Open("standard.txt")

	if err != nil {
		panic(err)
	}

	return b

}

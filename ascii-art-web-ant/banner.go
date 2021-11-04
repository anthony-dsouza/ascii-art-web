package ascii

import "os"

func Banner(str string) *os.File {

	b, err := os.Open("fonts/" + str)

	if err != nil {
		panic(err)
	}

	return b

}

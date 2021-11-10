package ascii

import "os"

func Banner(str string) (*os.File, error) {
	b, err := os.Open("fonts/" + str)

	return b, err
}

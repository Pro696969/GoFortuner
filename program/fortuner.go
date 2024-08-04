package fortuner

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

// var f *os.File
var f embed.FS

func err_handler(e error) {
	if e != nil {
		panic(e)
	}
}

func random_intn(min, max int) int {
	return min + rand.Intn(max-min)
}

func random_fortuner_allfortunes() *os.File {
	random_num := random_intn(0, 3)
	files := [3]string{"fortunes", "literature", "riddles"}
	choice := files[random_num]
	fname := "assets/" + choice
	f, err := os.Open(fname)
	err_handler(err)
	return f
}

func random_fortuner_literature() *os.File {
	f, err := os.Open("assets/literature")
	err_handler(err)
	return f
}

func random_fortuner_riddles() *os.File {
	f, err := os.Open("assets/riddles")
	err_handler(err)
	return f
}

func random_fortuner_fortunes() *os.File {
	f, err := os.Open("assets/fortunes")
	err_handler(err)
	return f
}

func random_fortune_printer(f *os.File) {
	b, err := io.ReadAll(f)
	err_handler(err)
	b_string := string(b)
	fortuner_array := strings.Split(b_string, "%")
	random_fortune := random_intn(1, len(fortuner_array))
	fmt.Println(fortuner_array[random_fortune])
}

func Program() {
	cli_args := os.Args
	// fmt.Println(len(cli_args))
	if len(cli_args) == 1 {
		f := random_fortuner_allfortunes()
		random_fortune_printer(f)
	} else {
		switch cli_args[1] {
		case "-f":
			f := random_fortuner_fortunes()
			random_fortune_printer(f)
		case "-l":
			f := random_fortuner_literature()
			random_fortune_printer(f)
		case "-r":
			f := random_fortuner_riddles()
			random_fortune_printer(f)
		default:
			fmt.Println("incorrect usage, try :")
			fmt.Println("fortune [-f/-r/-l]")
		}
	}
	// defer f.Close()
}

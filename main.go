package main

import (
	"embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//go:embed assets/fortunes
var f_fortunes embed.FS

//go:embed assets/literature
var f_lit embed.FS

//go:embed assets/riddles
var f_riddles embed.FS

func err_handler(e error) {
	if e != nil {
		panic(e)
	}
}

func random_intn(min, max int) int {
	return min + rand.Intn(max-min)
}

func random_fortuner_allfortunes() []byte {
	files := [3]string{"fortunes", "literature", "riddles"}
	random_num := random_intn(0, 3)
	choice := files[random_num]
	fname := "assets/" + choice
	if fname == "assets/fortunes" {
		data, err := f_fortunes.ReadFile(fname)
		err_handler(err)
		return data
	} else if fname == "assets/literature" {
		data, err := f_lit.ReadFile(fname)
		err_handler(err)
		return data
	} else {
		data, err := f_riddles.ReadFile(fname)
		err_handler(err)
		return data
	}
}

func random_fortuner_literature() []byte {
	data, err := f_lit.ReadFile("assets/literature")
	err_handler(err)
	return data
}

func random_fortuner_riddles() []byte {
	data, err := f_riddles.ReadFile("assets/riddles")
	err_handler(err)
	return data
}

func random_fortuner_fortunes() []byte {
	data, err := f_fortunes.ReadFile("assets/fortunes")
	err_handler(err)
	return data
}

func random_fortune_printer(f []byte) {
	b_string := string(f)
	fortuner_array := strings.Split(b_string, "%")
	random_fortune := random_intn(1, len(fortuner_array))
	fmt.Println(fortuner_array[random_fortune])
}

func main() {
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
}

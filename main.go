package main

import (
	"embed"
	"fmt"
	"math/rand"
)

var data []byte

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
		data, _ = f_fortunes.ReadFile(fname)
	} else if fname == "assets/literature" {
		data, _ = f_lit.ReadFile(fname)
	} else {
		data, _ = f_riddles.ReadFile(fname)
	}
	// data, _ := f_fortunes.ReadFile(fname)
	fmt.Println(string(data))
	return data
}

func main() {
	random_fortuner_allfortunes()
}

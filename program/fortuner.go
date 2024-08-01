package fortuner

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

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
	// fmt.Println(choice)
	fname := "assets/" + choice
	// fmt.Println(fname)
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

	// fmt.Println(b_string)

	fortuner_array := strings.Split(b_string, "%")
	random_fortune := random_intn(1, len(fortuner_array))

	fmt.Println(fortuner_array[random_fortune])
}

func Program() {
	// cli_args := os.Args

	// fmt.Println(string(f))
	// scanner := bufio.NewScanner(f)
	// scanner.Split(bufio.ScanLines)
	//
	// count := 0
	// for scanner.Scan() {
	//   count++
	// }
	//
	// fmt.Println("number of lines scanned ", count)

	// f := random_fortuner_allfortunes()
	f := random_fortuner_fortunes()
	defer f.Close()
	random_fortune_printer(f)
	// b, err := io.ReadAll(f)
	// err_handler(err)
	// b_string := string(b)

	// // fmt.Println(b_string)

	// fortuner_array := strings.Split(b_string, "%")
	// random_fortune := random_intn(1, len(fortuner_array))

	// fmt.Println(fortuner_array[random_fortune])

}

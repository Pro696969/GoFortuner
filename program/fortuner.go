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

func random_fortuner_all() *os.File {
	f, err := os.Open("assets/all_fortunes")
	err_handler(err)
	return f
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

	f := random_fortuner_all()
	defer f.Close()
	b, err := io.ReadAll(f)
	err_handler(err)
	b_string := string(b)

	// fmt.Println(b_string)

	fortuner_array := strings.Split(b_string, "%")
	random_fortune := random_intn(1, len(fortuner_array))

	fmt.Println(fortuner_array[random_fortune])

}

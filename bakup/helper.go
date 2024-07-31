package helper

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


var files []string

func RandInt(min int, max int) int {
  return min + rand.Intn(max-min) // so as to choose any number bw 1 to 6
}

func PathWalk(path string, f os.FileInfo, err error) error {
  if filepath.Ext(path) == ".dat" {
    return nil
  }

  if f.IsDir() {
    return nil
  }

  files = append(files, path)
  return nil
}

func Program() {
  fortuneCommand := exec.Command("fortune", "-f")
  pipe, err := fortuneCommand.StderrPipe()

  if err != nil {
    panic(err)
  }

  fortuneCommand.Start()
  outputStream := bufio.NewScanner(pipe)
  outputStream.Scan()
  line := outputStream.Text()
  root := line[strings.Index(line,"/"):]  //  extract - /usr/share/games/fortunes all fortunes are found here
  
  // fmt.Println(path)
  err = filepath.Walk(root, PathWalk)
  if err != nil {
    panic(err)
  }



  // printing the file tree structure 
  for i := 1; i < len(files); i++ {
    fmt.Println(files[i])
  }
  // fmt.Println(len(files)) 6 files found
  // fmt.Println(files)

  random_num := RandInt(1,len(files))
  random_filename := files[random_num]
  // fmt.Println(random_filename) //from which file the quote is taken from 

  // we can now randomly choose a file.
  // time to read contents of a random file, each quote is seperated by a %

  f, err := os.Open(random_filename)
  if err != nil {
    panic(err)
  }

  defer f.Close()

  b, err := io.ReadAll(f)
  if err != nil {
    panic(err)
  }

  b_string := string(b)

  quotes_array := strings.Split(b_string, "%")
  random_quote := RandInt(1, len(quotes_array))

  fmt.Println(quotes_array[random_quote])



}


package helper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
var files []string

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

  // printing the file tree structure 
  // for i := 1; i < len(files); i++ {
  //   fmt.Println(files[i])
  // }

  // fmt.Println(files)
  if err != nil {
    panic(err)
  }


}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	myInput := githubactions.GetInput("INPUT_MYINPUT")
	output := fmt.Sprintf("Hello %s", myInput)
	err := listFiles()
	check(err)
	githubactions.SetOutput("myOutput", output)
}

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func listFiles()  error {
	action := githubactions.New()
	return filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		action.Infof("found file:",path)
		return nil
	})

}
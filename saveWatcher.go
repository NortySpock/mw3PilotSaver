package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var initalListOfFiles []os.FileInfo

//filesToWatch = dict()
var userRequestedFile string = ""
var targetFileExtension string = ".sav"
var secondsBetweenRefreshes int = 5
var archiveFolderName string = "backupPilots"
var filename = "hello.blah"
var extension = filepath.Ext(filename)

func main() {
	initalListOfFiles, err := ioutil.ReadDir("./")

	if err != nil {
		panic(err)
	}
	for _, file := range initalListOfFiles {
		fmt.Println(file.Name())
	}

	if len(userRequestedFile) <= 0 {
		fmt.Println("No specific file to watch was requested.")
		fmt.Println("Watching all *" + targetFileExtension + " files in this folder.")
	}
	for _, file := range initalListOfFiles {
		fmt.Println("filename: " + file.Name() + " has ext:" + filepath.Ext(file.Name()))
	}
}

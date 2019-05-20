package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var initalListOfFiles []os.FileInfo

var userRequestedFile string = ""
var targetFileExtension string = ".txt"
var secondsBetweenRefreshes int = 5
var archiveFolderName string = "backupPilots"

func main() {

	//get the inital list of files
	initalListOfFiles, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	//set up the target list of files,
	//as only a handful have the right file extension
	// filename -> MD5sum
	var targetFilenameWithMD5sum map[string]string
	targetFilenameWithMD5sum = make(map[string]string)

	if len(userRequestedFile) <= 0 {
		fmt.Println("No specific file to watch was requested.")
		fmt.Println("Watching all *" + targetFileExtension + " files in this folder.")

		for _, file := range initalListOfFiles {
			if filepath.Ext(file.Name()) == targetFileExtension {
				targetFilenameWithMD5sum[file.Name()] = "<blank>"
			}
		}
	} else {
		fmt.Println("Looking for " + userRequestedFile + " as requested")
		for _, file := range initalListOfFiles {
			if file.Name() == userRequestedFile {
				targetFilenameWithMD5sum[file.Name()] = "<blank>"
			}
		}
	}

	fmt.Println(" ")

	if len(targetFilenameWithMD5sum) <= 0 {
		fmt.Println("Found no files to watch, exiting...")
		os.Exit(0)
	}

	fmt.Println("Found " + strconv.Itoa(len(targetFilenameWithMD5sum)) + " file(s) to watch.")
	for key := range targetFilenameWithMD5sum {
		fmt.Println("Watching:", key)
	}

	fmt.Println(" ")
	fmt.Println("Use CTRL-C to exit this watcher ")
	fmt.Println(" ")

}

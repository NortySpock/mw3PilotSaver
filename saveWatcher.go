package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var initalListOfFiles []os.FileInfo

var userRequestedFile string = ""
var targetFileExtension string = ".txt"
var timeBetweenRefreshes string = "3s"
var archiveFolderName string = "backupPilots"

func main() {
    fmt.Println("This go program is incomplete and should not be expected to work.")
    fmt.Println("The python program in the same repository should work properly.")
    fmt.Println(" ")


	var secondsBetweenRefreshes time.Duration
	secondsBetweenRefreshes, err := time.ParseDuration(timeBetweenRefreshes)
	if err != nil {
		panic(err)
	}

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

	var done bool = false
	for !done {
		for key, value := range targetFilenameWithMD5sum {
			hashMD5, err := wrappedMD5sum(key)
			if err == nil {
				if(value != hashMD5){
                    targetFilenameWithMD5sum[key] = hashMD5
                    fmt.Println("updated "+key+" was backed up at " + value)
                }
			} else {
				fmt.Println("burped on " + key)
			}

		}
		time.Sleep(secondsBetweenRefreshes)
	}

}



func wrappedMD5sum(filename string) (string, error) {
	var md5ret string
	f, err := os.Open(filename)
	if err != nil {
		return md5ret, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return md5ret, err
	}
	hashInBytes := h.Sum(nil)[:16]
	md5ret = hex.EncodeToString(hashInBytes)
	return md5ret, nil
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// is item in array
func inArray(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// is damn file
func isDumnFile(fileName string) bool {
	if isMusicFile(fileName) != true {
		return false
	}
	matched, _ := regexp.MatchString(`(?m)^\..*.mp3$`, fileName)
	return matched
}

// is music file
func isMusicFile(filename string) bool {
	allowedExtensions := []string{".mp3"}
	fileExtension := filepath.Ext(filename)

	return inArray(allowedExtensions, fileExtension) == true
}

// is allowed file to skip
func isAllowedFile(filename string) bool {
	ignoredFiles := []string{"GROMUSB2.CFG"}
	return isMusicFile(filename) || inArray(ignoredFiles, filename) == true
}

// count digits
func CountDigits(i int) int {
	return len(strconv.FormatInt(int64(i), 10))
}

// Generate new filename
func generateFilename(i int, precision int, fileExt string) string {
	return fmt.Sprintf("%0"+strconv.FormatInt(int64(precision), 10)+"d", i) + fileExt
}

// Check if file exists in range
func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	} else {
		return false
	}
}

// Rename file to 0000 format
func renameFile(fileName string, files []os.FileInfo, path string) bool {
	var numLevel int
	if CountDigits(len(files)) > 4 {
		numLevel = CountDigits(len(files))
	} else {
		numLevel = 4
	}

	for cntr := 1; cntr < int(math.Pow10(numLevel)); cntr++ {
		newFilename := generateFilename(cntr, numLevel, filepath.Ext(fileName))
		if fileExists(path+"/"+newFilename) != true {
			os.Rename(path+"/"+fileName, path+"/"+newFilename)
			fmt.Println("Rename file " + fileName + " => " + newFilename)

			break
		}
	}

	return true
}

// Scan dir
func scanDir(path string, action string, level int) {
	fmt.Println("Scan dir: " + path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Cannot scan dir: " + path)
	}

	for _, f := range files {
		if f.IsDir() {
			scanDir(path+"/"+f.Name(), action, level+1)
		} else {
			if action == "clean" {
				if isAllowedFile(f.Name()) != true {
					os.Remove(path + "/" + f.Name())
					fmt.Println("Remove file: " + path + "/" + f.Name())
				}
			} else if action == "fix" {
				if isDumnFile(f.Name()) == true {
					os.Remove(path + "/" + f.Name())
					fmt.Println("Remove file: " + path + "/" + f.Name())
					continue
				}
				if isMusicFile(f.Name()) == true && isAllowedFile(f.Name()) == true {
					matched, err := regexp.MatchString(`(?m)^\d*.mp3$`, f.Name())
					if err != nil {
						panic(err)
					}
					if matched != true {
						renameFile(f.Name(), files, path)
					}
				}

			} else {
				panic("Action not implemented")
			}
		}
	}

}

// Main function
func main() {
	volumePath := flag.String("volume", "/Volumes", "Path to GROM stick volume")
	action := flag.String("action", "Action", "fix OR clean")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	scanDir(*volumePath, *action, 0)
}

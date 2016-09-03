package main

import (
	"fmt"
	"gosubrename/dirtools"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	if len(argsWithoutProg) != 1 {
		fmt.Println("Incorrect number of arguments")
		return
	}

	dirPath := argsWithoutProg[0]
	exists, err := dirtools.Dircheck(dirPath)
	if exists {

	} else {
		fmt.Println("Directory path is not correct")
		fmt.Println(err)
	}

}

// RenameSubs check .avi and .srt files in a directory
// if the subtitles names are not inline with the video names, they are changed accordingly
func RenameSubs(dirpath string) {
	avifiles, err := dirtools.GetFiles(dirpath, dirtools.Avi)
	srtfiles, err := dirtools.GetFiles(dirpath, dirtools.Srt)

	//copy subtitles for safety
	destDirPath := dirpath + dirtools.Separator + "test"
	os.Mkdir(destDirPath, 0755)
	err = dirtools.CopyExtFiles(dirpath, destDirPath, dirtools.Srt)
	if err != nil {
		fmt.Println("Issue while copying srt files to: ", destDirPath)
		fmt.Println(err)
		return
	}

	// initialize map from episode number to videos titles
	rgx := regexp.MustCompile("\\d+")
	videos := make(map[string]string)
	for _, f := range avifiles {
		filename := f.Name()
		res := rgx.FindString(filename)
		if res != "" {
			var extension = filepath.Ext(filename)
			var name = filename[0 : len(filename)-len(extension)]
			videos[res] = name
		}
	}

	// replace the srt title by the avi title
	for _, f := range srtfiles {
		res := rgx.FindString(f.Name())
		if res != "" {
			title := videos[res]
			err = os.Rename(dirpath+dirtools.Separator+f.Name(), dirpath+dirtools.Separator+title)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

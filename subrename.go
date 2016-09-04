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

	arglen := len(argsWithoutProg)
	if arglen != 1 && arglen != 3 {
		fmt.Println("Incorrect number of arguments: ", arglen)
		fmt.Println(help())
		return
	}

	dirPath := argsWithoutProg[0]
	fmt.Println("Checking path : " + dirPath)
	exists, err := dirtools.Dircheck(dirPath)
	if !exists {
		fmt.Println("Directory path is not correct")
		fmt.Println(err)
		fmt.Println(help())
		return
	} else {
		if len(argsWithoutProg) == 3 {
			renameSubs(dirPath, argsWithoutProg[1], argsWithoutProg[2])
		} else {
			renameSubs(dirPath, "", "")
		}
	}
}

// RenameSubs check .avi and .srt files in a directory
// if the subtitles names are not inline with the video names, they are changed accordingly
func renameSubs(dirpath, videoext, subext string) {
	//taking extension parameters into account
	if videoext == "" {
		videoext = dirtools.Avi
	}
	if subext == "" {
		subext = dirtools.Srt
	}

	fmt.Printf("Renaming subtitles for extensions %s and %s\n", videoext, subext)

	avifiles, err := dirtools.GetFiles(dirpath, videoext)
	if len(avifiles) < 1 {
		fmt.Printf("No video files found for extension %s\n", videoext)
		return
	}

	srtfiles, err := dirtools.GetFiles(dirpath, subext)
	if len(avifiles) < 1 {
		fmt.Printf("No subtitles found for extension %s\n", subext)
		return
	}

	//copy subtitles for safety
	destDirPath := dirpath + dirtools.Separator + "test"
	fmt.Printf("Backing up subtitles in directory %s\n", destDirPath)
	os.Mkdir(destDirPath, 0755)
	err = dirtools.CopyExtFiles(dirpath, destDirPath, subext)
	if err != nil {
		fmt.Println("Issue while copying srt files to: ", destDirPath)
		fmt.Println(err)
		return
	}

	// initialize map from episode number to videos titles
	fmt.Println("Creating map of video titles")
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
	fmt.Println("Renaming wrong subtitles")
	for _, f := range srtfiles {
		res := rgx.FindString(f.Name())
		if res != "" {
			title := videos[res]
			err = os.Rename(dirpath+dirtools.Separator+f.Name(), dirpath+dirtools.Separator+title+subext)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// Help prints a helper message for the command line usage
func help() string {
	help := `
	Usage for the command line execution:
		gosubrename arg1 arg2 arg3
		arg1, mandatory : full path to the directory where you want to rename the subtitles
		arg2, optional : extension of the video files. Default is .avi
		arg3, optional : extension of the subtitles files. Default is .srt
		Example : gosubrename C:/temp/myDir .mkv .srt
			  gosubrename C:/temp/myDir
	`
	return help
}

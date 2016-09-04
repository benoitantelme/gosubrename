package dirtools

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const Filename = "filename"

// TestDirCheck tests behavior when there is no directory
func TestNotExist(t *testing.T) {
	badpath := "notachance"
	boolean, err := Dircheck(badpath)
	if boolean != false || err == nil {
		t.Error("Did not detect non existing path for ", badpath)
	}
}

// TestDirCheck tests behavior when the path points to a file
func TestNotADir(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "sometempfile")
	filename := file.Name()
	defer os.Remove(filename)

	if err != nil {
		fmt.Println(err)
		t.Error("Could not create a temp file ", filename)
	}

	boolean, err := Dircheck(file.Name())
	if boolean != false || err == nil {
		t.Error("Did not detect non directory path for ", filename)
	}
}

// TestDirCheck tests correct verification of a directory
func TestDirCheck(t *testing.T) {
	boolean, err := Dircheck(os.TempDir())
	if boolean != true || err != nil {
		t.Error("Did not detect directory path for ", os.TempDir())
	}
}

// TestCountFilesWithExt tests correct count of files for a specific extension
func TestCountFilesWithExt(t *testing.T) {
	dirpath, err := CreateDirWithFiles(10, Filename, Avi, "")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dirpath)

	allfiles, _ := ioutil.ReadDir(dirpath)

	nbr := CountFilesWithExt(allfiles, Avi)

	if nbr != 10 {
		t.Error("Could not count number of files with extension .avi ", os.TempDir())
	}
}

// TestGetFiles tests correct retrieval of files
func TestGetFiles(t *testing.T) {
	dirpath, err := CreateDirWithFiles(12, Filename, Avi, "")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dirpath)

	avifiles, err := GetFiles(dirpath, Avi)

	if len(avifiles) != 12 {
		t.Error("Could not get the files in dir ", dirpath)
	}
}

// TestCopyExtFiles tests correct copy of files
func TestCopyExtFiles(t *testing.T) {
	dirpath, err := CreateDirWithFiles(12, Filename, Srt, "")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dirpath)

	destDirPath := dirpath + Separator + "test"
	os.Mkdir(destDirPath, 0755)

	err = CopyExtFiles(dirpath, destDirPath, Srt)
	if err != nil {
		t.Error(err)
	}

	avifiles, err := GetFiles(destDirPath, Srt)
	if len(avifiles) != 12 {
		t.Error("Could not get copied the files in created dir ", dirpath+"/test")
	}

}

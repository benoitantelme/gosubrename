package dirtools

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const Filename = "filename"

func TestNotExist(t *testing.T) {
	badpath := "notachance"
	boolean, err := Dircheck(badpath)
	if boolean != false || err == nil {
		t.Error("Did not detect non existing path for ", badpath)
	}
}

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

func TestDirCheck(t *testing.T) {
	boolean, err := Dircheck(os.TempDir())
	if boolean != true || err != nil {
		t.Error("Did not detect directory path for ", os.TempDir())
	}
}

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

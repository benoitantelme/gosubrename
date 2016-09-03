package main

import (
	"gosubrename/dirtools"
	"os"
	"regexp"
	"testing"
)

func TestRenameSubs(t *testing.T) {
	rgx := regexp.MustCompile("wrongfilename")
	testdir := os.TempDir() + dirtools.Separator + "testsubrename"
	dirtools.CreateDirWithFiles(10, "filename", dirtools.Avi, testdir)
	defer os.RemoveAll(testdir)
	dirtools.CreateDirWithFiles(10, "wrongfilename", dirtools.Srt, testdir)
	RenameSubs(testdir)

	srtfiles, err := dirtools.GetFiles(testdir, dirtools.Srt)
	if err != nil {
		t.Error(err)
	}

	for _, f := range srtfiles {
		filename := f.Name()

		res := rgx.MatchString(filename)
		if res == true {
			t.Error("File name has not been modified for .srt files")
		}
	}

}

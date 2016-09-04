package main

import (
	"gosubrename/dirtools"
	"os"
	"regexp"
	"testing"
)

// TestRenameSubs tests correct renaming of subtitles
func TestRenameSubs(t *testing.T) {
	rgx := regexp.MustCompile("wrongfilename")
	testdir := os.TempDir() + dirtools.Separator + "testsubrename"
	dirtools.CreateDirWithFiles(10, "filename", dirtools.Avi, testdir)
	defer os.RemoveAll(testdir)
	dirtools.CreateDirWithFiles(10, "wrongfilename", dirtools.Srt, testdir)
	renameSubs(testdir, "", "")

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

// TestRenameWrongExtensionSubs tests behavior when the video files have another extension that specified (default one is .avi)
func TestRenameWrongExtensionSubs(t *testing.T) {
	rgx := regexp.MustCompile("wrongfilename")
	testdir := os.TempDir() + dirtools.Separator + "testsubrename"
	dirtools.CreateDirWithFiles(10, "filename", ".mkv", testdir)
	defer os.RemoveAll(testdir)
	dirtools.CreateDirWithFiles(10, "wrongfilename", dirtools.Srt, testdir)
	renameSubs(testdir, "", "")

	srtfiles, err := dirtools.GetFiles(testdir, dirtools.Srt)
	if err != nil {
		t.Error(err)
	}

	for _, f := range srtfiles {
		filename := f.Name()

		res := rgx.MatchString(filename)
		if res == false {
			t.Error("File name has been modified while extension does not correspond")
		}
	}
}

// TestRenameNoSubs tests behavior when there is no subtitles for the video files
func TestRenameNoSubs(t *testing.T) {
	rgx := regexp.MustCompile("filename")
	testdir := os.TempDir() + dirtools.Separator + "testsubrename"
	dirtools.CreateDirWithFiles(10, "filename", ".mkv", testdir)
	defer os.RemoveAll(testdir)

	renameSubs(testdir, "", "")

	avifiles, err := dirtools.GetFiles(testdir, dirtools.Srt)
	if err != nil {
		t.Error(err)
	}

	for _, f := range avifiles {
		filename := f.Name()

		res := rgx.MatchString(filename)
		if res == false {
			t.Error("Issue while no subtitles were present")
		}
	}
}

// TestRenameNoVideoFiles tests behavior when the video files are not present
func TestRenameNoVideoFiles(t *testing.T) {
	rgx := regexp.MustCompile("wrongfilename")
	testdir := os.TempDir() + dirtools.Separator + "testsubrename"
	defer os.RemoveAll(testdir)
	dirtools.CreateDirWithFiles(10, "wrongfilename", dirtools.Srt, testdir)
	renameSubs(testdir, "", "")

	srtfiles, err := dirtools.GetFiles(testdir, dirtools.Srt)
	if err != nil {
		t.Error(err)
	}

	for _, f := range srtfiles {
		filename := f.Name()

		res := rgx.MatchString(filename)
		if res == false {
			t.Error("Issue while no video files are present")
		}
	}
}

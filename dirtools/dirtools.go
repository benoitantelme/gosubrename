package dirtools

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const Avi = ".avi"
const Srt = ".srt"
const Separator = string(filepath.Separator)

// CreateDirWithFiles creates a directory dirpath or os.TempDir() + Separator + "test" if dirpath == ""
// inside that directory, it creates nbr files with name filename + X + extension
func CreateDirWithFiles(nbr int, filename, extension, dirpath string) (string, error) {
	if dirpath == "" {
		dirpath = os.TempDir() + Separator + "test"
	}
	os.Mkdir(dirpath, 0755)

	for i := 0; i < nbr; i++ {
		filename := fmt.Sprintf(dirpath+Separator+filename+"%d"+extension, i)
		err := ioutil.WriteFile(filename, nil, 0644)

		if err != nil {
			fmt.Println(err)
			return "", errors.New("Could not create a temp file " + filename)
		}
	}
	return dirpath, nil
}

// Dircheck will check if a file exists and is a directory
func Dircheck(path string) (bool, error) {
	// check if the path exist
	src, err := os.Stat(path)
	if err != nil {
		fmt.Println("Path does not exist")
		return false, err
	}

	// check if the path is a directory
	if !src.IsDir() {
		fmt.Println("Path is not a directory")
		return false, errors.New("Path " + path + " is not a directory")
	}

	return true, nil
}

// GetFiles will get the files in the path that have the given extension
func GetFiles(path string, extension string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Could not get files from path ", path)
		return nil, err
	}

	avis := make([]os.FileInfo, CountFilesWithExt(files, extension))

	var nbr int
	for _, f := range files {
		if filepath.Ext(f.Name()) == extension {
			avis[nbr] = f
			nbr += 1
		}
	}

	return avis, nil
}

// CountFileWithExt return the number of files with a certain extension in a list of files info
func CountFilesWithExt(files []os.FileInfo, ext string) int {
	var nbr int
	for _, f := range files {
		if filepath.Ext(f.Name()) == ext {
			nbr += 1
		}
	}
	return nbr
}

// copyExtFiles copy files with a specific extension from an initial directory to a destination directory
func CopyExtFiles(dir, dest, ext string) error {
	files, err := GetFiles(dir, ext)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f != nil {
			err = copyFile(dir+Separator+f.Name(), dest+Separator+f.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherwise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func copyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}

	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}

	if err = os.Link(src, dst); err == nil {
		return
	}

	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		return
	}

	err = out.Sync()
	return
}

package labfile

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Unzip a file on disk
func Unzip(src, dest string) ([]string, error) {
	var ListOfExtractedFiles []string
	var fileName string

	CreatePathIfNotExists(dest)

	r, err := zip.OpenReader(src)
	if err != nil {
		return ListOfExtractedFiles, err
	}

	defer func() {
		if err := r.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) (string, error) {
		var fileName string

		rc, err := f.Open()
		if err != nil {
			log.Println(err.Error())
			return "", err
		}

		defer func() {
			if err := rc.Close(); err != nil {
				log.Println(err.Error())
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			fileName = path
			// Sometimes the list of files to unzip is not in the right order
			subPath := GetFilePath(path)
			if subPath != "" {
				CreatePathIfNotExists(subPath)
			}
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Println(err.Error())
				return "", err
			}

			defer func() {
				if err := f.Close(); err != nil {
					log.Println(err.Error())
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				log.Println(err.Error())
				return "", err
			}
		}

		return fileName, nil
	}

	for _, f := range r.File {
		fileName, err = extractAndWriteFile(f)
		if err != nil {
			log.Println(err.Error())
			//return err
		} else {
			if fileName != "" {
				ListOfExtractedFiles = append(ListOfExtractedFiles, fileName)
			}
		}
	}

	return ListOfExtractedFiles, nil
}

// Unzip a file on disk
func UnzipSpecific(src, dest, specificFile string) ([]string, error) {
	var ListOfExtractedFiles []string
	var fileName string
	var extractAll bool

	if specificFile == "*.*" {
		extractAll = true
	} else {
		extractAll = false
	}

	CreatePathIfNotExists(dest)

	r, err := zip.OpenReader(src)
	if err != nil {
		return ListOfExtractedFiles, err
	}

	defer func() {
		if err := r.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) (string, error) {
		var fileName string

		rc, err := f.Open()
		if err != nil {
			log.Println(err.Error())
			return "", err
		}

		defer func() {
			if err := rc.Close(); err != nil {
				log.Println(err.Error())
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			if extractAll || strings.Contains(path, specificFile) {
				fileName = path
				// Sometimes the list of files to unzip is not in the right order
				subPath := GetFilePath(path)
				if subPath != "" {
					CreatePathIfNotExists(subPath)
				}
				f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					log.Println(err.Error())
					return "", err
				}

				defer func() {
					if err := f.Close(); err != nil {
						log.Println(err.Error())
					}
				}()

				_, err = io.Copy(f, rc)
				if err != nil {
					log.Println(err.Error())
					return "", err
				}
			} else {
				fileName = ""
			}
		}

		return fileName, nil
	}

	for _, f := range r.File {
		fileName, err = extractAndWriteFile(f)
		if err != nil {
			log.Println(err.Error())
			//return err
		} else {
			if fileName != "" {
				ListOfExtractedFiles = append(ListOfExtractedFiles, fileName)
			}
		}
	}

	return ListOfExtractedFiles, nil
}

// Get the file size of the given file on disk
func GetFileSize(fileName string) int64 {
	file, err := os.Stat(fileName)
	if err != nil {
		// Could not obtain stat, handle error
	}

	return file.Size()
}

// Get the extension from a full filename
func GetFileExtension(fileName string) string {
	return filepath.Ext(strings.ToLower(fileName))
}

// Get the file name (including extension) from a full filename
func GetFileBase(fullFileName string) string {
	return filepath.Base(fullFileName)
}

// Get the file name (without extension) from a full filename
func GetFileName(fullFileName string) string {
	fileName := GetFileBase(fullFileName)
	fileParts := split(fileName, ".")

	if len(fileParts) > 0 {
		return fileParts[0]
	}

	return fileName
}

// Get the path from a full filename
func GetFilePath(fileName string) string {
	path, err := filepath.Abs(filepath.Dir(fileName))
	if err != nil {
		return ""
	}

	return path
}

// Delete a file on disk
func RemoveFile(fileName string) {
	os.Remove(fileName)
}

// Delete a path on disk
func RemoveAll(path string) {
	os.RemoveAll(path)
}

// Move a file on disk
func MoveFile(originalFileName, destinationFileName string) error {
	return os.Rename(originalFileName, destinationFileName)
}

// Copy file on disk
func CopyFile(src, dest string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}

	defer in.Close()

	out, err := os.Create(dest)
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

// Check if a file exists on disk
func CheckIfFileExist(filename string) (fileExists bool) {
	if _, err := os.Stat(filename); err == nil {
		fileExists = true
	}

	return
}

// Get the list of files in the given path on disk
func GetFullFilesList(dir, filter string) (fileList []string, err error) {
	err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		apply := true

		// Cut File Name
		index := strings.Index(path, dir)

		if index >= 0 {
			path = path[index+len(dir):]
		}

		// Check Filter on extension
		if filter != "" {
			apply = GetFileExtension(path) == ".svg"
		}

		// Add in List if apply
		if apply {
			fileList = append(fileList, path)
		}

		return nil
	})

	return
}

// Save Text in a file on disk
func SaveText(fileName string, content []byte) error {
	// Link to permission chmod : http://ss64.com/bash/chmod.html
	return ioutil.WriteFile(fileName, content, 0644)
}

// Load Text from a file on disk
func LoadText(fileName string) (file []byte, err error) {
	file, err = ioutil.ReadFile(fileName)

	return
}

func CopyDir(source string, dest string) (err error) {
	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	return
}

func CompleteDir(path string) string {
	// 47 is code for /
	if path == "" || path[len(path)-1] == 47 {
		return path
	}

	return path + "/"
}

func CreatePathIfNotExists(path string) string {
	// Create directory structure on disk if not already exists
	if path != "" {
		// 0777 (octal) => 0b 111 111 111 (binary) => permissions rwxrwxrwx
		os.MkdirAll(path, 0777)
	}

	return path
}

func ChangePath(path string) {
	os.Chdir(path)
}

func GetMime(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	return getMimeFromFile(f)
}

/*
Private functions
*/

func getMimeFromFile(file *os.File) (string, error) {
	buffer := make([]byte, 128)

	// Read the first 128 bytes to get the content type
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	return http.DetectContentType(buffer[:n]), nil
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

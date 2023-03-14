package basic_functions

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// version 1.0.1

func Read_file(file string) ([]byte, error) {
	if _, err := os.Stat(file); err == nil {
		if err != nil {
			return []byte(""), err
		}
	}
	return os.ReadFile(file)
}

func Write_file(file string, data string) error {
	// save the file to disk
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err2 := f.WriteString(data)
	if err2 != nil {
		return err
	}
	return nil
}

func Delete_file(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

func File_or_directory_exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			// File or directory does not exist
			return false
		} else {
			// Some other error. The file may or may not exist
			return false
		}
	}
	return true
}

func List_files_in_dir(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}

func List_files_by_extension(dir string, extension string) []string {
	var extfiles []string

	var dirfiles, err = List_files_in_dir(dir)
	if err != nil {
		log.Fatalln("List_files_by_extension - Error: " + err.Error())
	}

	for _, file := range dirfiles {
		if strings.HasSuffix(file.Name(), extension) {
			extfiles = append(extfiles, file.Name())
		}
	}
	return extfiles
}

func Current_directory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// the executable directory
	directory := filepath.Dir(ex)
	return directory
}

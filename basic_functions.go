package basic_functions

import (
	"archive/zip"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

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

func Download_file(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
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

func Read_file_in_zip(zip_path string, file_path string) (string, error) {
	r, err := zip.OpenReader(zip_path)
	if err != nil {
		return "", err
	}
	defer r.Close()

	file_data := ""
	for _, f := range r.File {
		if !strings.Contains(f.Name, file_path) {
			continue
		}

		buffer := new(strings.Builder)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(buffer, rc)
		if err != nil {
			log.Fatal(err)
		}
		file_data = buffer.String()
		break
	}
	return file_data, err
}

func Serialize(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		dataEncoder := gob.NewEncoder(file)
		dataEncoder.Encode(object)
		file.Close()
	}
	return err
}

func Deserialize(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
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

func Calc_md5(data string) string {
	hash := md5.Sum([]byte(data))
	md5_string := hex.EncodeToString(hash[:])
	return md5_string
}

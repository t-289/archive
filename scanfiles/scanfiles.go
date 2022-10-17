package scanfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/t-289/archive/crud"
)

func ScanFolder(root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		//splitPath := strings.Split(path, "/")
		//parentPath := splitPath[len(splitPath)-2]
		if !info.IsDir() {
			//fileHash := hsfile.HashFile(path)
			ftype, _ := getType(path)
			fmt.Printf("\n Root: %s | Path: %s | FileName: %s | Type: %s", root, path, info.Name(), ftype)
			//fileType, _ := analize(path)
			//saveData(fileHash.Sum(nil), info.Name(), path, info.Size())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func saveData(hsval []byte, fileName string, filePath string, fileSize int64) {
	// new path
	newPath := ""

	// move file
	queryString := fmt.Sprintf("INSERT INTO files_info (hash, file_name, old_path, new_path, file_size) VALUES (%x, %s, %s, %s, %v)", hsval, fileName, filePath, newPath, fileSize)
	db := crud.DBInsert(queryString)

	fmt.Printf("DB Insert: %v", db)
}

func getType(filePath string) (string, error) {
	buf, _ := ioutil.ReadFile(filePath)

	contentType := ""

	// read file
	if filetype.IsImage(buf) {
		contentType = "image"
	} else if filetype.IsExtension(buf, "pdf") {
		contentType = "document"
	} else if filetype.IsDocument(buf) {
		contentType = "document"
	} else if filetype.IsArchive(buf) {
		contentType = "archive"
	} else if filetype.IsAudio(buf) {
		contentType = "audio"
	} else if filetype.IsFont(buf) {
		contentType = "font"
	} else if filetype.IsVideo(buf) {
		contentType = "video"
	} else {
		contentType = "unknow"
	}

	return contentType, nil
}

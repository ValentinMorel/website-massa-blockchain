package reader

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"

	"techtest/rpc"
)

func Read(data *rpc.GetDatastoreEntriesResponse) (files map[string][]byte, err error) {
	// Decode all the values contained in the array in b64
	result, _ := base64.StdEncoding.DecodeString(string(data.Result[0].CandidateValue))

	zipReader, _ := zip.NewReader(bytes.NewReader(result), int64(len(result)))
	content := make(map[string][]byte)

	// Separate the file in the zip archive
	for _, file := range zipReader.File {
		log.Printf("found file ! name => %s\n", file.Name)
		fc, _ := file.Open()
		defer fc.Close()
		contentZip, _ := ioutil.ReadAll(fc)
		content[file.Name] = contentZip
	}
	if err != nil {
		return nil, err
	}
	return content, nil
}

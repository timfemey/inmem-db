package helper

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"

	"github.com/cespare/xxhash"
)

func Hash(key string) uint64 {
	h := xxhash.Sum64String(key)
	return h
}

func ConvertToByte(initialVal any) ([]byte, error) {
	byteBuf, err := json.Marshal(initialVal)
	if err != nil {
		return nil, err
	}
	return byteBuf, nil
}

func convertToAny(initialVal []byte) (any, error) {

	var decodedData any
	err := json.Unmarshal(initialVal, &decodedData)
	if err != nil {
		return nil, err
	}
	return decodedData, nil

}

func Compress(dataIn any) (bytes.Buffer, error) {
	var compressedData bytes.Buffer
	zlibScripter := zlib.NewWriter(&compressedData)

	data, err := ConvertToByte(dataIn)
	if err != nil {
		return compressedData, err
	}
	_, err2 := zlibScripter.Write(data)
	if err2 != nil {
		return compressedData, err
	}
	zlibScripter.Close()

	return compressedData, nil

}

func DeCompress(compressedData *bytes.Buffer) (any, error) {

	zlibReader, err := zlib.NewReader(compressedData)
	if err != nil {
		return "", err
	}

	decompressedData, err := io.ReadAll(zlibReader)
	if err != nil {

		return "", err
	}

	data, err := convertToAny(decompressedData)
	if err != nil {
		return "", err
	}

	return data, nil
}

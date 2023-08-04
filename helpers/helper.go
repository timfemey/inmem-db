package helper

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"

	"github.com/cespare/xxhash"
)

func Hash(key string) uint64 {
	h := xxhash.Sum64String(key)
	return h
}

func convertToByte(initialVal any) ([]byte, error) {
	var byteBuf bytes.Buffer
	encoder := gob.NewEncoder(&byteBuf)
	err := encoder.Encode(initialVal)
	if err != nil {
		return []byte(""), err
	}
	return byteBuf.Bytes(), nil
}

func convertToAny(initialVal []byte) (any, error) {
	buffer := bytes.NewBuffer(initialVal)
	var decodedData any
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&decodedData)
	if err != nil {
		return "", err
	}
	return decodedData, nil
}

func Compress(data any) (bytes.Buffer, error) {
	var compressedData bytes.Buffer
	gzipScripter := gzip.NewWriter(&compressedData)
	inBytes, err := convertToByte(data)
	if err != nil {
		return compressedData, err
	}
	_, err2 := gzipScripter.Write(inBytes)
	if err2 != nil {
		return compressedData, err
	}
	gzipScripter.Close()
	return compressedData, nil
}

func DeCompress(compressedData bytes.Buffer) (any, error) {

	gzipReader, err := gzip.NewReader(&compressedData)
	if err != nil {
		return "", err
	}
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return "", err
	}
	data, err := convertToAny(decompressedData)
	if err != nil {
		return "", err
	}
	return data, nil
}

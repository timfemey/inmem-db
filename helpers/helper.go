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

func Compress(data any) (any, error) {
	var compressedData bytes.Buffer
	gzipScripter := gzip.NewWriter(&compressedData)
	inBytes, err := convertToByte(data)
	if err != nil {
		return "", err
	}
	_, err2 := gzipScripter.Write(inBytes)
	if err2 != nil {
		return "", err
	}
	gzipScripter.Close()
	inByte, err3 := convertToByte(compressedData)
	if err3 != nil {
		return "", err3
	}
	data, err4 := convertToAny(inByte)
	if err4 != nil {
		return "", err4
	}

	return data, nil
}

func DeCompress(compressedData any) (any, error) {
	inBytes, err1 := convertToByte(compressedData)
	if err1 != nil {
		return "", err1
	}
	inBytesBuffer := bytes.NewBuffer(inBytes)
	gzipReader, err := gzip.NewReader(inBytesBuffer)
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

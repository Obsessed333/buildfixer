package main


import(
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
	"strings"
)

func DecodePoBCode(code string) ([]byte, error){

	code = strings.ReplaceAll(code, "-", "+")
	code = strings.ReplaceAll(code, "_", "/")

	data, err := base64.StdEncoding.DecodeString(code)
	if err != nil{
		return nil, err
	}

	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil{
		return nil, err
	}
	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)
	if err != nil{
		return nil, err
	}

	return out.Bytes(), nil
}
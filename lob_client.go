package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

const lobAddressBaseUrl = "https://api.lob.com/v1/addresses/"
const lobPostcardBaseUrl = "https://api.lob.com/v1/postcards"

// https://gist.github.com/andrewmilson/19185aab2347f6ad29f5
// https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
func CreatePostCard(fromAddress, toAddress, frontFileName string) string {
	frontFile, _ := os.Open(frontFileName)
	defer frontFile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	frontPart, _ := writer.CreateFormFile("front", filepath.Base(frontFile.Name()))
	io.Copy(frontPart, frontFile)

	_ = writer.WriteField("back", "<body>hello, back!</body>")
	_ = writer.WriteField("description", "golang multipart test")
	_ = writer.WriteField("to", toAddress)
	_ = writer.WriteField("from", fromAddress)

	writer.Close()

	req, err := http.NewRequest("POST", lobPostcardBaseUrl, body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(os.Getenv("LOB_API_TEST_KEY")+":")))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}

func GetAddress(addressId string) string {
	lobAddressUrl := lobAddressBaseUrl + addressId

	req, err := http.NewRequest("GET", lobAddressUrl, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(os.Getenv("LOB_API_TEST_KEY")+":")))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}

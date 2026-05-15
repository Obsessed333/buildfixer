package main

import(
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchFromPobbIn(pobURL string) (string, error){

	rawURL := strings.Replace(pobURL, "pobb.in/", "pobb.in/pob/", 1)
	fmt.Printf("Downloading compressed code from: %s\n", rawURL)

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil{
		return "", err
	}
	req.Header.Add("User-Agent", "build-fixer-agent") 

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200{
		return "", fmt.Errorf("pobb.in error: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}


	
	return string(bodyBytes), nil
}
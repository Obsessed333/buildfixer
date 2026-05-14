package main

import(
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"github.com/obsessed333/buildfixer/internal/models"

)

func FetchFromPobbIn(pobURL string) (*models.PoBData, error){

	rawURL := strings.Replace(pobURL, "pobb.in/", "pobb.in/pob/", 1)

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil{
		return nil, err
	}
	req.Header.Add("User-Agent", "build-fixer-agent") 

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200{
		return nil, fmt.Errorf("pobb.in error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	xmlBytes, err := DecodePoBCode(string(body))
	if err != nil{
		return nil, fmt.Errorf("failed to decompress PoB string: %v", err)
	}
	var pobData models.PoBData
	err = xml.Unmarshal(xmlBytes, &pobData)
	if err != nil{
		return nil, fmt.Errorf("failed to parse XML: %v", err)
	}
	return &pobData, nil
}
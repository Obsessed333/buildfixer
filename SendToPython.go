package main


import(
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PythonRequest struct {
	XMLData string `json:"xml_data"`
	Budget  string `json:"budget"`
}

func SendToPythonAgent(xmlString string, budget string) (string, error) {
	url := "http://127.0.0.1:8000/analyze"

	payload := PythonRequest{
		XMLData: xmlString,
		Budget:  budget,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not connect to Python service: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("python service error (%s): %s", resp.Status, string(body))
	}

	return string(body), nil
}
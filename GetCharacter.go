package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/obsessed333/buildfixer/internal/models"
)

func GetCharacter(accountName, charName string) (*models.CharacterData, error) {
	cacheFile := fmt.Sprintf("cache_%s_%s.json", accountName, charName)

	_, err := os.Stat(cacheFile)
	if err == nil{
		fmt.Println("--- Loading from Local Cache ---")
		fileData, err := ioutil.ReadFile(cacheFile)
		var result models.CharacterData
		err = json.Unmarshal(fileData, &result)
		return &result, err
	}

	fmt.Println("--- Cache Missing: Fetching from PoE API ---")
	apiData, err := fetchFromPoE(accountName, charName)
	if err != nil{
		return nil, err
	}

	file, err := json.MarshalIndent(apiData, "", " ")
	if err != nil{
		return nil, err
	}
	_ = ioutil.WriteFile(cacheFile, file, 0644)

	return apiData, nil
}
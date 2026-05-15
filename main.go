package main


import (
	"fmt"
)



func main(){
	testURL := "https://pobb.in/l8Tyfca2Jjaa"

	compressedCode, err := FetchFromPobbIn(testURL)
	if err != nil{
		fmt.Printf("Error downloading build: %V\n", err)
		return
	}

	printHelper(compressedCode)
}
	
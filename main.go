package main


import (
	"fmt"
)



func main(){
	testURL := "https://pobb.in/l8Tyfca2Jjaa"

	pobData, err := FetchFromPobbIn(testURL)
	if err != nil{
		fmt.Printf("Error execution failed: %v\n", err)
		return
	}

	printHelper(pobData)
}
	
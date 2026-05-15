package main

import(
	"fmt"
)
func printHelper(compressedCode string) {

	fmt.Println("Sending raw compressed URL-safe Base64 string to Python calculation server...")
	pythonResult, err := SendToPythonAgent(compressedCode, "50 Divine")
	if err != nil{
		fmt.Printf("Network handoff failed: %v\n", err)
		return
	}
	fmt.Println("\n--- Response received from Python Service ---")
	fmt.Println(pythonResult)
	}
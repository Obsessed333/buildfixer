package main


import (
	"fmt"
	"github.com/obsessed333/buildfixer/internal/models"
)



func main(){
	char, err := GetCharacter("obsessedw-7623", "RAPE_PLUNDER_RAVAGE")
	if err != nil{
		panic(err)
	}

	fmt.Printf("Successfully loaded %s, a level %d %s\n",
		char.Character.Name,
		char.Character.Level,
		char.Character.Class)
}
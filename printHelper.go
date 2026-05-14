package main

import(
	"fmt"
	"strings"
	"github.com/obsessed333/buildfixer/internal/models"
)

func printHelper(pobData *models.PoBData) {

		activeIdx := 0
		for i, set := range pobData.Items.ItemSets {
			if set.ID == pobData.Items.ActiveItemSet {
				activeIdx = i
				break
			}
		}
		var activeSlots []models.Slot
		if len(pobData.Items.ItemSets) > 0 {
			activeSlots = pobData.Items.ItemSets[activeIdx].Slots
		}

		fmt.Println("\n--- Build Successfully Downloaded ---")
		fmt.Printf("Class: %s\n", pobData.Build.Class)
		fmt.Printf("Level: %d\n", pobData.Build.Level)
		fmt.Printf("Total Items Indexed: %d\n", len(pobData.Items.ItemList))
		fmt.Printf("Total Equipment Slots Used: %d\n", len(activeSlots))

		fmt.Println("\n--- Equipped Gear Layout ---")

		itemMap := make(map[int]string)
		for _, item := range pobData.Items.ItemList {
			itemMap[item.ID] = item.Raw
		}

		for _, slot := range activeSlots {
			if rawText, exists := itemMap[slot.ItemID]; exists{

				lines := strings.Split(strings.TrimSpace(rawText), "\n")
				itemName := lines[0]
				if strings.HasPrefix(lines[0], "Rarity") && len(lines) > 1{
					itemName = lines[1]
				}
				fmt.Printf("[%s]: %s\n", slot.Name, itemName)
			}
		}
	}
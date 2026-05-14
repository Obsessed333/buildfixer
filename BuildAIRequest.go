package main

import(
	"github.com/obsessed333/buildfixer/internal/models"
)


func BuildAIRequest(pobData *models.PoBData, budget string) models.AIGapAnalysisRequest{
	gearMap := make(map[string]string)


	itemTextMap := make(map[int]string)
	for _, item := range pobData.Items.ItemList {
		itemTextMap[item.ID] = item.Raw
	}
	activeIdx := 0
	for i, set := range pobData.Items.ItemSets {
		if set.ID == pobData.Items.ActiveItemSet {
			activeIdx = i
			break
		}
	}
	var activeSlots []models.Slot
	if len(pobData.Items.ItemSets) > 0 && activeIdx < len(pobData.Items.ItemSets) {
		activeSlots = pobData.Items.ItemSets[activeIdx].Slots
	}

	for _, slot := range activeSlots {
		if text, exists := itemTextMap[slot.ItemID]; exists {
			gearMap[slot.Name] = text
		}
	}

	return models.AIGapAnalysisRequest{
		CharacterClass: pobData.Build.Class,
		CurrentLevel:   pobData.Build.Level,
		CurrentGear:    gearMap,
		Budget:         budget,
	}
}
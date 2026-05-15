package models

import "encoding/xml"


type PoBData struct {
	XMLName xml.Name `xml:"PathOfBuilding"`
	Build   Build    `xml:"Build"`
	Skills  Skills   `xml:"Skills"`
	Tree    Tree     `xml:"Tree"`
	Items   Items    `xml:"Items"`
}

type PobbInStatsResponse struct {
	Stats struct {
		Life          int     `json:"life"`
		EnergyShield  int     `json:"energyShield"`
		Mana          int     `json:"mana"`
		FireResist    int     `json:"fireResist"`
		ColdResist    int     `json:"coldResist"`
		LightningResist int   `json:"lightningResist"`
		ChaosResist   int     `json:"chaosResist"`
		TotalDPS      float64 `json:"totalDPS"`
	} `json:"stats"`
	Keystones []string `json:"keystones"` // e.g., ["Mind Over Matter", "Chaos Inoculation"]
}

type Build struct {
	Level int    `xml:"level,attr"`
	Class string `xml:"className,attr"`
}

type Skills struct {
	SkillSets []SkillSet `xml:"SkillSet"`
}

type SkillSet struct {
	Skills []PoBSkill `xml:"Skill"`
}

type PoBSkill struct {
	Enabled string      `xml:"enabled,attr"`
	GemList []PoBGem    `xml:"Gem"`
}

type PoBGem struct {
	NameSpec    string `xml:"nameSpec,attr"`
	Level       int    `xml:"level,attr"`
	Quality     int    `xml:"quality,attr"`
	Enabled     string `xml:"enabled,attr"`
}

type Tree struct {
	ActiveSpec int    `xml:"activeSpec,attr"`
	Specs      []Spec `xml:"Spec"`
}

type Spec struct {
	Nodes string `xml:"nodes,attr"` 
}

type Items struct {
	ActiveItemSet int       `xml:"activeItemSet,attr"`
	ItemList      []PoBItem `xml:"Item"`
	ItemSets      []ItemSet `xml:"ItemSet"`
}

type ItemSet struct {
	ID    int    `xml:"id,attr"`
	Slots []Slot `xml:"Slot"` 
}

type PoBItem struct {
	ID   int    `xml:"id,attr"`
	Raw  string `xml:",chardata"` 
}

type Slot struct {
	Name   string `xml:"name,attr"`   // e.g., "Body Armour"
	ItemID int    `xml:"itemId,attr"` // References PoBItem.ID
}

type AIGapAnalysisRequest struct {
	CharacterClass string            `json:"class"`
	CurrentLevel   int               `json:"level"`
	CurrentGear    map[string]string `json:"current_gear"`      
	Budget         string            `json:"budget"`          
	TreeNodes      string            `json:"tree_nodes"`        
}
package models

type CharacterData struct {
	Items      []Item `json:"items"`
	Character  struct {
		Name       string `json:"name"`
		Class      string `json:"class"`
		Level      int    `json:"level"`
	} `json:"character"`
}

type Item struct {
	Name        string   `json:"name"`        
	BaseType    string   `json:"typeLine"`    /
	Slot        string   `json:"inventoryId"` 
	Links       int      `json:"-"`           
	ExplicitMods []string `json:"explicitMods"`
	ImplicitMods []string `json:"implicitMods"`
	EnchantMods  []string `json:"enchantMods"`
	Sockets      []Socket `json:"sockets"`
}

type Socket struct {
	Group int    `json:"group"` // Items in the same group are linked
	Attr  string `json:"attr"`  // Strength (R), Dexterity (G), Intelligence (B)
}

type PassiveTreeData struct {
	Hashes []int `json:"hashes"` // These IDs correspond to nodes on the tree
}

type AIGapAnalysisRequest struct {
	AccountName    string            `json:"account_name"`
	CharacterClass string            `json:"class"`
	CurrentLevel   int               `json:"level"`
	CurrentGear    map[string]string `json:"current_gear"` // Slot: ItemName
	KeyStats       map[string]int    `json:"stats"`        // Life: 3500, ChaosRes: -20
	Budget         string            `json:"budget"`       // 10 Chaos" or "50 Divine
	MetaComparison []string          `json:"missing_meta_items"` // Missing chase items(Replica Farrul's Fur for Flicker Strike for example)
}
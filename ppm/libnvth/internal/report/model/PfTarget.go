package model

// PfTarget pftarget
type PfTarget struct {
	FuncName   string   `json:"func_name" bson:"func_name"`
	URL        string   `json:"url" bson:"url"`
	Parameters []string `json:"parameters" bson:"parameters"`
}

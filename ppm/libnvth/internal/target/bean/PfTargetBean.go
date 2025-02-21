package bean

// PfTargetBean bean
type PfTargetBean struct {
	FuncName   string   `json:"func_name" bson:"func_name"`
	URL        string   `json:"url" bson:"url"`
	Parameters []string `json:"parameters" bson:"parameters"`
}

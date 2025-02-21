package bean

// Parameter parameters
type Parameter struct {
	Type  string `json:"type" bson:"type"`
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

// TargetDetailBean bean for user who is assigned to the scope
type TargetDetailBean struct {
	FuncName   string      `json:"func_name" bson:"func_name"`
	HTTPMethod string      `json:"http_method" bson:"http_method"`
	URL        string      `json:"url" bson:"url"`
	Severity   string      `json:"severity" bson:"severity"`
	Parameters []Parameter `json:"parameters" bson:"parameters"`
	Remark     string      `json:"remark" bson:"remark"`
}

package bean

// VulTemplateBean bean
type VulTemplateBean struct {
	ID              string `json:"id" bson:"_id"`
	CategoryID      string `json:"category_id" bson:"category_id"`
	Rating          string `json:"rating" bson:"rating"`
	Name            string `json:"name" bson:"name"`
	IdentifierName  string `json:"identifier_name" bson:"identifier_name"`
	Overview        string `json:"overview" bson:"overview"`
	DetectionMethod string `json:"detection_method" bson:"detection_method"`
	Description     string `json:"description" bson:"description"`
	Condition       string `json:"condition" bson:"condition"`
	PossibleImpact  string `json:"possible_impact" bson:"possible_impact"`
	Countermeasure  string `json:"countermeasure" bson:"countermeasure"`
	Remarks         string `json:"remarks" bson:"remarks"`
}

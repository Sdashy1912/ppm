package dao

// VulTemplateDAOParam param for Create, Update
type VulTemplateDAOParam struct {
	ID              string `bson:"_id"`
	CategoryID      string `bson:"category_id"`
	Rating          string `bson:"rating"`
	Name            string `bson:"name"`
	IdentifierName  string `bson:"identifier_name"`
	Overview        string `bson:"overview"`
	DetectionMethod string `bson:"detection_method"`
	Description     string `bson:"description"`
	Condition       string `bson:"condition"`
	PossibleImpact  string `bson:"possible_impact"`
	Countermeasure  string `bson:"countermeasure"`
	Remarks         string `bson:"remarks"`
}

package bean

// VulTemplateBean bean
type VulTemplateBean struct {
	ID             string `json:"id" bson:"_id"`
	IdentifierName string `json:"identifier_name" bson:"identifier_name"`
	Name           string `json:"name" bson:"name"`
}

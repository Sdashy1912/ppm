package bean

// VulCategoryGetBean bean
type VulCategoryGetBean struct {
	ID        string            `json:"id" bson:"_id"`
	Name      string            `json:"name" bson:"name"`
	Templates []VulTemplateBean `json:"templates" bson:"templates"`
}

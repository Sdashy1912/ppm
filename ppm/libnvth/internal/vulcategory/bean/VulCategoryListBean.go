package bean

// VulCategoryListBean bean
type VulCategoryListBean struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

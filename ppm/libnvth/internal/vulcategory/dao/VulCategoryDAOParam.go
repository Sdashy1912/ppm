package dao

// VulCategoryDAOParam param for Create, Update
type VulCategoryDAOParam struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

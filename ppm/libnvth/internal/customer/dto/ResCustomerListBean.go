package dto

// ResCustomerListBean bean
type ResCustomerListBean struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Industry    string `json:"industry"`
}

package dto

// ResCustomerGetBean bean
type ResCustomerGetBean struct {
	ID          string                   `json:"id"`
	Email       string                   `json:"email"`
	Name        string                   `json:"name"`
	PhoneNumber string                   `json:"phone_number"`
	Address     string                   `json:"address"`
	Industry    string                   `json:"industry"`
	Projects    []ResCustomerProjectBean `json:"projects"`
}

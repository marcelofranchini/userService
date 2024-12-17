package domain

type UserCreateRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Document     string `json:"document" bson:"document"`
	DocumentType string `json:"documentType" bson:"documentType"`
	Phone        string `json:"phone" bson:"phone"`
}

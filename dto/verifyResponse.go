package dto

type VerifiedReponse struct {
	IsAuthorized bool `json:"is_authorized"`
}
type NotVerifiedReponse struct {
	IsAuthorized bool   `json:"is_authorized"`
	Message      string `json:"message"`
}

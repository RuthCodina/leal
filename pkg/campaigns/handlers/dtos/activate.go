package dtos

type ActivateDTO struct {
	Name string `json:"name", validate:"required"`
	Id   int    `json:"branch_id", validate:"required"`
}

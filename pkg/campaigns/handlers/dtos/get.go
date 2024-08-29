package dtos

type GetAllByBranchDTO struct {
	BranchId int  `json:"branch_id", validate:"required"`
	Activate bool `json:"activate", validate:"required"`
}

type GetAllByCommerceDTO struct {
	CommerceId int  `json:"commerce_id", validate:"required"`
	Activate   bool `json:"activate", validate:"required"`
}

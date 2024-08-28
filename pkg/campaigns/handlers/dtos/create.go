package dtos

type CampaignCreateDTO struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	BranchOffice int    `json:"branch_office" validate:"required"`
}

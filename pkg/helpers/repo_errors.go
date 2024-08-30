package helpers

import "errors"

var (
	ErrNoExist          = errors.New("the data that you are browsing no exist")
	ErrActivateCampaign = errors.New("can't activate the campaign because is the branch has one campaign activated")
	ErrAccumulation     = errors.New("no cashback, branch no have active campaign")
)

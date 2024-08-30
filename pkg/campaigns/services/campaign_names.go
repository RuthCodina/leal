package services

const (
	TEXACO_1 string = "texaco mayo sucursal 5"
	TEXACO_2 string = "texaco mayo sucursal 7"
)

var Names map[string]interface{} = map[string]interface{}{
	"texaco mayo sucursal 5": CampaignTexaco1,
	"texaco mayo sucursal 7": CampaignTexaco2,
}

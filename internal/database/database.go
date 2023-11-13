package database

type LTVDataJSON struct {
	CampaignID string  `json:"CampaignId"`
	Country    string  `json:"Country"`
	Users      uint    `json:"Users"`
	LTV1       float64 `json:"Ltv1"`
	LTV2       float64 `json:"Ltv2"`
	LTV3       float64 `json:"Ltv3"`
	LTV4       float64 `json:"Ltv4"`
	LTV5       float64 `json:"Ltv5"`
	LTV6       float64 `json:"Ltv6"`
	LTV7       float64 `json:"Ltv7"`
}

type LTVDataCSV struct {
	CampaignID string  `csv:"CampaignId"`
	Country    string  `csv:"Country"`
	UserID     uint64  `csv:"UserId"`
	LTV1       float64 `csv:"Ltv1"`
	LTV2       float64 `csv:"Ltv2"`
	LTV3       float64 `csv:"Ltv3"`
	LTV4       float64 `csv:"Ltv4"`
	LTV5       float64 `csv:"Ltv5"`
	LTV6       float64 `csv:"Ltv6"`
	LTV7       float64 `csv:"Ltv7"`
}

package models

type DolarCotation struct {
	USDBRL DolarCotationInfos `json:"USDBRL"`
}

type DolarCotationInfos struct {
	Code        string `json:"code"`
	Codein      string `json:"codein"`
	Name        string `json:"name"`
	High        string `json:"high"`
	Low         string `json:"low"`
	VarBid      string `json:"varBid"`
	PctChange   string `json:"pctChange"`
	Bid         string `json:"bid"`
	Ask         string `json:"ask"`
	Timestamp   string `json:"timestamp"`
	CreatedDate string `json:"createdDate"`
}

package models

type Transference struct {
	AccountIDSender  int     `json:"account_sender"`
	AccountIDReciber int     `json:"account_reciber"`
	Amount           float64 `json:"amount"`
	RecordDate       string  `json:"record_date"`
}

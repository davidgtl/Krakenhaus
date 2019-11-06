package models

type Activity struct {
	ID int64 `json:"patient_id"`
	Name string `json:"activity"`
	Start int64 `json:"start"`
	End int64 `json:"end"`
}

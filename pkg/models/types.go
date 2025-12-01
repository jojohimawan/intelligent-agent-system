package models

type DecodedSignal struct {
	Source string
	Param  string
	Value  interface{}
	Unit   string
}

type VSSPoint struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
	Unit  string      `json:"unit"`
}

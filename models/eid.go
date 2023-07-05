package model

type EidData struct {
	Name   string `json:"name"`
	Type   string `json:"idtype"`
	Number string `json:"idnum"`
}

type ResJSON struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Data    *EidData `json:"data,omitempty"`
}

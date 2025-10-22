package model

type URLRequest struct {
	URLs    []string `json:"urls"`
	Timeout int      `json:"timeout"`
}

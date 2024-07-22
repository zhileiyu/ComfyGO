package client

type queueReq struct {
	Clear  string   `json:"clear"`
	Delete []string `json:"delete"`
}

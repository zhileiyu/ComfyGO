package client

type queueReq struct {
	Clear  string   `json:"clear"`
	Delete []string `json:"delete"`
}

type promptReq struct {
	ClientID string               `json:"client_id"`
	Prompt   map[string]promptObj `json:"prompt"`
}

type promptObj struct {
}

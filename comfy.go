package comfyGO

import "github.com/zhileiyu/comfyGO/internal/client"

const defaultConfPath = "comfy.toml"

type ComfyClient struct {
	clientImp *client.CFImp
}

func NewComfyClient(configFile ...string) *ComfyClient {
	comfyClient := &ComfyClient{}
	if len(configFile) == 0 {
		comfyClient.clientImp = client.New(defaultConfPath)
	} else {
		comfyClient.clientImp = client.New(defaultConfPath)
	}
	return comfyClient
}

func (client *ComfyClient) ISServerOK() bool {
	return client.clientImp.ServerAvailable()
}

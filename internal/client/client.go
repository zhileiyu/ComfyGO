package client

import (
	"fmt"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"net/http"
)

type CFImp struct {
	endPoint string
	server   ServerInfo
	hc       *http.Client
}

func New(configFile string) *CFImp {
	config, err := loadConfig(configFile)
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	reqBaseUrl := fmt.Sprintf("http://%s:%d", config.BaseUrl, config.Port)
	imp := &CFImp{
		endPoint: reqBaseUrl,
		hc:       &http.Client{},
		server: ServerInfo{
			Endpoint: reqBaseUrl,
		},
	}
	imp.checkServerInfo()
	return imp
}

func (imp *CFImp) ServerAvailable() bool {
	imp.checkServerInfo()
	return imp.server.Available
}

func (imp *CFImp) checkServerInfo() {
	resp, err := imp.hc.Get(imp.endPoint + systemStats)
	if err != nil {
		logger.Info(err)
		imp.serverOffline()
	}
	if resp.StatusCode != http.StatusOK {
		logger.Info("not 200")
		imp.serverOffline()
	}
	info := systemInfo{}
	err = bindHttpRes(resp, &info)
	if err != nil {
		imp.serverOffline()
	}
	imp.server.Available = true
	imp.server.systemInfo = info
}

func (imp *CFImp) ServerInfo(refresh bool) ServerInfo {
	if refresh {
		imp.checkServerInfo()
	}
	return imp.server
}

func (imp *CFImp) serverOffline() {
	imp.server = ServerInfo{
		Available: false,
	}
}

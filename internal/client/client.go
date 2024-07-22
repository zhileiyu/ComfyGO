package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"github.com/zhileiyu/comfyGO/internal/utils"
	"log"
	"net/http"
)

type CFImp struct {
	endPoint  string
	server    serverInfo
	hc        *http.Client
	wsConn    *websocket.Conn
	clientID  string
	available bool
}

func New(configFile string) *CFImp {
	config, err := loadConfig(configFile)
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	httpBaseUrl := fmt.Sprintf("http://%s:%d", config.BaseUrl, config.Port)
	imp := &CFImp{
		endPoint: httpBaseUrl,
		hc:       &http.Client{},
		server: serverInfo{
			endpoint: httpBaseUrl,
		},
		clientID: utils.UniqueStr(),
	}
	imp.checkServerInfo()
	if imp.server.available {
		wsBaseUrl := fmt.Sprintf("ws://%s:%d/ws?clientID=%s", config.BaseUrl, config.Port, imp.clientID)
		imp.wsConn, _, err = websocket.DefaultDialer.Dial(wsBaseUrl, nil)
		if err != nil {
			return nil
		}
		logger.Info("setup ws", wsBaseUrl)
	} else {
		return nil
	}
	return imp
}

func (imp *CFImp) ServerAvailable() bool {
	imp.checkServerInfo()
	return imp.server.available
}

func (imp *CFImp) checkServerInfo() {
	info := systemInfo{}
	err := imp.GetData(systemStats, &info)
	if err != nil {
		imp.serverOffline()
	}
	imp.server.available = true
	imp.server.systemInfo = info
}

func (imp *CFImp) ServerInfo(refresh bool) serverInfo {
	if refresh {
		imp.checkServerInfo()
	}
	return imp.server
}

func (imp *CFImp) serverOffline() {
	imp.server = serverInfo{
		available: false,
	}
}

func (imp *CFImp) PostEnqueue() error {
	err := imp.PostJson(queue, nil, nil)
	return err
}

func (imp *CFImp) PostJson(path string, reqData interface{}, resData interface{}) error {
	var jsonData []byte
	if reqData != nil {
		jsonData, _ = json.Marshal(reqData)
	}

	resp, err := http.Post(imp.endPoint+path, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()
	if !errors.Is(err, nil) {
		//logger.Error("Failed to send POST request: %v", err)
		log.Printf("Failed to send POST: %v", errors.Is(err, nil))
		log.Fatal(err)

		return err
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		imp.checkServerInfo()
		logger.Error("Failed to send POST request: %d", resp.StatusCode)
		return nil
	}

	if resData == nil {
		return nil
	}

	err = utils.BindHttpRes(resp, resData)
	if err != nil {
		return err
	}
	return nil
}

func (imp *CFImp) GetData(path string, data interface{}) error {
	resp, err := imp.hc.Get(imp.endPoint + path)
	defer resp.Body.Close()
	if err != nil {
		logger.Info(err)
		imp.serverOffline()
		return err
	}
	if resp.StatusCode != http.StatusOK {
		logger.Info("not 200")
		imp.serverOffline()
		return err
	}
	err = utils.BindHttpRes(resp, data)
	if err != nil {
		return err
	}
	return nil
}

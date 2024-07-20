package client

import (
	"encoding/json"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"io"
	"net/http"
)

func bindHttpRes(res *http.Response, target interface{}) error {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error()
		return err
	}
	err = json.Unmarshal(body, target)
	if err != nil {
		logger.Error()
		return err
	}
	return nil
}

package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"io"
	"net/http"
)

func UniqueStr() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}

func BindHttpRes(res *http.Response, target interface{}) error {
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

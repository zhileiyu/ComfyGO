package comfy

import (
	"encoding/json"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"github.com/zhileiyu/comfyGO/internal/utils"
)

type workflowData struct {
	LastNodeID int     `json:"last_node_id"`
	LastLinkID int     `json:"last_link_id"`
	Nodes      []*node `json:"nodes"`
	Links      []*link `json:"links"`
	Version    float64 `json:"version"`
}

type Workflow struct {
	Data     *workflowData
	jsonData []byte
	ID       string
}

type link []interface{}

func NewWorkflowFromJson(jsonData []byte) *Workflow {
	wd := new(workflowData)
	if err := json.Unmarshal(jsonData, wd); err != nil {
		logger.Error("load workflow json fail: " + err.Error())
		return nil
	}
	wf := &Workflow{
		Data:     wd,
		jsonData: jsonData,
		ID:       utils.UniqueStr(),
	}
	return wf
}

// AddNode, return node id
func (wf *workflowData) AddNode(nc nodeClass) int {
	nid := len(wf.Nodes)
	n := &node{
		NodeID: nid,
	}
	wf.Nodes = append(wf.Nodes, n)
	return nid
}

func (wf *Workflow) LoadData(jsonData []byte) error {
	err := json.Unmarshal(wf.jsonData, &wf.Data)
	if err != nil {
		logger.Error("load workflow data err:", err)
		return err
	}
	wf.jsonData = jsonData
	return nil
}

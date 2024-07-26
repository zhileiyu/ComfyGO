package comfyGO

import (
	"github.com/zhileiyu/comfyGO/internal/client"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"os"
)

type Workshop struct {
	//load balance
	cp *client.Pool
}

func NewWorkshop(configFile ...string) *Workshop {
	var clientPool *client.Pool
	if len(configFile) == 0 {
		clientPool = client.NewPool(defaultConfPath)
	} else {
		clientPool = client.NewPool(configFile[0])
	}
	return &Workshop{
		cp: clientPool,
	}
}

// Ready check if comfy server ready
//func (ws *Workshop) Ready() bool {
//	return ws.client.ServerAvailable()
//}

func (workshop *Workshop) LoadWorkflowFromFile(filename string) *Workflow {
	workflowJson, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Error("workflow file not exist")
			return nil
		} else {
			logger.Error("read workflow file error: " + err.Error())
			return nil
		}
	}
	c := workshop.cp.ProperClient()
	wid := c.CreateWorkflow(workflowJson)
	return &Workflow{
		id: wid,
		ws: workshop,
	}
}

package comfyGO

type Workflow struct {
	ws *Workshop
	id string
}

func (wf *Workflow) Enqueue() *Job {
	wf.ws.cp.ProperClient().PromptEnqueue(wf.id)
}

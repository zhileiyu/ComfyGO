package comfy

type node struct {
	NodeID        int           `json:"id"`
	NodeType      nodeClass     `json:"type"`
	Order         int           `json:"order"`
	Inputs        []nodeInputs  `json:"inputs"`
	Outputs       []nodeOutputs `json:"outputs"`
	Mode          int           `json:"mode"`
	Properties    string        `json:"properties"`
	WidgetsValues []interface{} `json:"widgets_values"`
}

type nodeInputs struct {
	Name      string     `json:"name"`
	InputType inputClass `json:"type"`
	Link      int        `json:"link"`
}

type nodeOutputs struct {
	Name      string     `json:"name"`
	InputType inputClass `json:"type"`
	Links     []int      `json:"links"`
	SlotIndex int        `json:"slot_index"`
}

type nodeClass int
type inputClass int
type outputClass int

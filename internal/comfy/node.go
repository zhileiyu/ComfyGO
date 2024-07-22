package comfy

type node struct {
	nodeID   int
	nodeType nodeClass
	order    int
	inputs   []nodeInputs
	outputs  []nodeOutputs
	mode     int
	version  string
}

type nodeInputs struct {
	name      string
	inputType inputClass
	link      int
}

type nodeOutputs struct {
	name      string
	inputType inputClass
	links     []int
	slotIndex int
}

type nodeClass int
type inputClass int
type outputClass int

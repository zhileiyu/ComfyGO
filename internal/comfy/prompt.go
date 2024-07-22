package comfy

type prompt struct {
	inputs map[interface{}]string
	pct    promptClassType
}

type promptClassType int

const (
	emptyPrompt promptClassType = iota
)

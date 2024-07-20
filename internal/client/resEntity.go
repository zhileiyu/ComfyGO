package client

//ref :

// systemInfo contains a system info and acc Devices infos
type systemInfo struct {
	System  *System   `json:"system"`
	Devices []*Device `json:"devices"`
}

// System contains system info
type System struct {
	OS             string `json:"os"`
	PythonVersion  string `json:"python_version"`
	EmbeddedPython bool   `json:"embedded_python"`
}

// Device contains gpu info
type Device struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Index          int    `json:"index"`
	VRAMTotal      int64  `json:"vram_total"`
	VRAMFree       int64  `json:"vram_free"`
	TorchVRAMTotal int64  `json:"torch_vram_total"`
	TorchVRAMFree  int64  `json:"torch_vram_free"`
}

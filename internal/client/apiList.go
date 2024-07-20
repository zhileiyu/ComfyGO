package client

//ref https://github.com/comfyanonymous/ComfyUI/blob/master/server.py

const (
	index          = "/"
	embeddings     = "/embeddings"
	extensions     = "/extensions"
	uploadImage    = "/upload/image"
	uploadMask     = "/upload/mask"
	view           = "/view"
	viewMetadata   = "/view_metadata/%s"
	systemStats    = "/system_stats"
	prompt         = "/prompt"
	objectInfo     = "/object_info"
	nodeObjectInfo = "/object_info/%s"
	history        = "/history"
	promptHistory  = "/history/{prompt_id}"
	queue          = "/queue"
	interrupt      = "/interrupt"
	free           = "/free"
)

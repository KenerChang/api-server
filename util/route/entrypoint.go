package route

import (
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Entrypoint struct {
	Path string
	Method string
	Callback HandleFunc
	Version int
}

func NewEntrypoint(method, path string, callback HandleFunc, version int) Entrypoint {
	return Entrypoint{
		Path: path, Method: method, Callback: callback, Version: version,
	}
}
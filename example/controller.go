package example

import (
	"github.com/KenerChang/api-server/util"
	"github.com/KenerChang/api-server/util/logger"
	"github.com/KenerChang/api-server/util/route"
	"net/http"
)

var (
	ModuleInfo route.Module
)

func init() {
	ModuleInfo = route.Module{
		Name: "example",
		Entrypoints: []route.Entrypoint{
			route.NewEntrypoint("GET", "get", handleExampleGet, 1),
		},
	}
}

func handleExampleGet(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Hello string `json:"hello"`
	}

	logger.Info.Println(r, "hello, I am a example")

	res := Response{
		Hello: "hi~",
	}
	util.WriteJSON(w, res)
}

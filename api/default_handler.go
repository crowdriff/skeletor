package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/crowdriff/skeletor/skeletor"
	"github.com/zenazn/goji/web"
)

type defaultRes struct {
	Prod    bool   `json:"prod"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

// defaultHandler handles HTTP requests for / & /version
func defaultHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	conf := skeletor.GetConf()
	res, err := json.Marshal(defaultRes{
		Status:  "ok",
		Prod:    conf.Prod,
		Version: conf.Version,
	})

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, `{"status": "error"}`)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, string(res))
}

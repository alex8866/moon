package server

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

const (
	VERSION = 1
)

// Add api endpoints in here, you are allowed to use httprouter syntax to define parameters
func (s *Server) DefineEndpoints() {
	s.Endpoint("version", API_GET|API_POST, VersionEndpoint)
}

func VersionEndpoint(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ver := struct {
		Version float64 `json:"ver"`
	}{
		VERSION,
	}
	log.Infof("[%s] hit version api\n", req.RemoteAddr)
	data, _ := json.Marshal(ver)
	//w.Header().Add("Content-type", "encoding/json")

	w.Write(data)
}

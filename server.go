package mock_api

import (
	"net"
	"net/http"
	"strings"
)

type Server map[string]map[string]string

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if apiList, ok := s[r.Method]; ok {
		if api, ok := apiList[strings.Trim(r.URL.Path, "/")]; ok {
			w.Write([]byte(api))
			return
		}
	}
	http.NotFound(w, r)
}

func Run(addr string, router map[string]map[string]string, quit chan bool) (err error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	go http.Serve(ln, Server(router))
	go func() {
		for {
			select {
			case <-quit:
				ln.Close()
				return
			default:
			}
		}
	}()

	return
}

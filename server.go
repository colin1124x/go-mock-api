package mock_api

import (
	"net"
	"net/http"
	"strings"
)

type Server struct {
	routerMap  map[string]map[string]string
	routerFunc func(http.ResponseWriter, *http.Request)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if apiList, ok := s.routerMap[r.Method]; ok {
		if api, ok := apiList[strings.Trim(r.URL.Path, "/")]; ok {
			w.Write([]byte(api))
			return
		}
	}

	if s.routerFunc != nil {
		s.routerFunc(w, r)
		return
	}

	http.NotFound(w, r)
}

func New() *Server {
	return &Server{
		routerMap: make(map[string]map[string]string),
	}
}

func (s *Server) Route(method, path, body string) {
	if _, ok := s.routerMap[method]; !ok {
		s.routerMap[method] = make(map[string]string)
	}

	s.routerMap["GET"][path] = body
}

func (s *Server) Map(m map[string]map[string]string) {
	for Method, mm := range m {
		switch Method {
		case "GET", "PUT", "POST", "DELET":
			for p, b := range mm {
				s.Route(Method, p, b)
			}
		}
	}
}

func (s *Server) Run(addr string, quit chan bool) (err error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	go http.Serve(ln, s)
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

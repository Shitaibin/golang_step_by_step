package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type server struct {
	data map[string]interface{}
}

func newServer() *server {
	return &server{data: make(map[string]interface{})}
}

func (s *server) inject(key string, obj interface{}) {
	s.data[key] = obj
}

func (s *server) start(port string) {
	for k, v := range s.data {
		http.HandleFunc(fmt.Sprintf("/%v", k), func(w http.ResponseWriter, req *http.Request) {
			if req.Method == HTTP_GET {
				v, err := json.Marshal(v)
				if err != nil {
					w.Write([]byte(fmt.Sprintf("marshal error: %v", err)))
					return
				}
				w.Write(v)
				return
			}

			w.Write([]byte("not support other method"))
		})
	}

	go func() {
		fmt.Printf("http server error: %v", http.ListenAndServe(fmt.Sprintf("::%v", port), nil))
	}()
}

package fake

import "net/http"

type router struct {
	funcCalls map[string]int
}

func NewRouter() *router {
	return &router{}
}

func (r *router) ServeMux() *http.ServeMux {
	return http.NewServeMux()
}

func (r *router) RegisterEndpoint(path string, handler func(http.ResponseWriter, *http.Request)) {
	called(r.funcCalls, "RegisterEndpoint")
}

func (r *router) GetFuncCalls(funcName string) int {
	value, exists := r.funcCalls[funcName]
	if !exists {
		return 0
	}
	return value
}

func called(funcCalls map[string]int, funcName string) {
	if funcCalls == nil {
		funcCalls = make(map[string]int)
	}

	value, exists := funcCalls[funcName]
	if !exists {
		funcCalls[funcName] = 1
		return
	}
	funcCalls[funcName] = value + 1
}

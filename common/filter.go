package common

import "net/http"

type FilterHandle func(rw http.ResponseWriter, r *http.Response) error

type Filter struct {
	filterMap map[string]FilterHandle
}

func NewFilterHandle() *Filter  {
	return &Filter{make(map[string]FilterHandle)}
}

func (f *Filter)RegisterFilterUri(uri string, handle FilterHandle)  {
	 f.filterMap[uri] = handle
}



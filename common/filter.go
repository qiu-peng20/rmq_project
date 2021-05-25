package common

import "net/http"

//声明一个新的数据类型
type FilterHandle func(rw http.ResponseWriter, r *http.Request) error

//用来存储需要拦截的URI	
type Filter struct {
	filterMap map[string]FilterHandle
}

func NewFilterHandle() *Filter  {
	return &Filter{make(map[string]FilterHandle)}
}

func (f *Filter)RegisterFilterUri(uri string, handle FilterHandle)  {
	 f.filterMap[uri] = handle
}

func (f *Filter)GetFilterHandle(uri string) FilterHandle {
	return f.filterMap[uri]
}

type WebHandle func(rw http.ResponseWriter, rq *http.Request)

// 执行拦截器
func (f *Filter)Handle(handle WebHandle) func(rw http.ResponseWriter, rq *http.Request) {
	return func(rw http.ResponseWriter, rq *http.Request) {
		for path, handle := range f.filterMap{
			if path == rq.RequestURI {
				err := handle(rw, rq)
				if err != nil {
					_, _ = rw.Write([]byte(err.Error()))
					return
				}
				//跳出循环
				break
			}
		}
		handle(rw, rq)
	}
}



package main

import (
	"goapi/controller"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Node struct {
	path     string
	children map[string]*Node
	handler  http.HandlerFunc
}

func newNode(path string) *Node {
	return &Node{
		path:     path,
		children: make(map[string]*Node),
	}
}

func (n *Node) insert(path string, handler http.HandlerFunc) {
	current := n

	for {
		i := 0
		for i < len(path) && i < len(current.path) && path[i] == current.path[i] {
			i++
		}

		if i == len(current.path) {
			path = path[i:]
			if path == "" {
				current.handler = handler
				return
			}
			child, exists := current.children[path[0:1]]
			if !exists {
				child = newNode(path)
				child.handler = handler
				current.children[path[0:1]] = child
				return
			}
			current = child
		} else {
			child := newNode(current.path[i:])
			child.children = current.children
			child.handler = current.handler

			current.path = current.path[:i]
			current.children = map[string]*Node{
				child.path[0:1]: child,
			}
			if i < len(path) {
				newChild := newNode(path[i:])
				newChild.handler = handler
				current.children[path[i:i+1]] = newChild
			} else {
				current.handler = handler
			}
			return
		}
	}
}

func (n *Node) search(path string) *Node {
	current := n

	for {
		if len(path) == 0 {
			return current
		}

		i := 0
		for i < len(path) && i < len(current.path) && path[i] == current.path[i] {
			i++
		}

		if i == len(current.path) {
			path = path[i:]
			if path == "" {
				return current
			}
			child, exists := current.children[path[0:1]]
			if !exists {
				return nil
			}
			current = child
		} else {
			return nil
		}
	}
}

type Router struct {
	root *Node
}

func NewRouter() *Router {
	return &Router{root: newNode("/")}
}

func (r *Router) Handle(path string, handler http.HandlerFunc) {
	r.root.insert(path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	node := r.root.search(path)
	if node != nil && node.handler != nil {
		node.handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func main() {
	router := NewRouter()

	// Dynamically load and register handlers
	autoRegisterHandlers(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func autoRegisterHandlers(router *Router) {
	controllers := []interface{}{
		&controller.HomeController{},
		&controller.AboutController{},
	}

	for _, ctrl := range controllers {
		ctrlType := reflect.TypeOf(ctrl)
		for i := 0; i < ctrlType.NumMethod(); i++ {
			method := ctrlType.Method(i)
			if strings.HasSuffix(method.Name, "Handler") {
				route := "/" + strings.ToLower(strings.TrimSuffix(method.Name, "Handler"))
				handler := func(w http.ResponseWriter, r *http.Request) {
					reflect.ValueOf(ctrl).MethodByName(method.Name).Call([]reflect.Value{
						reflect.ValueOf(w),
						reflect.ValueOf(r),
					})
				}
				router.Handle(route, handler)
			}
		}
	}
}

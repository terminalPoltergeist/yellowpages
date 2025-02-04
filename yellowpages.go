package yellowpages

import (
	"errors"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

/*
A type alias for HTTP verbs

Uses all HTTP verbs supported by github.com/julienschmidt/httprouter

https://www.w3.org/Protocols/HTTP/Methods.html
*/
type method int

const (
    GET method = iota
    HEAD
    POST
    PUT
    DELETE
    OPTIONS
    PATCH
)

/*
An Endpoint colocates an http method and a path
*/
type Endpoint struct {
    Action method
    Path string
}

/*
A Service maps Endpoints to handlers

A handler generates responses for requests to the Endpoint
*/
type Service map[Endpoint]httprouter.Handle

/*
Register registers each Endpoint in the Service with the router
*/
func (s Service) Register(r *httprouter.Router) error {
    for endpoint, handler := range s {
        switch endpoint.Action {
        case GET:
            r.GET(endpoint.Path, handler)
        case HEAD:
            r.HEAD(endpoint.Path, handler)
        case POST:
            r.POST(endpoint.Path, handler)
        case PUT:
            r.PUT(endpoint.Path, handler)
        case DELETE:
            r.DELETE(endpoint.Path, handler)
        case OPTIONS:
            r.OPTIONS(endpoint.Path, handler)
        case PATCH:
            r.PATCH(endpoint.Path, handler)
        default:
            return errors.New(fmt.Sprintf("Unknown Endpoint Action for Path [%s]", endpoint.Path))
        }
    }
    return nil
}

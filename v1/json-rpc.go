package main

import (
	"io"
	"net"
	"net/http"
	"net/rpc"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  {
	return c.in.Read(p)
}

func (c *HttpConn) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}

func (c *HttpConn) Close() error {
	return nil
}

type Test struct{}

type HelloArgs struct {
	Name string
}

func (test *Test) Hello(args *HelloArgs, result *string) error {
	*result = "Hello " + args.Name
	return nil
}

//func (test *Test) Hello1(args *HelloArgs, result *string) error {
//    *result = "ti pidor " + args.Name
//    return nil
//}

func main() {

	server := rpc.NewServer()
	server.Register(&Test{})

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/rpc" {
			serverCodec := jsonrpc2.NewServerCodec(&HttpConn{in: r.Body, out: w}, server )

			w.Header().Set( "Content-type", "application/json" )
			w.WriteHeader(200)

			if err1 := server.ServeRequest(serverCodec) ; err1 != nil {
				http.Error(w, "Error while serving JSON request", 500)
				return
			}
		} else {
			http.Error(w, "Unknown request", 404)
		}
	} ) )
}
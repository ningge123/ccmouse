package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(port string, service interface{}) error {
	rpc.Register(service)

	listen, err := net.Listen("tcp", port)

	if err != nil {
		return err
	}

	for  {
		connect, err := listen.Accept()

		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(connect)
	}

	return nil
}

func NewClient(post string) (*rpc.Client, error)  {
	client, err := net.Dial("tcp", post)

	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(client), nil
}
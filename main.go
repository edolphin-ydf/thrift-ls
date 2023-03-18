package main

import (
	"context"
	"io"
	"log"
	"os"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/multierr"
)

type readWriteCloser struct {
	readCloser  io.ReadCloser
	writeCloser io.WriteCloser
}

func (r *readWriteCloser) Read(b []byte) (int, error) {
	return r.readCloser.Read(b)
}

func (r *readWriteCloser) Write(b []byte) (int, error) {
	return r.writeCloser.Write(b)
}

func (r *readWriteCloser) Close() error {
	return multierr.Append(r.readCloser.Close(), r.writeCloser.Close())
}

func main() {
	file, err := os.Create("/Users/bytedance/thriftlsp.log")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("start")
	conn := jsonrpc2.NewConn(
		jsonrpc2.NewStream(
			&readWriteCloser{
				readCloser:  os.Stdin,
				writeCloser: os.Stdout,
			},
		),
	)

	ctx := context.Background()
	conn.Go(ctx, protocol.ServerHandler(&Server{}, nil))
	<-ctx.Done()

	log.Println("exit")
}

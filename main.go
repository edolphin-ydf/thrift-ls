package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var logger *zap.Logger

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

func initLog() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	logPath := filepath.Join(user.HomeDir, ".thrift-lsp.log")
	logConfig := zap.NewDevelopmentConfig()
	logConfig.OutputPaths = []string{logPath}
	logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err = logConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initLog()

	log.Println("start")
	conn := jsonrpc2.NewConn(
		jsonrpc2.NewStream(
			&readWriteCloser{
				readCloser:  os.Stdin,
				writeCloser: os.Stdout,
			},
		),
	)

	ctx := protocol.WithLogger(context.Background(), logger)
	conn.Go(ctx, protocol.ServerHandler(&Server{}, nil))
	<-ctx.Done()

	log.Println("exit")
}

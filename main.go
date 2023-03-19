package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

var (
	logFile   = flag.String("logfile", "", "log file")
	logLevel = flag.Int("loglevel", int(zap.DebugLevel), "log level")
)

func initLog() {
	if *logFile == "" {
		// set logger to zap no op logger
		logger = zap.NewNop()
		return
	}
	var err error
	logConfig := zap.NewDevelopmentConfig()
	logConfig.OutputPaths = []string{*logFile}
	logConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(*logLevel))
	logger, err = logConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

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

package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jwenz723/grpcdemo/messaging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	serverAddrFlag = flag.String("server_addr", "", "The server address in the format of host:port")
	useStreaming   = flag.Bool("use_streaming", false, "Setting this will use grpc streaming instead of repeated single messages")
	waitNanos      = flag.Int("wait_nanos", 500, "The number of nanoseconds to wait before sending messages (this applies to both single and stream messages)")
	waitDuration   time.Duration
)

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	//logger, _ := zap.NewDevelopment() // This will cause logs to be written for every grpc req/resp
	defer logger.Sync()
	grpc_zap.ReplaceGrpcLogger(logger)

	logger.Info("test")
	serverAddr := "localhost:8080"
	waitDuration = time.Duration(*waitNanos) * time.Nanosecond

	if *serverAddrFlag != "" {
		serverAddr = *serverAddrFlag
	} else if serverAddrEnv, ok := os.LookupEnv("SERVER_ADDR"); ok {
		serverAddr = serverAddrEnv
	}

	// Start prometheus metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	// Note: load balancing between servers requires an L7 load balancer
	// (like linkerd) between this client and the servers
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_zap.UnaryClientInterceptor(logger),
			grpc_prometheus.UnaryClientInterceptor)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_zap.StreamClientInterceptor(logger),
			grpc_prometheus.StreamClientInterceptor)))

	if err != nil {
		logger.Fatal("fail to dial",
			zap.Error(err))
	}
	defer conn.Close()

	client := messaging.NewMessagingServiceClient(conn)

	if *useStreaming {
		waitc := make(chan struct{})
		stream, err := client.StreamMessages(context.Background())
		if err != nil {
			logger.Fatal("error opening stream",
				zap.Error(err))
		}
		defer stream.CloseSend()

		n := &messaging.Message{Sender: "client", Message: "A streamed message from a client"}
		go func() {
			for {
				_, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				if err != nil {
					logger.Fatal("error receiving message",
						zap.Error(err))
				}

				if err := stream.Send(n); err != nil {
					logger.Fatal("error sending message",
						zap.Error(err))
				}

				//handleReceivedMessage(m, "stream", logger)

				time.Sleep(waitDuration)
			}
		}()

		// Send a message to start the back and forth Notes
		if err := stream.Send(&messaging.Message{Sender: "client", Message: "Startup"}); err != nil {
			logger.Fatal("error sending startup message",
				zap.Error(err))
		}

		logger.Info("sent startup message")
		<-waitc
	} else {
		for {
			_, err := client.SendMessage(context.Background(), &messaging.Message{Sender: "client", Message: "A single message from a client"})
			if err != nil {
				logger.Error("error sending message",
					zap.Error(err))
			}

			//handleReceivedMessage(m, "single", logger)

			time.Sleep(waitDuration)
		}
	}
}

func handleReceivedMessage(m *messaging.Message, receiveType string, logger *zap.Logger) {
	//logger.Info("received from server",
	//	zap.String("sender", m.Sender),
	//	zap.String("message", m.Message),
	//	zap.Int("wait_nanos", *waitNanos))
}

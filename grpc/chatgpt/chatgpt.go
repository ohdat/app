package chatgpt

import (
	"context"
	"github.com/ohdat/opb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client opb.ChatgptClient
}

func NewChatgptClient(address string) (*GrpcClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &GrpcClient{
		conn:   conn,
		client: opb.NewChatgptClient(conn),
	}, nil
}

var once sync.Once

var grpcClient *GrpcClient

func GetChatgptClient() *GrpcClient {
	once.Do(func() {
		var address = viper.GetString("grpc.chatgpt")
		var err error
		grpcClient, err = NewChatgptClient(address)
		if err != nil {
			panic(err)
		}
	})
	return grpcClient
}

func (s *GrpcClient) Close() {
	s.conn.Close()
}

func (s *GrpcClient) SendMessage(wsId, message, ConversationId, ParentMessageId string) (err error) {
	var req = new(opb.ChatgptRequest)
	req.Message = message
	req.WsId = wsId
	req.ConversationId = ConversationId
	req.ParentMessageId = ParentMessageId
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var res *opb.ChatgptResponse
	res, err = s.client.SendMessage(ctx, req)
	if err != nil {
		return
	}
	res.Message = message
	return
}

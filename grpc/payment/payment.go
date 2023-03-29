package payment

import (
	"context"
	"errors"
	"github.com/ohdat/app/response"
	"github.com/ohdat/opb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client opb.PaymentClient
}

func NewPaymentClient(address string) (*GrpcClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &GrpcClient{
		conn:   conn,
		client: opb.NewPaymentClient(conn),
	}, nil
}

var once sync.Once

var grpcClient *GrpcClient

func GetPaymentClient() *GrpcClient {
	once.Do(func() {
		var address = viper.GetString("grpc.payment")
		var err error
		grpcClient, err = NewPaymentClient(address)
		if err != nil {
			panic(err)
		}
	})
	return grpcClient
}

func (s *GrpcClient) Close() {
	s.conn.Close()
}

func (s *GrpcClient) ProducerRecharge(wallet string, amount int) (err error) {
	var req = new(opb.ProducerRechargeRequest)
	var res *opb.ProducerRechargeReply
	req.WalletAddress = wallet
	req.Amount = int64(amount)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err = s.client.ProducerRecharge(ctx, req)
	if err != nil {
		return
	}
	return
	var code = int(res.GetCode())
	if code != 200 {
		if response.ErrCode(code).Code() != 0 {
			err = response.ErrCode(code)
		}
		err = errors.New(res.GetMessage())
	}
	return
}

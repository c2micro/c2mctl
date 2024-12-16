package service

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/c2micro/c2mcli/internal/middleware"
	"github.com/c2micro/c2mshr/defaults"
	managementv1 "github.com/c2micro/c2mshr/proto/gen/management/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var mgmtConn = &grpcConn{}

type grpcConn struct {
	ctx  context.Context
	conn *grpc.ClientConn
	svc  managementv1.ManagementServiceClient
}

// инициализация подключения к mgmt серверу
func Init(ctx context.Context, host string, token string) error {
	var err error
	mgmtConn.ctx = ctx

	if mgmtConn.conn, err = grpc.NewClient(
		host,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithUnaryInterceptor(middleware.UnaryClientInterceptor(token)),
		grpc.WithStreamInterceptor(middleware.StreamClientInterceptor(token)),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(defaults.MaxProtobufMessageSize),
			grpc.MaxCallSendMsgSize(defaults.MaxProtobufMessageSize),
		),
	); err != nil {
		return err
	}

	mgmtConn.svc = managementv1.NewManagementServiceClient(mgmtConn.conn)
	return nil
}

func Close() error {
	if mgmtConn.conn != nil {
		return mgmtConn.conn.Close()
	}
	return nil
}

func getSvc() managementv1.ManagementServiceClient {
	return mgmtConn.svc
}

func AddOperator(username string) (*managementv1.Operator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().NewOperator(ctx, &managementv1.NewOperatorRequest{Username: username})
	if err != nil {
		return nil, err
	}
	return rep.GetOperator(), nil
}

func ListOperators() ([]*managementv1.Operator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().GetOperators(ctx, &managementv1.GetOperatorsRequest{})
	if err != nil {
		return nil, err
	}
	return rep.GetOperators(), nil
}

func RegenOperator(username string) (*managementv1.Operator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().RegenerateOperator(ctx, &managementv1.RegenerateOperatorRequest{Username: username})
	if err != nil {
		return nil, err
	}
	return rep.GetOperator(), nil
}

func RevokeOperator(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := getSvc().RevokeOperator(ctx, &managementv1.RevokeOperatorRequest{Username: username})
	return err
}

func AddListener() (*managementv1.Listener, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().NewListener(ctx, &managementv1.NewListenerRequest{})
	if err != nil {
		return nil, err
	}
	return rep.GetListener(), nil
}

func ListListeners() ([]*managementv1.Listener, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().GetListeners(ctx, &managementv1.GetListenersRequest{})
	if err != nil {
		return nil, err
	}
	return rep.GetListeners(), nil
}

func RegenListener(id int64) (*managementv1.Listener, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().RegenerateListener(ctx, &managementv1.RegenerateListenerRequest{Lid: id})
	if err != nil {
		return nil, err
	}
	return rep.GetListener(), nil
}

func RevokeListener(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := getSvc().RevokeListener(ctx, &managementv1.RevokeListenerRequest{Lid: id})
	return err
}

func GetCertCA() (*managementv1.GetCertCAResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return getSvc().GetCertCA(ctx, &managementv1.GetCertCARequest{})
}

func GetCertOperator() (*managementv1.GetCertOperatorResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return getSvc().GetCertOperator(ctx, &managementv1.GetCertOperatorRequest{})
}

func GetCertListener() (*managementv1.GetCertListenerResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return getSvc().GetCertListener(ctx, &managementv1.GetCertListenerRequest{})
}

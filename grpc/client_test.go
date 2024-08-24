package grpc

import (
	"context"
	"github.com/stretchr/testify/require"
	servicegrpc "go20240218/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	cc, err := grpc.NewClient(":8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	client := servicegrpc.NewUserServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	resp, err := client.GetById(ctx, &servicegrpc.GetByIdRequest{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.User)
}

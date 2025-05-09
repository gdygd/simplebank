package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	db "github.com/gdygd/simplebank/db/sqlc"
	"github.com/gdygd/simplebank/token"
	"github.com/gdygd/simplebank/util"
	"github.com/gdygd/simplebank/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server

}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, username string, role string, duration time.Duration) context.Context {
	ctx := context.Background()
	accessToken, _, err := tokenMaker.CreateToken(username, role, duration)
	require.NoError(t, err)

	bearerToekn := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToekn,
		},
	}
	return metadata.NewIncomingContext(ctx, md)
}

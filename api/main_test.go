package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/slavik22/bank/db/sqlc"
	"github.com/slavik22/bank/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		SecretKey:           util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

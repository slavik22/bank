package util

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)

	if err != nil {
		log.Fatal(err)
	}

	require.NotEmpty(t, password)
	require.NoError(t, err)
	err = CheckPassword(password, hashedPassword)
}

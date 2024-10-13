package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/vyctor/go-finance-backend/utils"
	"testing"
)

func createRandomUser(t *testing.T) User {

	var userCreateArgs = CreateUserParams{
		Username: "usename",
		Password: "password",
		Email:    utils.RandomEmail(),
	}

	userId, err := testQueries.CreateUser(context.Background(), userCreateArgs)
	require.NoError(t, err)
	require.NotEmpty(t, userId)

	userFromDb, err := testQueries.GetUserById(context.Background(), userId)

	require.NoError(t, err)

	require.Equal(t, userCreateArgs.Username, userFromDb.Username)
	require.Equal(t, userCreateArgs.Password, userFromDb.Password)
	require.Equal(t, userCreateArgs.Email, userFromDb.Email)
	require.NotEmpty(t, userFromDb.CreatedAt)
	require.NotEmpty(t, userFromDb.UpdatedAt)

	return userFromDb
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	userFromDb, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.Equal(t, userFromDb.ID, user1.ID)
}

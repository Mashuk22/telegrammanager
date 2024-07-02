package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	roleName := "Regular"

	role, err := testQueries.GetRole(context.Background(), 1)
	if err == nil {
		role, err = testQueries.CreateRole(context.Background(), roleName)
		if err != nil {
			require.NoError(t, err)
		}
		require.NotEmpty(t, role.Name)
	}

	require.Equal(t, roleName, role.Name)

	userArg := CreateUserParams{
		ChatID:       12341234,
		Username:     sql.NullString{String: "username", Valid: true},
		FirstName:    sql.NullString{String: "firstname", Valid: true},
		LastName:     sql.NullString{String: "lastname", Valid: true},
		RoleID:       1,
		IsSubscribed: sql.NullBool{Bool: true, Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), userArg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, userArg.ChatID, user.ChatID)
	require.Equal(t, userArg.Username, user.Username)
	require.Equal(t, userArg.FirstName, user.FirstName)
	require.Equal(t, userArg.LastName, user.LastName)
	require.Equal(t, userArg.RoleID, user.RoleID)
	require.Equal(t, userArg.IsSubscribed, user.IsSubscribed)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}

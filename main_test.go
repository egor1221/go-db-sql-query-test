package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		require.NoError(t, err)
	}

	clientID := 1

	cl, err := selectClient(db, clientID)

	require.NoError(t, err)
	require.Equal(t, cl.ID, clientID)
	require.NotEmpty(t, cl.Birthday)
	require.NotEmpty(t, cl.Email)
	require.NotEmpty(t, cl.FIO)
	require.NotEmpty(t, cl.Login)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		require.NoError(t, err)
	}

	clientID := -1

	cl, err := selectClient(db, clientID)

	require.Equal(t, err, sql.ErrNoRows)
	assert.Empty(t, cl.Birthday)
	assert.Empty(t, cl.Email)
	assert.Empty(t, cl.FIO)
	assert.Empty(t, cl.ID)
	assert.Empty(t, cl.Login)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		require.NoError(t, err)
	}

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)

	cl.ID = id

	require.NoError(t, err)
	require.NotEmpty(t, id)

	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, cl, client)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		require.NoError(t, err)
	}

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
}

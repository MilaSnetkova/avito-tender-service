package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func setupContainerDatabase(t *testing.T) (testcontainers.Container, string) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image: "postgres:13", // или другая версия
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
		},
		ExposedPorts: []string{"5432/tcp"},
	}

	db, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	port, err := db.MappedPort(ctx, "5432")
	require.NoError(t, err)

	dbURL := "postgres://testuser:testpassword@localhost:" + port.Port() + "/testdb"

	return db, dbURL
}

func TestDatabase(t *testing.T) {
	dbContainer, dbURL := setupContainerDatabase(t)
	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	// Устанавливаем переменные окружения для теста
	os.Setenv("POSTGRES_CONN", dbURL)

	// Запускаем тесты с реальной базой данных
}

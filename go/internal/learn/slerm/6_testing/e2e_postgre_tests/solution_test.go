package e2e_postgre_tests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	db *pgx.Conn
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbName := "postgres"
	dbUser := "postgres"
	dbPassword := "postgres"

	pc, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:14-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	//req := testcontainers.ContainerRequest{
	//	Image:        "postgres:14-alpine",
	//	ExposedPorts: []string{"5432"},
	//	Env: map[string]string{
	//		"POSTGRES_DB":       dbName,
	//		"POSTGRES_USER":     dbUser,
	//		"POSTGRES_PASSWORD": dbPassword,
	//	},
	//	WaitingFor: wait.ForLog("Ready to accept connections"),
	//}
	//
	//pc, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
	//	ContainerRequest: req,
	//	Started:          true,
	//})
	//if err != nil {
	//	log.Fatalf("Could not start postgre: %s", err)
	//}
	//defer func() {
	//	if err := pc.Terminate(ctx); err != nil {
	//		log.Fatalf("Could not stop redis: %s", err)
	//	}
	//}()

	// Get container host and port
	host, err := pc.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get PostgreSQL container host: %s", err)
	}
	port, err := pc.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Could not get PostgreSQL container port: %s", err)
	}

	connString := fmt.Sprintf("user=postgres password=postgres dbname=postgres host=%s port=%s sslmode=disable", host, port.Port())
	db, err = pgx.Connect(ctx, connString)
	if err != nil {
		return
	}

	code := m.Run()
	db.Close(ctx)
	os.Exit(code)
}

func TestQueryExecution(t *testing.T) {
	pCtx := context.Background()

	ctx, cancel := context.WithTimeout(pCtx, 5*time.Second)
	defer cancel()

	// Insert test data
	_, err := db.Exec(ctx, "CREATE TABLE test_table (id SERIAL PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)
	_, err = db.Exec(ctx, "INSERT INTO test_table (name) VALUES ('test_name')")
	assert.NoError(t, err)

	// Perform query
	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM test_table").Scan(&count)
	assert.NoError(t, err)

	// Assert the result
	assert.Equal(t, 1, count)
}

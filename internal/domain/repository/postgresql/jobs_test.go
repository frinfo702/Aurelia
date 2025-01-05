package postgresql_test

import (
	"Aurelia/internal/domain/repository/postgresql"
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestFindAll(t *testing.T) {
	ctx := context.Background()
	pgContainer, err := postgres.Run(ctx,
		"postgres:14",
		postgres.WithInitScripts(filepath.Join(".", "testdata", "setupDB.sql")),
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(3*time.Minute),
		),
	)
	if err != nil {
		log.Fatalf("failed to start postgres container: %s", err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate mysql container: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	tdb, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)

	// ドキュメントはctxを渡している！どうする？！
	// 一度dbのまま作ってみて、あとでpgxに変更
	jobRepo := postgresql.NewJobRepository(tdb)
	assert.NoError(t, err)

	foundJobs, err := jobRepo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, foundJobs, 3)

}

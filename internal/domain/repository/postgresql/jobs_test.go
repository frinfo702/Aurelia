package postgresql_test

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	db  *sql.DB
	err error
)

// loadSQLFile は指定されたファイルパスのSQLを読み込み、返します。
func loadSQLFile(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to read SQL file: %w", err)
	}

	return string(content), nil
}

func TestMain(m *testing.M) {
	db, err := setupTestContainer()
	if err != nil {
		log.Fatalf("failed to setup test container: %v", err)
	}

	err = setupTestData()
	if err != nil {
		log.Fatalf("failed to setup test data: %v", err)
	}
	// テストの実行
	code := m.Run()

	err = teardownTestData()
	if err != nil {
		log.Printf("failed to teardown test data: %v", err)
	}

	// データベース接続のクローズ
	if err := db.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	}

	os.Exit(code)
}

func setupTestContainer() (*sql.DB, error) {
	ctx := context.Background()

	// PostgreSQLコンテナの設定
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(60),
	}

	// コンテナの起動
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("failed to start container: %v", err)
	}

	// テスト終了時にコンテナを終了
	defer func() {
		if err := postgresC.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %v", err)
		}
	}()

	// ホストとポートの取得
	host, err := postgresC.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get container host: %v", err)
	}

	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("failed to get container port: %v", err)
	}

	// データベース接続文字列の作成
	dsn := fmt.Sprintf("host=%s port=%s user=testuser password=testpass dbname=testdb sslmode=disable", host, port.Port())

	// データベース接続
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db, nil
}

func setupTestData() error {
	// setup.sql の実行
	setupSQL, err := loadSQLFile("./testdata/setupDB.sql")
	if err != nil {
		log.Fatalf("failed to load setup.sql: %v", err)
	}

	_, err = db.Exec(setupSQL)
	if err != nil {
		log.Fatalf("failed to execute setup.sql: %v", err)
	}
	return nil
}

func teardownTestData() error {
	// teardown.sql の実行
	teardownSQL, err := loadSQLFile("testdata/cleanupDB.sql")
	if err != nil {
		log.Printf("failed to load teardown.sql: %v", err)
	} else {
		_, err = db.Exec(teardownSQL)
		if err != nil {
			log.Printf("failed to execute teardown.sql: %v", err)
		}
	}
	return nil
}

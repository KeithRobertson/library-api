package migrations

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	// Step 1: connect without keyspace
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042

	adminSession, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("cannot connect to Cassandra: %v", err)
	}
	defer adminSession.Close()

	// Step 2: ensure keyspace exists
	createKeyspace(adminSession)

	// Step 3: create a NEW session bound to the keyspace
	cluster.Keyspace = "library"
	keyspaceSession, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("cannot connect to keyspace: %v", err)
	}
	defer keyspaceSession.Close()

	// Step 4: run migrations using the keyspace session
	driver, err := cassandra.WithInstance(keyspaceSession, &cassandra.Config{
		KeyspaceName: "library",
	})
	if err != nil {
		log.Fatalf("cannot create Cassandra driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"cassandra",
		driver,
	)
	if err != nil {
		log.Fatalf("migration init error: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}

func createKeyspace(session *gocql.Session) {
	if err := session.Query(`
    CREATE KEYSPACE IF NOT EXISTS library
    WITH replication = {
        'class': 'SimpleStrategy',
        'replication_factor': 1
    };
`).Exec(); err != nil {
		log.Fatalf("failed to create keyspace: %v", err)
	}
}

func NewAppSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	cluster.Keyspace = "library"
	cluster.Consistency = gocql.Quorum

	return cluster.CreateSession()
}

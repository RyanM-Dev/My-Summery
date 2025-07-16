### Key Points
- **Dockertest Overview**: Dockertest is a Go library that simplifies running integration tests by spinning up Docker containers for third-party services like databases, ensuring realistic and isolated test environments.
- **Ease of Use**: It seems likely that Dockertest’s straightforward API makes it easy to set up and tear down containers, reducing the complexity of testing against real services.
- **Cross-Platform**: The evidence leans toward Dockertest being effective on Windows, macOS, and Linux, providing consistent test environments across platforms.
- **No Controversy**: There’s no apparent debate around Dockertest; it’s widely regarded as a reliable tool for integration testing in Go.

### What is Dockertest?
Dockertest is a Go library that helps you run integration tests by launching Docker containers for services like MySQL, PostgreSQL, MongoDB, or Redis. Instead of mocking these services, you can test against real instances, which are created and destroyed automatically during tests. This approach ensures your tests reflect real-world behavior.

### Why Use It?
Dockertest is ideal when your application depends on external services. It eliminates the need for complex mocks, which can be hard to maintain, especially when service schemas or APIs change. By using Docker, you can quickly spin up isolated containers, run tests, and clean up afterward, making your tests reliable and repeatable.

### How to Get Started?
Install Dockertest with:
```sh
go get -u github.com/ory/dockertest/v3
```
Then, use it to run a container, connect to it, and clean up after tests. The examples below show how to set it up for various services.

### Basic Example
Here’s a simple example for testing with a MySQL container:
1. Create a Docker pool.
2. Run a MySQL container.
3. Wait for it to be ready.
4. Run your tests.
5. Clean up the container.

This process ensures your tests are isolated and don’t leave behind unused containers.

---



# 🚦 Comprehensive Guide to Dockertest for Go Integration Testing

Dockertest is a powerful Go library that simplifies integration testing by enabling developers to run Docker containers for third-party services like databases, message queues, or APIs during tests. As of July 13, 2025, it’s a go-to tool for Go developers who need to test applications against real services in a controlled, isolated environment. This guide provides a detailed exploration of Dockertest, complete with verbose examples, tips, and best practices to help you master the package. 🌟

## 📚 Introduction to Dockertest
Dockertest allows you to spin up Docker containers for services like MySQL, PostgreSQL, MongoDB, or Redis, run your tests against them, and tear them down afterward. This approach eliminates the need for complex mocking, which can be error-prone and hard to maintain when service schemas or APIs change. By using real services, Dockertest ensures your tests are realistic and reliable.

### 🌟 Why Use Dockertest?
- **Isolation**: Each test runs in its own container, preventing interference.
- **Realism**: Tests against actual services catch issues that mocks might miss.
- **Consistency**: Provides a uniform test environment across Windows, macOS, and Linux.
- **Ease of Use**: Simplifies container management with a straightforward API.

### 🎯 When to Use Dockertest?
- When your application interacts with external services like databases or APIs.
- For integration tests that need to verify real-world behavior.
- When testing different service versions or configurations.

## 🎯 Installation
To install Dockertest, run:

```sh
go get -u github.com/ory/dockertest/v3
```

This fetches the latest version (v3) of Dockertest, making the following packages available:
- `github.com/ory/dockertest/v3`
- `github.com/ory/dockertest/v3/docker`

Import it in your test code:

```go
import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)
```

## 🎯 Using Dockertest
Dockertest’s workflow involves creating a Docker pool, running a container, waiting for it to be ready, running tests, and cleaning up. The `pool.Retry` function is key to handling service startup delays, ensuring tests don’t fail prematurely.

### 📌 Basic Usage
Here’s a complete example for running a MySQL container:

```go
package dockertest_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// Create a pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// Verify Docker connection
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Run MySQL container
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Wait for the database to be ready
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// Run tests
	code := m.Run()

	// Clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSomething(t *testing.T) {
	// Example test: Query the database
	var result int
	err := db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		t.Fatalf("Query failed: %s", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
```

**How It Works**:
- **Setup**: `dockertest.NewPool("")` creates a pool using the default Docker endpoint.
- **Connection Check**: `pool.Client.Ping()` ensures Docker is accessible.
- **Container Creation**: `pool.Run` starts a MySQL 5.7 container with a root password.
- **Retry Logic**: `pool.Retry` waits until the database is ready by attempting to connect and ping.
- **Test Execution**: `m.Run()` runs the tests, and `TestSomething` queries the database.
- **Cleanup**: `pool.Purge(resource)` removes the container.

**Sample Failure Output**:
If the database connection fails:
```
Could not connect to database: dial tcp 127.0.0.1:3306: connect: connection refused
```

**💡 Tip**: Set `pool.MaxWait` (e.g., `pool.MaxWait = 120 * time.Second`) to allow more time for slow-starting services.

### 📌 Examples for Common Services
Dockertest supports various services. Below are detailed examples for popular ones:

#### 🎯 PostgreSQL
```go
package dockertestpsql_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var db *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	resource.Expire(120)

	pool.MaxWait = 120 * time.Second
	if err := pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestQuery(t *testing.T) {
	var result int
	err := db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		t.Fatalf("Query failed: %s", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
```

**How It Works**:
- Uses `RunWithOptions` for more control, setting `AutoRemove` to true.
- Sets a 120-second expiration with `resource.Expire(120)`.
- Connects to PostgreSQL using the `lib/pq` driver.

**💡 Tip**: Use `listen_addresses = '*'` to allow connections from outside the container.

#### 🎯 MongoDB
```go
package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "5.0",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=root",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = pool.Retry(func() error {
		var err error
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://root:password@localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), nil)
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := dbClient.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Could not disconnect MongoDB: %s", err)
	}

	os.Exit(code)
}

func TestMongoInsert(t *testing.T) {
	collection := dbClient.Database("testdb").Collection("testcoll")
	_, err := collection.InsertOne(context.TODO(), map[string]string{"key": "value"})
	if err != nil {
		t.Fatalf("Insert failed: %s", err)
	}
}
```

**How It Works**:
- Runs a MongoDB 5.0 container with authentication.
- Uses `mongo.Connect` to establish a connection.
- Disconnects the client after tests to free resources.

**💡 Tip**: Always disconnect MongoDB clients to avoid resource leaks.

#### 🎯 Redis
```go
package redis_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/ory/dockertest/v3"
)

var db *redis.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("redis", "3.2", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		db = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})
		return db.Ping().Err()
	}); err != nil {
		log.Fatalf("Could not connect to Redis: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestRedisSetGet(t *testing.T) {
	err := db.Set("key", "value", 0).Err()
	if err != nil {
		t.Fatalf("Set failed: %s", err)
	}
	val, err := db.Get("key").Result()
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}
	if val != "value" {
		t.Errorf("Expected value, got %s", val)
	}
}
```

**How It Works**:
- Runs a Redis 3.2 container without additional configuration.
- Uses the `go-redis/redis` client to interact with Redis.

**💡 Tip**: Use `resource.GetPort` to dynamically get the container’s port mapping.

## 🎯 Advanced Features
Dockertest supports complex testing scenarios, such as custom images, networks, and multiple containers.

### 📌 Building and Running Custom Images
You can build a custom Docker image from a Dockerfile:

```go
resource, err := pool.BuildAndRun("my-postgres-test-image", "./db/image/Dockerfile", []string{})
if err != nil {
	log.Fatalf("Could not start resource: %s", err)
}
```

**Dockerfile Example**:
```dockerfile
FROM postgres:latest
# Add custom configurations
ENV POSTGRES_DB custom_db
```

**How It Works**:
- `pool.BuildAndRun` builds the image from the specified Dockerfile and runs it.
- Useful for custom service configurations.

**💡 Tip**: Ensure the Dockerfile is in the correct context directory.

### 📌 Managing Networks
For multi-container setups, create a Docker network:

```go
network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
	Name: "app-datastore",
})
if err != nil {
	log.Fatalf("Could not create network: %s", err)
}
resource, err := pool.RunWithOptions(&dockertest.RunOptions{
	Repository: "postgres",
	Tag:        "latest",
	NetworkID:  network.ID,
}, func(config *docker.HostConfig) {
	config.NetworkMode = "bridge"
	config.AutoRemove = true
})
```

**How It Works**:
- Creates a network for containers to communicate.
- Assigns the network to containers via `NetworkID`.

**💡 Tip**: Use `findNetwork` to check for existing networks before creating new ones.

### 📌 Multiple Containers
Here’s an example of running an application and database container:

```go
func CreateLocalTestContainer() (*LocalTestContainer, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
		return nil, err
	}

	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{Name: "app-datastore"})
	if err != nil {
		log.Fatalf("Could not create network: %s", err)
	}

	dbresource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
		NetworkID: network.ID,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
	})
	if err != nil {
		log.Fatalf("Could not start dbresource: %s", err)
	}

	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s:%s/dbname?sslmode=disable", strings.Trim(dbresource.Container.Name, "/"), "5432")

	appresource, err := pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		Dockerfile: "Dockerfile",
		ContextDir: ".",
		Platform:   "linux/amd64",
	}, &dockertest.RunOptions{
		Name:      "app",
		Env:       []string{fmt.Sprintf("DB_CONN_URL=%s", databaseUrl)},
		NetworkID: network.ID,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
	})
	if err != nil {
		log.Fatalf("Could not start appresource: %s", err)
	}

	return &LocalTestContainer{
		appName:      appresource.Container.Name,
		dbName:       dbresource.Container.Name,
		appcontainer: appresource,
		dbcontainer:  dbresource,
		pool:         pool,
		network:      network.ID,
	}, nil
}
```

**How It Works**:
- Runs a PostgreSQL container and an application container.
- Uses a network to allow communication.
- Builds the application container from a Dockerfile.

**💡 Tip**: Use `pool.MaxWait` to handle long startup times for application containers.

## 🎯 Troubleshooting and FAQ
- **⚠️ Out of Disk Space**:
  Clean up unused Docker volumes and images with `docker-cleanup-volumes`.

- **⚠️ Removing Old Containers**:
  If containers persist, set an expiration:
  ```go
  resource.Expire(60) // Kill after 60 seconds
  ```
  Or use `AutoRemove`:
  ```go
  config.AutoRemove = true
  ```

- **⚠️ Connection Issues**:
  Increase `pool.MaxWait` or debug with `docker logs <container_id>`.

## 🎯 Best Practices
- **Use `pool.Retry`**: Always wait for services to be ready.
- **Set Timeouts**: Adjust `pool.MaxWait` for slow services.
- **Clean Up**: Use `pool.Purge` and `AutoRemove` to free resources.
- **Use Networks**: For multi-container tests, ensure proper communication.
- **Log Errors**: Use detailed logging to diagnose issues.

## 📋 Summary Table: Dockertest Features

| Feature                | Purpose                                      | Example                                      | When to Use                          |
|------------------------|----------------------------------------------|----------------------------------------------|---------------------------------------|
| Basic Container        | Run a single service container               | `pool.Run("mysql", "5.7", []string{...})`   | Simple integration tests             |
| Custom Images          | Build and run custom Docker images           | `pool.BuildAndRun("my-image", "./Dockerfile", []string{})` | Custom service configurations        |
| Networks               | Enable container communication               | `pool.Client.CreateNetwork(...)`            | Multi-container setups               |
| AutoRemove             | Auto-clean containers after tests            | `config.AutoRemove = true`                  | Prevent resource leaks               |
| Retry Mechanism        | Wait for service readiness                   | `pool.Retry(func() error { ... })`          | Handle startup delays                |

## 📚 Additional Resources
- **Official Documentation**: [Dockertest on GitHub](https://github.com/ory/dockertest)
- **Examples**: [Dockertest Examples](https://github.com/ory/dockertest/tree/master/examples)
- **Go Testing Guide**: [Go Testing Documentation](https://go.dev/doc/code#Testing)

This guide equips you with everything needed to master Dockertest for Go integration testing, ensuring robust and realistic tests. 🚀


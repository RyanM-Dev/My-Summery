### 📖 Chapter Summary

Chapter 4 transitions from theoretical gRPC concepts to the concrete implementation of a Go microservice using **Hexagonal Architecture** (Ports and Adapters). The primary goal is to build a loosely coupled, highly testable application by isolating the core business logic from external technical concerns like databases, gRPC stubs, and CLIs. This section outlines the architectural blueprint and the necessary project folder structure to ensure maintainability as the system scales.

#### 🆕 Key New Concepts & Technologies

|**Concept / Technology**|**Short Description**|**Why It Matters in gRPC Microservices**|
|:--|:--|:--|
|**Hexagonal Architecture**|A pattern that separates the application core from external dependencies via interfaces.|Ensures business logic remains agnostic of the transport layer (gRPC) or database (MySQL).|
|**Driver (Primary) Actors**|External entities that trigger the application (e.g., a gRPC client or CLI).|Identifies the entry points of your microservice system.|
|**Driven (Secondary) Actors**|External entities triggered by the application (e.g., a database or third-party API).|Manages out-of-process dependencies without polluting the business logic.|
|**Ports**|Interfaces defining how actors interact with the application.|Acts as a strict contract that must be implemented by adapters.|
|**Adapters**|Concrete implementations that translate technology-specific data into domain models.|Allows switching between different technologies (e.g., REST vs. gRPC) without changing core logic.|

#### 📌 Concepts Explained

- **Hexagonal Architecture Fundamentals**: Understanding the "Ports and Adapters" system.
- **Actor Classification**: Differentiating between Driver and Driven actors.
- **The Application Core**: Keeping business logic technology-agnostic.
- **The Ports & Adapters Layer**: Defining interfaces vs. concrete implementations.
- **Project Layout**: Organizing folders for maximum isolation.

### 🔧 Hexagonal Architecture (Ports and Adapters)

- **Clear explanation**: Proposed by Alistair Cockburn, this pattern aims to create components that connect via ports (interfaces) and adapters (concrete implementations). The application core sits in the center (the "Hexagon") and is surrounded by layers that handle external communication.
- **Importance in gRPC microservices**: It makes the microservice transport-neutral. If you need to add a REST endpoint alongside gRPC, you simply add a new adapter without touching the `PlaceOrder` logic.
- **Design Principles**: Use **Dependency Injection** to plug adapters into ports at runtime. This allows for "Pluggable" components, which are essential for swapping real databases with mocks during testing.

### 🔧 Driver vs. Driven Actors

- **Driver Actors**: These are "Primary" actors that initiate a conversation with the hexagon (e.g., a user making a gRPC request).
- **Driven Actors**: These are "Secondary" actors that the hexagon initiates a conversation with (e.g., the application saving an order to a MySQL database).
- **Relationship**: Drivers use **Driver Ports** to command the hexagon; the hexagon uses **Driven Ports** to command Driven Actors.

---

#### 💻 Code Examples & Implementation

**Code Example: Project Folder Structure (Book Version)** The book suggests a standard hierarchy to enforce the separation of concerns.

```
order/
├── cmd/
│   └── main.go           # Entry point & Dependency Injection
├── config/
│   └── config.go         # Environment configurations
├── internal/
│   ├── adapters/         # Concrete implementations (db, grpc)
│   ├── application/      # Core logic (api, domain)
│   └── ports/            # Interfaces (contracts)
└── go.mod                # Dependency management
```

**Updated Modern Version (2026 Best Practices)** In 2026, senior Go engineers often follow the **Standard Go Project Layout** more strictly, emphasizing the use of `pkg` for code that's safe for other projects to import, while keeping everything else in `internal`.

```
order/
├── cmd/
│   └── ordersvc/         # Renamed to service name for clarity
│       └── main.go
├── internal/
│   ├── adapter/          # Plural vs Singular (Modern preference varies)
│   │   ├── handler/      # Driver Adapters (gRPC handlers)
│   │   └── repository/   # Driven Adapters (DB repositories)
│   ├── application/      # Core business logic
│   │   ├── domain/       # Entities/Models
│   │   └── service/      # Use case orchestration
│   └── port/             # Interface definitions
├── proto/                # Local proto definitions if not centralized
├── buf.yaml              # Buf configuration (2026 Standard)
├── buf.gen.yaml          # Buf generation config
└── go.mod                # Go 1.24+ module
```

**Explanation:**

- **`internal/`**: Ensures that these packages cannot be imported by external projects, preventing "accidental" public APIs.
- **`cmd/`**: Only contains the "glue" code—loading config, initializing adapters, and starting the server.
- **`buf.yaml`**: Replaces raw `protoc` commands with a modern configuration-driven approach for gRPC generation.

---

#### 📊 Figures & Diagrams

**Figure 4.1: Hexagonal architecture with ports**

- **Description**: A visual representation showing the "Application" core in the center, flanked by "Driver" and "Driven" sides.
- **Flow**:
    - **Driver Side**: Web/CLI adapters use ports to reach the Application.
    - **Driven Side**: The Application uses ports to reach MySQL/Other App adapters.
- **2026 Clarification**: The hexagon shape is arbitrary; what matters is that the core (inner layer) has no knowledge of the outer layers (adapters).

**Figure 4.2: Project folder structure of a Go microservice**

- **Description**: A directory tree mapping hexagonal concepts to Go packages.
- **Relationships**:
    - `internal/application/core/domain` contains the pure business entities.
    - `internal/ports` contains the Go interfaces that define what the system _can_ do.
    - `internal/adapters/grpc` is where the gRPC-specific server code lives.

---

📋 Combined Practical Examples **The Architectural "Contract" (Port Definition)** In 2026, we ensure our ports are clearly defined before implementing any logic.

```
// internal/port/service.go
package port

import "github.com/username/microservices/order/internal/application/domain"

// OrderService is a Driver Port (The Application interface)
type OrderService interface {
    PlaceOrder(order domain.Order) (domain.Order, error)
}

// OrderRepository is a Driven Port (The Database interface)
type OrderRepository interface {
    Save(order *domain.Order) error
    Get(id string) (domain.Order, error)
}
```

💡 Pro Tips & Best Practices

1. **Keep the Hexagon Pure**: The code inside `internal/application` should **never** import `google.golang.org/grpc` or any SQL driver.
2. **Interface Segregation**: Don't create one massive "Port" interface. Separate them by responsibility (e.g., `OrderPort`, `PaymentPort`).
3. **Avoid Package Bloat**: Don't put everything in one package. Use `internal/adapters/db` for database logic and `internal/adapters/grpc` for gRPC logic to prevent circular dependencies.
4. **The `internal` Rule**: Use the `internal` directory for any code that is not intended to be a public library. Go's compiler strictly enforces this.
5. **Standard Layout**: While the book uses a specific layout, always check the [Golang Standards Project Layout](https://github.com/golang-standards/project-layout) for industry-wide conventions as of 2026.
6. **Dependency Injection (DI)**: In 2026, consider using lightweight DI containers like **Uber-fx** or **Google Wire** if your hexagon grows large, though manual injection in `main.go` is preferred for clarity in smaller services.

🧪 Practice & Review **Questions**

1. What is the main motivation for using Hexagonal Architecture in a microservice?
2. Define the difference between a Driver Actor and a Driven Actor.
3. Why are ports usually implemented as Interfaces in Go?
4. What role does the `adapters` folder play in this architecture?
5. Why should the `Application` layer be "technology-agnostic"?
6. Which folder usually contains the "Dependency Injection" logic and application entry point?
7. Explain the "Dependency Rule" in Hexagonal Architecture.

**Answers**

1. To build loosely coupled application components that are isolated from external dependencies like databases or UI, making them easier to test and maintain.
    
2. Driver Actors (Primary) trigger communication with the application (e.g., a CLI or gRPC request); Driven Actors (Secondary) are triggered _by_ the application (e.g., a database saving data).
    
3. Interfaces act as contracts. This allows the core logic to call methods without knowing the underlying implementation (like whether the database is MySQL or MongoDB).
    
4. It contains concrete implementations that translate technology-specific requests into domain-friendly formats that the application hexagon can understand.
    
5. To ensure that business rules are not tied to specific frameworks or protocols, allowing the service to evolve or switch technologies with minimal code changes.
    
6. The `cmd/` folder, specifically `main.go`, is where adapters are initialized and plugged into ports via dependency injection.
    
7. Outer layers (Adapters) always depend on inner layers (Application Core). The core never depends on anything outside itself.

### 📖 Section Summary

This section details the initial technical setup of a Go-based microservice following **Hexagonal Architecture** (Ports and Adapters). It covers the initialization of the Go module, the creation of a strictly isolated folder hierarchy, the definition of core business entities (domain models), and the establishment of **Ports** (interfaces) which serve as the strict contract between the application logic and external infrastructure.

#### 🆕 Key New Concepts & Technologies

|**Concept / Technology**|**Short Description**|**Why It Matters in gRPC Microservices**|
|:--|:--|:--|
|**Go Modules (`go.mod`)**|Go's built-in dependency management system.|Ensures reproducible builds and handles versioning of internal stubs and external libraries.|
|**Domain Models**|Pure Go structs representing business entities.|Keeps business rules independent of transport layers (gRPC) or persistence layers (SQL).|
|**Ports (Interfaces)**|Contractual definitions of application capabilities.|Enables loose coupling, allowing for pluggable adapters and easy mocking during tests.|
|**Dependency Injection**|Passing required dependencies into a struct constructor.|Decouples the application from concrete implementations, enhancing flexibility and testability.|
|**`internal` Directory**|A special Go directory that restricts package access.|Enforces strict architectural boundaries by preventing external modules from importing private logic.|

#### 📌 Concepts Explained

- **Project Initialization**: Establishing the module path and folder hierarchy.
- **Domain-Driven Design (Core)**: Modeling business entities like `Order` and `OrderItem`.
- **The Orchestration Layer (API)**: Implementing the service logic that coordinates use cases.
- **Port Definitions**: Defining the "Driver" and "Driven" contracts.

### 🔧 Project Initialization

- **Explanation**: Setting up a microservice requires a clean module path and a folder structure that supports the "Ports and Adapters" pattern to ensure modularity.
- **Microservices Context**: In 2026, module paths are strictly lowercase to avoid OS-level case sensitivity issues, and each service is isolated within an `internal` folder to prevent "leaking" core logic [165, History].
- **Design Principle**: Use the `cmd/` directory exclusively for the application's entry point and dependency injection "glue" code.

### 🔧 The Application Core (The Hexagon)

- **Explanation**: This is the heart of the service, containing domain models and API logic. It is "technology-agnostic," meaning it has no knowledge of gRPC or databases.
- **Theory**: The core receives requests through **Driver Ports** and communicates with external systems (like databases) through **Driven Ports**.

### 🔧 Port Interfaces

- **Explanation**: Ports are Go interfaces that define interactions between actors and the application. **Driver Ports** define what the application offers; **Driven Ports** define what the application requires from infrastructure.
- **Principle**: By depending on interfaces rather than concrete implementations, the application remains "pluggable".

#### 💻 Code Examples & Implementation

**Code Example: Project Initialization (Book Version)**

```
# Initialize the project
mkdir -p microservices/order
cd microservices/order
go mod init GitHub.com/<username>/microservices/order

# Create folder hierarchy
mkdir cmd mkdir config
mkdir -p internal/adapters/db
mkdir -p internal/adapters/grpc
mkdir -p internal/application/core/api
mkdir -p internal/application/core/domain
mkdir -p internal/ports
```

**Updated Modern Version (2026 Best Practices)**

```
# Go 1.24+ standard initialization (Lowercase paths)
mkdir -p order && cd order
go mod init github.com/username/microservices/order

# 2026 Layout: Emphasis on Service terminology
mkdir -p cmd/ordersvc           # Explicit binary name
mkdir -p internal/adapter/db    # Modern preference for singular
mkdir -p internal/adapter/grpc
mkdir -p internal/application/domain
mkdir -p internal/application/service  # Renamed 'api' to 'service'
mkdir -p internal/port          # Consistent naming
```

**Explanation:**

- **Lowercase Paths**: GitHub repository names and Go module paths should be lowercase to prevent "module not found" errors in case-sensitive environments [History].
- **Flat Names**: Modern Go projects often prefer `port` and `adapter` (singular) over plurals for cleaner package imports.

---

**Code Example: Domain Model Definition (Book Version)**

```
// internal/application/core/domain/order.go
package domain

type Order struct {
    ID         int64       `json:"id"`
    CustomerID int64       `json:"customer_id"`
    Status     string      `json:"status"`
    OrderItems []OrderItem `json:"order_items"`
    CreatedAt  int64       `json:"created_at"`
}
```

**Updated Modern Version (2026 Best Practices)**

```
// internal/application/domain/order.go
package domain

import "time"

type Order struct {
    ID         int64       `json:"id"`
    CustomerID int64       `json:"customer_id"`
    Status     string      `json:"status"`
    OrderItems []OrderItem `json:"order_items"`
    CreatedAt  time.Time   `json:"created_at"` // 2026: Use time.Time for entities
}

// NewOrder handles initial domain state
func NewOrder(custID int64, items []OrderItem) Order {
    return Order{
        CreatedAt:  time.Now().UTC(), // Modern standard: UTC by default
        Status:     "Pending",
        CustomerID: custID,
        OrderItems: items,
    }
}
```

**Explanation:**

- **Rich Types**: Senior Go engineers prefer `time.Time` over `int64` for domain entities to leverage Go's native time manipulation and validation methods [History].

---

**Code Example: Port Definitions (Book Version)**

```
// internal/ports/api.go
type APIPort interface {
    PlaceOrder(order domain.Order) (domain.Order, error)
}

// internal/ports/db.go
type DBPort interface {
    Get(id string) (domain.Order, error)
    Save(*domain.Order) error
}
```

**Updated Modern Version (2026 Best Practices)**

```
// internal/port/service.go
package port

import (
    "context"
    "github.com/username/microservices/order/internal/application/domain"
)

// OrderService defines what the core can do (Driver Port)
type OrderService interface {
    // 2026 Standard: Context is mandatory for gRPC-aware services
    PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}

// OrderRepository defines what the core needs (Driven Port)
type OrderRepository interface {
    Get(ctx context.Context, id int64) (domain.Order, error)
    Save(ctx context.Context, order *domain.Order) error
}
```

**Explanation:**

- **Context Propagation**: Every port method in a modern gRPC service **must** include `context.Context` to allow for cancellation signals and tracing to flow from the gRPC handler through the core to the database.

#### 📊 Figures & Diagrams

**Figure 4.2: Project folder structure of a Go microservice**

- **Description**: A directory tree mapping hexagonal architecture concepts to Go packages.
- **Key Components**:
    - `cmd/main.go`: The entry point for dependency injection.
    - `internal/ports`: Interfaces defining method signatures.
    - `internal/application`: The business logic heart.
- **2026 Improvement**: The structure is slightly flatter, with "service" replacing "core/api" to align with modern DDD (Domain Driven Design) nomenclature.

📋 Combined Practical Examples **The 2026 Core Service Implementation**

```
// internal/application/service/order_service.go
package service

import (
	"context"
	"github.com/username/microservices/order/internal/application/domain"
	"github.com/username/microservices/order/internal/port"
)

type Application struct {
	repo port.OrderRepository // Dependency on the port
}

func NewApplication(repo port.OrderRepository) *Application {
	return &Application{repo: repo}
}

func (a *Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	// Persistence through the driven port
	err := a.repo.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
```

💡 Pro Tips & Best Practices

1. **Context is Non-Negotiable**: Never define a port method without `context.Context`. You will need it later for gRPC timeouts.
2. **Lowercase Module Paths**: Always use lowercase for your module URLs during `go mod init` to avoid "module not found" errors on different operating systems [History].
3. **Encapsulate State**: Keep your domain entities pure; logic related to status transitions (e.g., setting an order to "Pending") should live in the core, not the database adapter.
4. **Interface Ownership**: In Hexagonal Architecture, the core (the consumer) owns and defines the Driven Port interfaces (`DBPort`), not the adapter.
5. **Use `internal` for Safety**: Strictly use the `internal/` directory for any package not intended for public consumption to prevent accidental API leakage.
6. **Avoid Global State**: Do not use global variables for database connections. Use constructor-based dependency injection to pass adapters into the core.

🧪 Practice & Review **Questions**

1. Why is the `internal` directory critical for microservice architectural integrity?
2. Define the difference between a Driver Port and a Driven Port.
3. Why should domain models use `time.Time` instead of `int64` timestamps in a modern Go project?
4. Explain how Dependency Injection allows for "pluggable" database adapters.
5. What is the significance of the `go mod init` module URL?
6. Why must every Port method include `context.Context` in a gRPC-centric system?
7. What happens if you try to import a package from an `internal` folder of another Go module?

**Answers**

1. It prevents other modules from importing private logic, ensuring that the microservice's internal implementation details remain encapsulated and cannot be "leaked."
    
2. A Driver Port (Primary) is the interface the core offers to trigger business logic; a Driven Port (Secondary) is what the core uses to interact with external infrastructure like a database.
    
3. `time.Time` provides a robust, semantic API for timezones, durations, and comparisons, whereas `int64` is just a raw number that loses business context.
    
4. By injecting an interface (Port) into the Application Core, you can swap concrete adapters (MySQL, MongoDB, or Mocks) at runtime without modifying the business logic.
    
5. It establishes the base URL for the project's packages, which Go uses to resolve internal imports and allows external consumers to find the module.
    
6. It allows the propagation of deadlines, cancellation signals, and distributed trace IDs across service boundaries.
    
7. The Go compiler will **forbid the import and throw an error** during the build process [History, 209].Go enforces a strict rule for any directory named `internal`: its contents are private and can only be imported by packages located within the same directory tree. This mechanism is essential in microservices for several reasons:
   - **Encapsulation**: It hides implementation details, ensuring that external modules cannot rely on your private logic.
   - **Preventing "Spaghettification"**: It stops uncontrolled and chaotic dependencies between different packages and teams.
   - **Maintenance**: You can refactor or change code inside `internal` freely without worrying about breaking someone else's external project

### 📖 Section Summary (Combined from both books)

This section details the implementation of **Driven Adapters**, specifically focusing on the persistence layer using Go. While the primary book (Babal) leverages **GORM** to rapidly build a MySQL database adapter, the secondary book (Shuiskov) provides a broader perspective on the **Repository Pattern**, emphasizing the importance of keeping database-specific logic isolated and ensuring that data entities remain distinct from domain models for better maintainability.

#### 🆕 Key New Concepts & Technologies

|**Concept / Technology**|**Short Description**|**Primary Book View (Babal)**|**Secondary Book View (Shuiskov)**|**Why It Matters**|
|:--|:--|:--|:--|:--|
|**GORM**|A popular Object-Relational Mapping (ORM) library for Go.|Used as the primary tool for database abstraction.|Briefly mentions ORMs but emphasizes clean interface-based repositories.|Accelerates development by automating SQL query generation.|
|**Repository Pattern**|An abstraction that mediates between the domain and data mapping layers.|Implicitly implemented through the DB adapter.|Explicitly defined as a core layer of the microservice.|Decouples the application core from the persistence technology.|
|**Data Mapping**|Converting data between different structures (DB Entity $\leftrightarrow$ Domain Model).|Essential for translating GORM structs to pure domain models.|Stresses explicit mapping functions to avoid leaking DB tags into business logic.|Maintains the integrity of Hexagonal Architecture boundaries.|
|**AutoMigrate**|Automatic schema synchronization.|Used for quick table creation during development.|Prefers explicit migration files for production-grade deployments [History].|Simplifies local development and initial prototyping.|

#### 📌 Concepts Explained

- **The Database Adapter**: Implementing the concrete logic for the `DBPort`.
- **Infrastructure Isolation**: Using GORM to handle the complexities of MySQL without leaking them to the core.
- **Model Transformation**: Why we must map database entities back to domain models.

### 🔧 The Database Adapter (Driven Adapter)

- **Explanation**: In Hexagonal Architecture, the database adapter is a **Driven Adapter**. It implements an interface (Port) defined by the application core. Babal implements this using a struct that wraps a `*gorm.DB` instance.
- **Shuiskov Nuance**: Shuiskov treats this as a "Repository" layer. He highlights that while in-memory repositories are great for tests, persistent repositories (like MySQL) must handle **durability**—ensuring data survives a service restart.
- **2026 Approach**: Modern Go engineers often use **sqlc** for type-safe SQL or **GORM** for complex relations, but always wrap these in a repository that accepts `context.Context` to support gRPC-level timeouts.

---

### 🔧 Model Mapping & Translation

- **Explanation**: Database entities often contain technology-specific tags (like `gorm:"primaryKey"`) or types (like `gorm.Model`). These should never reach the application core.
- **Implementation**: The adapter's `Get` and `Save` methods are responsible for "unmarshalling" database records into clean domain structs.
- **Importance**: This allows you to change your database schema (e.g., renaming a column) by only updating the mapping logic in the adapter, without touching a single line of business logic.

---

#### 💻 Code Examples & Implementation

**Code Example: DB Adapter Structure (Primary Book Version)**

```
// From Babal, page 56
type Adapter struct {
    db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
    db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
    if openErr != nil {
        return nil, fmt.Errorf("db connection error: %v", openErr)
    }
    err := db.AutoMigrate(&Order{}, OrderItem{})
    return &Adapter{db: db}, nil
}
```

**Code Example: Repository Interface (Secondary Book Version)**

```
// Shuiskov prefers explicit interfaces on the receiving side
type metadataRepository interface {
    Get(ctx context.Context, id string) (*model.Metadata, error)
    Put(ctx context.Context, id string, metadata *model.Metadata) error
}
```

**Updated Modern Version (2026 Best Practices)**

```
package repository

import (
	"context"
	"github.com/username/order/internal/application/domain"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

// 2026: Constructors return the interface (Port)
func NewGormAdapter(db *gorm.DB) *GormAdapter {
	return &GormAdapter{db: db}
}

func (a *GormAdapter) Save(ctx context.Context, order *domain.Order) error {
	// 2026 Standard: Use WithContext for tracing and cancellation
	return a.db.WithContext(ctx).Create(order).Error
}
```

**Explanation**:

- **`WithContext(ctx)`**: Mandatory in 2026. This allows a gRPC timeout or a client cancellation to propagate all the way to the database, killing expensive queries early.
- **Mapping Logic**: Notice the `Save` method directly accepts the domain model, but internally, GORM might use an `OrderEntity` struct to handle DB tags.

---

#### 📊 Figures & Diagrams

**Figure 4.2: Project folder structure (Revisited)**

- **Description**: Shows the `internal/adapters/db` folder containing the concrete implementation.
- **Differences**: Babal places the mapping logic directly inside the methods. Shuiskov suggests using helper functions like `MetadataToProto` (or in this case `ToDomain`) to keep the methods clean and readable.

---

📚 Further Reading from "Microservices with Go, 2nd Edition"

1. **Chapter 7: Storing Service Data**: A deep dive into MySQL integration and the trade-offs between different database types.
2. **Durability and Persistence**: Understanding why the move from in-memory repositories to external databases is the first step toward production-grade reliability.
3. **Data Caching Techniques**: How to use an in-memory repository (like the one in Babal) as a **cache** layer in front of a real database to reduce latency.

🔨 Challenge from the secondary book: **Implement a "Read-Through Cache"**. Modify your `OrderService` (The Core) to depend on _two_ repositories: a primary MySQL adapter and a secondary in-memory cache adapter. If an order isn't in the cache, fetch it from MySQL and populate the cache.

---

💡 Pro Tips & Best Practices (Combined)

1. **Context is King**: Always pass `context.Context` from the gRPC handler through the core to the DB adapter. In 2026, failing to do so is considered a senior-level mistake.
2. **Don't Over-ORM**: Use GORM for basic CRUD, but don't be afraid to use raw SQL for complex reports. GORM's `Raw()` function is your friend.
3. **Internal Module Paths**: Ensure your imports use the full module path (e.g., `github.com/user/project/internal/...`) to avoid package resolution issues [History].
4. **AutoMigrate is for Dev Only**: In production (2026), use migration tools like `golang-migrate` or `atlas` to manage schema changes safely.
5. **Errors as Constants**: Define your "Not Found" errors in the `port` or `repository` package so both the core and the adapter can check them using `errors.Is()`.
6. **Zero Trust Mapping**: Never trust a database struct in your business logic. Map it to a domain struct immediately upon retrieval.

---

🧪 Practice & Review **Questions**

1. Why does the `Adapter` struct wrap `*gorm.DB` instead of having the core logic use GORM directly?
2. What is the role of `AutoMigrate` in the `NewAdapter` function?
3. Explain why we convert `orderEntity.CreatedAt` (from DB) to a different format for the domain model.
4. How does Shuiskov's view on "Durability" justify the use of an external database over an in-memory map?
5. Why should a `Get` method in an adapter return a domain model instead of a GORM model?
6. What happens if the `gorm.Open` call fails during application startup?

**Answers**

1. To ensure the business logic remains technology-agnostic. The core should not know that you are using GORM or MySQL.
2. It ensures that the required database tables exist and match the struct definitions, simplifying the initial setup.
3. Database types are often tied to specific storage formats (like `time.Time` with specific tags). Domain models need clean Go types (like `int64` or pure `time.Time`) for easier logic processing.
4. In-memory data is lost when the service restarts or crashes. External databases provide durable storage that persists across failures.
5. To maintain architectural boundaries. If you return a GORM model, you "leak" the database implementation details into your core logic.
6. The `NewAdapter` function returns an error, which should cause the `main.go` file to call `log.Fatalf` and stop the service from running in an invalid state.

### 📖 Section Summary (Combined from both books)

This section covers the transformation of a technology-agnostic application core into a network-accessible service via a **gRPC Driver Adapter**. While the primary book (Babal) focuses on the mechanics of setting up the listener and registering the server, the secondary book (Shuiskov) provides a more robust approach to **data mapping** and **rich error handling**, ensuring that the gRPC layer remains a thin translation boundary that doesn't leak transport concerns into the business logic.

#### 🆕 Key New Concepts & Technologies

|**Concept / Technology**|**Short Description**|**Primary Book View (Babal)**|**Secondary Book View (Shuiskov)**|**Why It Matters**|
|:--|:--|:--|:--|:--|
|**gRPC Server Reflection**|Allows clients to query the server for its schema.|Enabled in development for `grpcurl` support.|Implicitly supports it for easier manual testing.|Simplifies debugging without needing local `.proto` files.|
|**Unimplemented Server**|An embedded struct for forward compatibility.|Mandatory to embed `UnimplementedOrderServer`.|Required to ensure the server satisfies the interface.|Prevents breaking the build when adding new RPC methods.|
|**TCP Listener**|The network socket that accepts incoming traffic.|Created using `net.Listen("tcp", ...)`.|Part of the `main` executable initialization.|Fundamental for making the service reachable over the network.|
|**Request Mapping**|Translating protobuf messages to domain entities.|Performed inline within the handler methods.|Prefers dedicated `ToProto` and `FromProto` helpers.|Decouples the API contract from the internal business models.|

#### 📌 Concepts Explained

- **The gRPC Adapter Structure**: Defining the bridge between the Port and the Network.
- **Server Lifecycle (Listen & Serve)**: Managing the startup and socket binding.
- **Reflection & Debugging**: Using `grpcurl` and server reflection for API discovery.
- **Translation Logic**: Mapping generated Protobuf code to pure Go domain models.

---

### 🔧 The gRPC Adapter implementation

- **Explanation**: The gRPC adapter is a **Driver Adapter**. It implements the gRPC server interface (generated from `.proto`) and "drives" the application core by calling the `APIPort`.
- **Shuiskov Nuance**: Shuiskov emphasizes that the handler should return **rich gRPC status codes** (e.g., `codes.InvalidArgument`) rather than raw errors to help clients understand exactly what went wrong.
- **2026 Approach**: In modern Go, we use **Buf** to manage the generated code and ensure the adapter is instrumented with interceptors for logging, tracing, and metrics from day one.

### 🔧 Server Lifecycle (Listen & Serve)

- **Mechanism**: Bind a TCP port (e.g., `:3000`), initialize `grpc.NewServer()`, register the service, and begin handling requests.
- **Production Tip**: Babal suggests enabling **Reflection** only in the "development" environment. In 2026, we strictly enforce this to prevent exposing API internals in production.

---

#### 💻 Code Examples & Implementation

**Code Example: gRPC Adapter Structure (Primary Book Version)**

```
// From Babal, page 58
type Adapter struct {
    api  ports.APIPort
    port int
    order.UnimplementedOrderServer // Forward compatibility
}
```

**Code Example: Mapping Strategy (Secondary Book / Shuiskov Version)**

```
// Shuiskov prefers dedicated mapping functions (p. 107)
func OrderToProto(o *domain.Order) *gen.Order {
    return &gen.Order{
        Id: o.ID,
        Status: o.Status,
        // ... mapping logic
    }
}
```

**Updated Modern Version (2026 Best Practices)**

```
package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/username/order/internal/application/domain"
	"github.com/username/order/gen" // Generated by Buf
)

type Adapter struct {
	api  port.OrderService
	gen.UnimplementedOrderServiceServer
}

// 2026: Handler returns gRPC specific status codes
func (a *Adapter) Create(ctx context.Context, req *gen.CreateOrderRequest) (*gen.CreateOrderResponse, error) {
	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	// Translation layer
	order, err := a.api.PlaceOrder(ctx, domain.Order{
		CustomerID: req.UserId,
		// ... mapping
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to place order: %v", err)
	}

	return &gen.CreateOrderResponse{OrderId: order.ID}, nil
}
```

**Explanation**:

- **`GetUserId()`**: Always use the generated getter methods to avoid nil-pointer panics if the request message is nil.
- **`status.Error`**: Translates internal logic errors into standard gRPC wire errors.
- **`context.Context`**: Mandatory propagation for deadlines and tracing.

---

#### 📊 Figures & Diagrams

**Figure 4.2: Project folder structure (Revisited)**

- **Description**: This figure highlights the `internal/adapters/grpc` folder.
- **Key Files**:
    - `grpc.go`: Contains the RPC method implementations (The "Handlers").
    - `server.go`: Contains the `Run()` logic and gRPC server initialization.
- **Cross-book comparison**: Babal combines the server and handlers in one package, while Shuiskov often labels these as `handler` and `controller` to separate networking from orchestration.

---

📚 Further Reading from "Microservices with Go, 2nd Edition"

1. **Chapter 5: Synchronous Communication**: A masterclass in gRPC vs. HTTP/JSON and why code generation is the "killer feature" for microservices.
2. **Context Propagation**: How to pass metadata (like Trace IDs) across service boundaries using Go contexts.
3. **Advanced Error Details**: Using `google.rpc.BadRequest` to return multi-field validation errors.

🔨 Challenge from the secondary book: **Implement a Global Recovery Interceptor.** In your `Run()` method, add a gRPC interceptor that catches any `panic` within your handlers and converts it into a `codes.Internal` gRPC error. This prevents your entire service from crashing if one request hits a bug.

---

💡 Pro Tips & Best Practices (Combined)

1. **Use `grpcurl` for Manual Testing**: It acts like `curl` for gRPC. Use the `-plaintext` flag for local dev.
2. **Getter Methods are Life Savers**: Protobuf getters (e.g., `req.GetItems()`) safely return defaults if the field is missing, preventing crashes [History].
3. **Environment-Aware Reflection**: Register `reflection.Register(grpcServer)` inside an `if config.GetEnv() == "development"` block only.
4. **Context Propagation**: Always pass the `ctx` from the gRPC handler down into the core service and then to the database adapter.
5. **Standardize Internal Error Codes**: Map domain errors to gRPC codes (e.g., `domain.ErrNotFound` $\rightarrow$ `codes.NotFound`) inside the adapter.
6. **Prefer Buf over Raw Scripts**: Use `buf generate` to avoid maintaining brittle `run.sh` scripts for your `.proto` files [History].

---

🧪 Practice & Review **Questions**

1. Why must you embed `Unimplemented<Service>Server` in your adapter struct?
2. What is the role of `net.Listen` in the gRPC server lifecycle?
3. How does `reflection` assist a developer using `grpcurl`?
4. Why should data mapping (Proto $\leftrightarrow$ Domain) happen in the adapter and not the core?
5. What happens if you call `grpcServer.Serve(listen)` before registering your service?
6. Which gRPC status code should you return if a user sends a missing mandatory field?
7. Explain the difference between `grpc.NewServer()` and `grpcServer.Serve()`.

**Answers**

1. It provides forward compatibility, allowing your code to compile even if you add new RPC methods to the `.proto` file but haven't implemented them in Go yet.
2. It creates a TCP listener on a specific port to accept incoming network connections before the gRPC server starts processing them.
3. It allows `grpcurl` to "discover" the available services and methods on a running server without needing a local copy of the `.proto` file.
4. To keep the application core technology-agnostic. The core should only know about Go domain models, not generated protobuf structs.
5. The server will start, but every request will return an "Unimplemented" error because no logic was linked to the incoming calls.
6. `codes.InvalidArgument` (equivalent to HTTP 400).
7. `NewServer()` initializes the server object and its settings/interceptors, while `Serve()` is a blocking call that starts the actual processing of network traffic.

### 📖 Chapter Summary (Combined from both books)

This concluding section of Chapter 4 focuses on the "Glue Layer"—the assembly of isolated hexagonal components into a functional, running microservice. It details the implementation of **Dependency Injection (DI)** and configuration management, moving from Babal’s environment-based setup to Shuiskov’s more robust, production-ready scaffolding. You will learn how to wire repositories to the application core and then to the transport adapters, culminating in the first manual verification using `grpcurl`.

#### 🆕 Key New Concepts & Technologies

|**Concept / Technology**|**Short Description**|**Primary Book View (Babal)**|**Secondary Book View (Shuiskov)**|**Why It Matters**|
|:--|:--|:--|:--|:--|
|**Dependency Injection**|Passing dependencies into a struct rather than creating them inside.|Manual wiring in `main.go`.|Foundational for testable layers.|Ensures loose coupling and allows for easy mocking in tests.|
|**12-Factor Config**|Storing configuration in the environment.|Strict `os.Getenv` usage.|Prefers YAML files mapped to structs.|Keeps credentials and env-specific settings out of source code.|
|**Main Package**|The entry point of the Go binary.|Resides in `cmd/` to wire layers.|Emphasizes consistency across services.|The central place where the system is bootstrapped.|
|**`grpcurl`**|A CLI tool to interact with gRPC servers.|Used for "cURL-like" debugging.|Essential for white-box verification.|Allows testing gRPC endpoints without a dedicated client.|
|**Server Reflection**|Exposing service metadata to clients.|Enabled for development only.|Implicitly supported for discovery.|Enables `grpcurl` to discover methods without `.proto` files.|

---

#### 📌 Concepts Explained

- **Application Bootstrapping**: Initializing infrastructure (DB, Logging) before logic.
- **The Wiring Pattern**: How adapters are plugged into ports via DI.
- **Environment vs. File Configuration**: Best practices for 2026.
- **Manual Endpoint Verification**: Using CLI tools for gRPC testing.

---

### 🔧 Dependency Injection & Bootstrapping

- **Babal’s Approach**: Manual initialization. You create the `dbAdapter`, pass it to `NewApplication`, then pass the application to `NewGRPCAdapter`.
- **Shuiskov Nuance**: Emphasizes using **Inversion of Control (IoC)** patterns to keep the `main` package clean.
- **2026 Practice**: While manual wiring is standard for small services, senior devs use **Uber-Fx** or **Google Wire** to handle large dependency graphs and lifecycle management (Start/Stop hooks).

### 🔧 Configuration Management

- **Explanation**: A microservice must behave differently in `development` (plaintext, logging) vs `production` (TLS, structured logs).
- **Design Principle**: Never hardcode database URLs or ports. In 2026, we strictly separate configuration files (YAML/TOML) from secrets (Vault/Env).

---

#### 💻 Code Examples & Implementation

**Code Example: Configuration Management**

|Primary Book (Babal)|Secondary Book (Shuiskov)|2026 Modern Best Practice|
|:--|:--|:--|
|`os.Getenv("PORT")`|`yaml.NewDecoder(f).Decode(&cfg)`|`env.Parse(&cfg)` (Struct Tags)|

**Updated Modern Version (2026 Best Practices)**

```
// config/config.go
package config

import (
	"github.com/caarlos0/env/v11" // 2026 standard for env-to-struct mapping
)

type Config struct {
	Env            string `env:"ENV" envDefault:"development"`
	DBURL          string `env:"DATA_SOURCE_URL,required"`
	Port           int    `env:"APPLICATION_PORT" envDefault:"3000"`
	PaymentService string `env:"PAYMENT_SERVICE_URL"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
```

---

**Code Example: Application Entry Point (`main.go`)**

```
// Primary Book Version (Babal) - p. 62
func main() {
    dbAdapter, _ := db.NewAdapter(config.GetDataSourceURL())
    app := api.NewApplication(dbAdapter)
    grpcAdapter := grpc.NewAdapter(app, config.GetApplicationPort())
    grpcAdapter.Run()
}

// 2026 Modern Version (Combined with Shuiskov's Lifecycle Logic)
func main() {
    // 1. Load Config
    cfg, _ := config.Load()

    // 2. Initialize Infrastructure (Driven Adapters)
    dbAdapter, _ := db.NewGormAdapter(cfg.DBURL)

    // 3. Initialize Core (Hexagon)
    orderSvc := service.NewApplication(dbAdapter)

    // 4. Initialize Handlers (Driver Adapters)
    grpcAdapter := grpc.NewAdapter(orderSvc, cfg.Port)

    // 5. Start with Graceful Shutdown (Modern Requirement)
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    go grpcAdapter.Run()

    <-ctx.Done() // Wait for SIGINT/SIGTERM
    grpcAdapter.Stop()
}
```

**Explanation**:

- **`signal.NotifyContext`**: Ensures the service closes DB connections and stops accepting gRPC requests gracefully before exiting.
- **Wiring**: Notice the dependency flow: `DB` $\rightarrow$ `Application` $\rightarrow$ `gRPC`.

---

### 🔧 Verification with `grpcurl`

- **Command (Babal)**: `grpcurl -d '{...}' -plaintext localhost:3000 Order/Create`.
- **Secondary Book Tip**: Shuiskov recommends using `grpcurl` to verify the **Error Model**.
- **2026 Context**: We often use **Buf Curl** (`buf curl`) or Postman/Insomnia's built-in gRPC support for a better UI experience.

---

#### 📊 Figures & Diagrams

**Figure 4.2: Project folder structure (Revisited)**

- **Description**: The diagram illustrates the completed service where `main.go` sits at the top, pointing into `internal/adapters` and `internal/application`.
- **Shuiskov Difference**: Shuiskov suggests adding a `pkg/` directory for any domain models or errors you wish to export to other services, while keeping the implementation in `internal/`.

---

📚 Further Reading from "Microservices with Go, 2nd Edition"

1. **Chapter 8: Setting Up Service Deployments**: Covers how to containerize the `main.go` binary using multi-stage Docker builds.
2. **Chapter 11: Reliability Overview**: A deep dive into the **Graceful Shutdown** pattern implemented in our 2026 version of `main.go`.
3. **Twelve-Factor App Configuration**: Detailed theory on why environment variables are safer than hardcoded strings.

🔨 Challenge from the secondary book: **Implement a Health Check Endpoint.** Modify your gRPC adapter to register the standard gRPC Health service. Use `grpcurl` to query the service status (`grpc.health.v1.Health/Check`) and ensure it returns `SERVING` only after the DB adapter has successfully connected.

---

💡 Pro Tips & Best Practices (Combined)

1. **Dependency Inversion**: Always inject the **Port (Interface)** into your Application core, never the concrete Adapter.
2. **Fail Fast**: If the `db.NewAdapter` call fails, use `log.Fatalf` immediately. Do not start the gRPC server if its dependencies are missing.
3. **Environment-Specific Reflection**: Reflection is a security risk. Register it _only_ if `cfg.Env == "development"`.
4. **Context Propagation**: The `main.go` file should provide the "root" context. Every adapter and service method should accept this context to support cancellations.
5. **Standardize Naming**: Stick to Shuiskov's naming (`controller`, `repository`, `handler`) across your entire company to reduce mental overhead when switching between services.
6. **Buf for Generation**: Use `buf generate` instead of long `protoc` commands in your build scripts to ensure consistent stubs for your adapters.

---

🧪 Practice & Review **Questions**

1. Why is `main.go` called the "Glue" layer of Hexagonal Architecture?
2. Explain why we use `context.Context` in the gRPC server's `Serve()` call.
3. What is the benefit of mapping environment variables to a `Config` struct during startup?
4. How does `grpcurl` "know" which methods exist on a server without having the `.proto` file?
5. Why should the database URL be a configuration parameter rather than a hardcoded string?
6. Explain the role of `UnimplementedOrderServer` in the gRPC adapter struct.
7. What is the difference between `grpc.NewServer()` and `grpcServer.Serve(listen)`?

**Answers**

1. It is where the technology-agnostic core is "glued" to technology-specific adapters (DB, gRPC) via Dependency Injection.
2. It allows for graceful shutdown and propagation of cancellation signals across all running goroutines.
3. It provides type safety, allows for default values, and prevents the logic from being littered with `os.Getenv` calls.
4. It uses **gRPC Server Reflection**, which allows the server to broadcast its API schema to the client.
5. To support the "12-Factor App" principle, allowing the same binary to run in dev, staging, and production by simply changing environment variables.
6. It provides forward compatibility, ensuring the adapter still satisfies the server interface even if new methods are added to the proto but not yet implemented in Go.
7. `NewServer()` initializes the server settings and interceptors; `Serve()` starts the blocking loop that listens on the network socket.
# 🚀 Swagger (OpenAPI) Integration in Go — Complete Guide

Let's organize your Swagger setup clearly, tailored specifically for **Go backend projects** using popular frameworks like **Gin** or **Fiber**.

---

## 🎯 **Objective**

Implement clean, professional, and **auto-generated Swagger documentation** for your Go APIs. We'll use structured comments, serve the Swagger UI at a convenient endpoint (`/docs`), and follow best practices.

---

## 🛠 **Tools & Installation**

First, pick the right tool:

📦 **`swaggo/swag`**  
This library reads specially formatted Go comments and generates Swagger-compatible JSON.

- Official repo: [swaggo/swag](https://github.com/swaggo/swag)
    

Install it along with middleware and UI assets:

```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger   # for Gin
go get -u github.com/gofiber/swagger      # for Fiber
go get -u github.com/swaggo/files         # Swagger UI assets
```

Ensure `$GOPATH/bin` is in your `PATH`. Then run:

```bash
swag init
```

✅ **Generated files**:

- `docs/docs.go`
    
- `docs/swagger.json`
    
- `docs/swagger.yaml`
    

---

## 🔖 **Basic Setup in Go**

### 📝 **1. Add Top-Level API Comments (in `main.go`)**

Always add this section at the very top of your main file:

```go
// @title       My Awesome API
// @version     1.0
// @description API server for my cool application
// @host        localhost:8080
// @BasePath    /api/v1

// Authentication (optional)
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

✅ **Explanation**:

- **`@BasePath`**: Prefix for your API routes.
    
- **Security definitions**: Sets global authentication settings (Bearer token).
    

---

### 🌐 **2. Initialize Swagger in Router**

**Example for Gin**:

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "your_project/docs"
)

func main() {
    r := gin.Default()
    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}
```

**Example for Fiber**:

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger"
    _ "your_project/docs"
)

func main() {
    app := fiber.New()
    app.Get("/docs/*", swagger.HandlerDefault)
    app.Listen(":8080")
}
```

---
# 📘 Swagger `@` Annotation Tags in Go — Deep Dive with Examples 🧠

This guide is crafted like a **senior Go developer teaching you Swagger (`swaggo`) annotations** — tag by tag, with clear purpose, real use cases, and **at least two working examples** each. Think of it as your **Swagger mastery checklist**.

---

## ✅ 1. `@Summary`

### 🔍 What it is:

A **one-line summary** of the endpoint — shown at the top in Swagger UI.

### 🧪 Syntax:

```go
// @Summary <short sentence>
```

### 📌 Examples:

```go
// @Summary Get a single user by ID
// @Summary Create a new blog post
```

### 💡 Best Practices:

- ✅ 5–10 words max
    
- ✅ Use **action verbs** like "Get", "Create", "Delete"
    

---

## ✅ 2. `@Description`

### 🔍 What it is:

A longer, optional description. Appears below the summary in the UI.

### 🧪 Syntax:

```go
// @Description Line 1
// @Description Line 2 (optional)
```

### 📌 Examples:

```go
// @Description Fetches a user from the database.
// @Description Requires a valid ID and token.

// @Description Uploads a file to cloud storage.
// @Description File must be under 5MB.
```

### 💡 Tips:

- Supports multiple lines
    
- Markdown formatting is often supported (e.g., **bold**, `code`)
    

---

## ✅ 3. `@Tags`

### 🔍 What it is:

Defines **logical groups** for endpoints in Swagger UI (side navigation).

### 🧪 Syntax:

```go
// @Tags users,admin
```

### 📌 Examples:

```go
// @Tags auth
// @Tags orders,internal
```

### 💡 Tips:

- ✅ Use multiple tags if endpoint fits multiple contexts
    
- ⚠️ Don’t over-tag — clutter reduces clarity
    

---

## ✅ 4. `@Accept`

### 🔍 What it is:

Defines which **Content-Type(s)** the endpoint accepts.

### 🧪 Syntax:

```go
// @Accept json
```

### 📌 Examples:

```go
// @Accept json
// @Accept multipart/form-data
```

### 💡 Common Values:

- `json`
    
- `xml`
    
- `plain`
    
- `multipart/form-data`
    

---

## ✅ 5. `@Produce`

### 🔍 What it is:

Defines the **response Content-Type** the server produces.

### 🧪 Syntax:

```go
// @Produce json
```

### 📌 Examples:

```go
// @Produce json
// @Produce plain
```

---

## ✅ 6. `@Param`

### 🔍 What it is:

Specifies input parameters — from path, query, headers, body, or formData.

### 🧪 Syntax:

```go
// @Param <name> <in> <type> <required> "<description>"
```

### 📌 Examples:

#### 🟢 Path

```go
// @Param id path string true "User ID"
```

#### 🔵 Query

```go
// @Param search query string false "Search keyword"
// @Param limit query int false "Result limit"
```

#### 🟣 Header

```go
// @Param Authorization header string true "Bearer token"
```

#### 🟠 Body

```go
// @Param body body dto.CreateUser true "Create user payload"
```

#### 🔴 FormData

```go
// @Param avatar formData file true "Profile picture"
// @Param name formData string true "User's full name"
```

---

## ✅ 7. `@Success`

### 🔍 What it is:

Defines **successful responses**, including status and return type.

### 🧪 Syntax:

```go
// @Success <status> {type} <model> "<description>"
```

### 📌 Examples:

```go
// @Success 200 {object} dto.User "User details"
// @Success 201 {object} map[string]string "New ID"
```

### 🔢 Common Types:

- `object`
    
- `array`
    
- `string`
    
- `file`
    
- `map[string]string`
    

---

## ✅ 8. `@Failure`

### 🔍 What it is:

Describes expected **error responses**.

### 🧪 Syntax:

```go
// @Failure <status> {type} <model> "<description>"
```

### 📌 Examples:

```go
// @Failure 400 {object} errs.Error "Validation failed"
// @Failure 401 {object} errs.Error "Unauthorized"
```

---

## ✅ 9. `@Router`

### 🔍 What it is:

Declares the route path and HTTP method. Must match actual implementation.

### 🧪 Syntax:

```go
// @Router /path [method]
```

### 📌 Examples:

```go
// @Router /users/{id} [get]
// @Router /products [post]
// @Router /orders/{order_id} [delete]
```

⚠️ **Path variables** (e.g., `{id}`) must be declared in `@Param` too.

---

## ✅ 10. `@Security`

### 🔍 What it is:

Marks the endpoint as requiring security/authentication.

### 🧪 Syntax:

```go
// @Security <name>
```

### 📌 Global Definition:

In `main.go`:

```go
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

### 📌 Usage Example:

```go
// @Security BearerAuth
```

✅ This enables the **"Authorize"** button in Swagger UI.

---

## ✅ 11. `@ID` _(Optional)_

### 🔍 What it is:

Gives the route a **unique identifier**. Used for SDK generation, testing, and documentation tools.

### 🧪 Syntax:

```go
// @ID unique-operation-name
```

### 📌 Examples:

```go
// @ID get-user-by-id
// @ID upload-profile-image
```

---

## 🧪 Full Annotated Example (Fiber)

```go
// GetUserByID returns a user by ID
// @Summary Retrieve user
// @Description Retrieves a user’s data using a unique ID
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} dto.User "User object"
// @Failure 400 {object} errs.Error "Invalid ID"
// @Failure 404 {object} errs.Error "User not found"
// @Security BearerAuth
// @Router /users/{id} [get]
// @ID get-user-by-id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error { ... }
```

---

## 📋 Summary Table: All `@` Tags at a Glance

|Tag|Role|Required?|
|---|---|---|
|`@Summary`|Short, one-line endpoint description|✅|
|`@Description`|Detailed explanation (multi-line supported)|✅|
|`@Tags`|Group endpoint(s) in Swagger UI|✅|
|`@Accept`|Input content type(s)|✅|
|`@Produce`|Output content type(s)|✅|
|`@Param`|Input details (path/query/header/body/formData)|✅|
|`@Success`|Success responses|✅|
|`@Failure`|Failure/error responses|✅|
|`@Router`|Path and method declaration|✅|
|`@Security`|Marks route as requiring auth|❌|
|`@ID`|Unique identifier (code generation/test hooks)|❌|

---

## 🚀 Next Steps

Would you like me to:

✅ Generate a Swagger-ready Go boilerplate for **Fiber** or **Gin**?  
✅ Build a DTO struct and `errs.Error` sample for real?  
✅ Show you a side-by-side view of code and Swagger UI?

Just say the word — and I’ll tailor it to your actual project.

## 🧑‍💻 **Best Practices**

🌟 **Use DTO structs with example tags**:

```go
type Genre struct {
    ID          string `json:"id" example:"123"`
    Name        string `json:"name" example:"Jazz"`
    Description string `json:"description" example:"Relaxing genre"`
}
```

🌟 **Group APIs logically using tags** (`@Tags genres, collections`).

🌟 **Automate Swagger generation** (`go:generate`, Makefile).

🌟 **Always document errors** (`@Failure`) clearly.

---

## 🛡 **Authentication Setup**

**Global Setup (`main.go`)**:

```go
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

**Endpoint Usage**:

```go
// @Security BearerAuth
```

In Swagger UI, you'll see an **Authorize** button allowing you to enter your token.

---

## 🔎 **Testing Your Swagger UI**

- Launch server (`go run main.go`).
    
- Visit:
    

```
http://localhost:8080/docs/index.html
```

Interact with your APIs through the clean UI.

---

## 📋 **Quick Reference Cheat Sheet**

|Annotation|Description|Example|
|---|---|---|
|`@Summary`|Short title|`Get genre by ID`|
|`@Description`|Detailed description|`Retrieve detailed genre info`|
|`@Tags`|API grouping|`genres, admin`|
|`@Accept`|Request content-type|`json`|
|`@Produce`|Response content-type|`json`|
|`@Param`|Inputs|`id path string true "Genre ID"`|
|`@Success`|Successful response|`200 {object} dto.Genre`|
|`@Failure`|Error response|`404 {object} errs.Error`|
|`@Security`|Authentication requirement|`BearerAuth`|
|`@Router`|Route definition|`/genres/{id} [get]`|

---

## 🌟 **Example (Fiber)**

Fully documented handler:

```go
// GetGenreByID retrieves genre by ID.
// @Summary Get genre
// @Description Fetches a genre with detailed information.
// @Tags genres
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Success 200 {object} dto.Genre
// @Failure 400 {object} errs.Error
// @Failure 404 {object} errs.Error
// @Security BearerAuth
// @Router /genres/{id} [get]
func (h *GenreHandler) GetGenreByID(c *fiber.Ctx) error { ... }
```

---

## 💡 **Next Steps**

- Structure and document your handlers.
    
- Add security.
    
- Automate Swagger generation.
    

You're all set to build clear, maintainable, and well-documented Go APIs!

# 🎯 Using Multiple `@Tags` in Swagger (OpenAPI) Correctly ✅

Here's your clean and structured Markdown edit for **Obsidian**, enriched with emojis for enhanced readability, loyalty to your provided data, and useful tips and examples:

---

## ✅ What Happens if You Use Multiple `@Tags` in Swagger?

When you specify **more than one tag** in the `@Tags` directive, like this:

```go
// @Tags users,admin
```

It means:

- ⚡ **Your endpoint appears under multiple sections in Swagger UI**.
    

---

## 📚 Clear Example:

If your Swagger UI sidebar looks like this:

```
▼ users
   ├─ GET /users/{id}
   └─ POST /users

▼ admin
   └─ GET /users/{id}  ⬅️ (same endpoint appears again)
```

👉 The endpoint tagged as `users,admin` appears in **both** sections.

---

## 🎖️ **Practical Use Cases**

|🛠️ Use Case|📝 Example|
|---|---|
|🔑 Endpoint appears in **user-facing** and **admin** areas.|`@Tags users,admin`|
|📦 Group under both **inventory** and **products**|`@Tags inventory,products`|
|🚀 API that’s both **internal** and **external**|`@Tags internal,external`|

---

## 📌 **Example in Fiber or Gin**:

Clearly show how you’d practically tag an admin-level delete endpoint:

```go
// DeleteUser deletes a user by ID.
// @Summary Delete a user
// @Description Admin-level deletion of a user account.
// @Tags users,admin
// @Param id path string true "User ID"
// @Success 204 {string} string "No Content"
// @Security BearerAuth
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error { ... }
```

🔖 **Tip**:

- Ensure you only use multiple tags when logical and meaningful, keeping your documentation intuitive and tidy.
    

---

## 🚩 **Caution**:

⚠️ Avoid overusing tags to prevent a cluttered and confusing UI:

❌ **Bad Example**:

```go
// @Tags api,admin,users,system,internal,external,debug
```

✅ **Good Example** (focused, minimal):

```go
// @Tags users,admin
```

---

## 🔑 **Quick Cheat Sheet**:

|📌 Behavior|📋 Result|
|---|---|
|🟢 One tag|✅ Clean grouping under single tag|
|🔵 Multiple tags|✅ Endpoint visible in each specified tag group|
|🔴 Too many tags|❌ UI clutter, confusion|

---

## 🧠 **Advanced Tip (Optional)**:

👉 You can customize the order or descriptions of tags directly in your generated Swagger YAML or annotations to further optimize your UI.

_Example (Swagger YAML):_

```yaml
tags:
  - name: users
    description: User-related operations
  - name: admin
    description: Admin-only endpoints
```

✅ **Benefit**: Clearer documentation organization.

---

🎓 **Example Tagging Structure (Fiber/Gin)**:

```go
// UpdateProduct updates product details.
// @Summary Update a product
// @Description Update details of a specific product.
// @Tags products,inventory
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param body body dto.ProductUpdate true "Updated product data"
// @Success 200 {object} dto.Product "Updated product details"
// @Failure 400 {object} errs.Error "Invalid input"
// @Failure 404 {object} errs.Error "Product not found"
// @Security BearerAuth
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error { ... }
```

---

✅ **Summary (TL;DR)**:

- ✅ **Good**: Use multiple tags logically.
    
- ❌ **Bad**: Excessive tagging.
    

Keep it clean and meaningful!


# 🎯 Complete Guide to `@Param` in Swagger for Go (`swaggo`) ✅

The `@Param` directive is one of the **most powerful** and **versatile** tools in Swagger. It defines how clients **send input to your API**, from path variables to query params, headers, full JSON bodies, and even file uploads!

---

## 🧱 Basic Syntax

```go
@Param <name> <in> <type> <required> "<description>"
```

|🔹 Field|📝 Meaning|
|---|---|
|`<name>`|Name of the parameter (e.g., `id`, `Authorization`)|
|`<in>`|Location: `path`, `query`, `header`, `body`, `formData`|
|`<type>`|Data type: `string`, `int`, `bool`, `file`, or a struct (only for `body`)|
|`<required>`|`true` or `false` — whether this field is mandatory|
|`<description>`|A short, clear human-readable explanation|

---

## 🔹 1. Path Parameters

📌 Used for route parts like `/users/{id}`

```go
// @Param id path string true "User ID"
```

⚠️ **Must be marked `true` (required)** — path variables are always required.

🧪 Example combo:

```go
// @Router /users/{id} [get]
// @Param  id path string true "User ID"
```

---

## 🔹 2. Query Parameters

📌 Used for URLs like `/users?page=2&limit=10`

```go
// @Param page query int false "Page number"
// @Param search query string false "Search keyword"
```

✅ Optional params? Use `false` as the fourth field.

---

## 🔹 3. Header Parameters

📌 For things like `Authorization`, `X-Request-ID`, or custom headers

```go
// @Param Authorization header string true "Bearer token"
```

🎯 Great for token-based authentication or request tracing.

---

## 🔹 4. Body Parameters

📌 For **full JSON body** payloads — use with POST/PUT

```go
// @Param body body dto.CreateUserRequest true "User create payload"
```

📦 DTO Struct Example:

```go
type CreateUserRequest struct {
    Name string `json:"name" example:"John"`
    Age  int    `json:"age" example:"30"`
}
```

✨ Add `example` tags to improve Swagger UI visualization.

---

## 🔹 5. FormData Parameters

📌 For file uploads or multipart forms

```go
// @Param avatar formData file true "User avatar image"
// @Param name formData string true "Full name"
```

⚠️ Must combine with:

```go
// @Accept multipart/form-data
```

---

## 🧪 Full Example — All Param Types Together

```go
// UploadProfile handles profile photo uploads
// @Summary Upload profile
// @Description Upload a user’s profile image with metadata
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "User ID"
// @Param name formData string true "Full name"
// @Param avatar formData file true "Profile image"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} errs.Error
// @Router /users/{id}/upload [post]
// @Security BearerAuth
```

---

## 🧠 Tips & Best Practices

✅ **Always match `@Param` path fields with `@Router` path variables**  
✅ Use **clear, short, and descriptive names**  
✅ Use DTOs for `body` and include `example` values  
✅ Use `true/false` properly for required vs optional  
✅ Swagger only **documents** — actual validation is your responsibility in handlers

---

## 📘 Optional: Supported Primitive Types

|Swagger Type|Go Type|Example|
|---|---|---|
|`string`|`string`|`"John"`|
|`int`|`int64`|`42`|
|`bool`|`bool`|`true` or `false`|
|`file`|`[]byte`|multipart upload|

---

## 🧪 Want More?

Let me know if you want:

- 📁 A DTO struct template with examples
    
- 🧩 A sample handler using all param types
    
- 📄 Preview of the Swagger UI output for validation
    

I'll tailor examples to **Fiber**, **Gin**, or your specific project structure.

---
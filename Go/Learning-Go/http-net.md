
## **HTTP Client:

---

## 🌐 Avoid Using `http.DefaultClient` in Production

### ❌ Why you shouldn't use `http.DefaultClient`:

```go
resp, err := http.Get("https://example.com") // Uses DefaultClient internally
```

- Go’s `net/http` package provides a default client instance called `DefaultClient`.
    
- **By default, it has no timeout** set, which is risky in production environments.
    
- It's better to create a custom client with a proper timeout:
    

```go
client := http.Client{
    Timeout: 10 * time.Second,
}
```

---

## 📦 Understanding the `http.Response` Structure

When you send a request and get a response, the `http.Response` object includes several important fields:

|Field|Description|
|---|---|
|`StatusCode`|Numeric status code (e.g., 200, 404)|
|`Status`|Text version of the status code (e.g., "200 OK")|
|`Header`|Response headers (`http.Header` type)|
|`Body`|Response body (`io.ReadCloser` type)|

### 🧪 Example: Reading JSON response

```go
decoder := json.NewDecoder(resp.Body)
var result MyStruct
err := decoder.Decode(&result)
```

---

## ⚠️ Avoid Using `http.Get`, `http.Post`, etc.

```go
http.Get("https://example.com")   // Not recommended
http.Post(...)
```

- These helper functions use the **default client**, which means they **don't have a timeout**.
    
- It's safer to use your own `http.Client` with custom settings.


## **HTTP Server:


Here’s your new set of Go (Golang) notes rewritten in clean, well-structured **English Markdown** with summaries and examples:

---

## 🧭 Writing HTTP Responses in the Correct Order

When writing an HTTP response using Go’s `net/http` package, you must follow a specific sequence of method calls:

### ✅ Step-by-step order:

1. **(Optional)** Call `Header()` if you want to set custom response headers:
    

```go
w.Header().Set("Content-Type", "application/json")
```

2. **Call `WriteHeader(statusCode)`** to set the HTTP status code:
    

```go
w.WriteHeader(http.StatusCreated) // e.g., 201
```

- All standard HTTP status codes are defined as constants in the `net/http` package (e.g., `http.StatusOK`, `http.StatusBadRequest`).
    
- **Note:** These are **untyped integers**, not a custom type.
    

3. **Call `Write([]byte)`** to write the response body:
    

```go
w.Write([]byte(`{"message": "success"}`))
```

> If you're returning a **200 OK** status (default), you can skip the `WriteHeader` step — Go will use 200 if it's not explicitly set.

---

## ⏱️ Setting Server Timeouts (Important for Security)

To prevent slowloris-style attacks or poor client behavior, always set **timeouts** for the server:

### 🔧 Recommended timeout settings:

Use `time.Duration` values for:

- `ReadTimeout` — max duration for reading the entire request
    
- `WriteTimeout` — max duration before timing out the response write
    
- `IdleTimeout` — max amount of time to wait for the next request when keep-alives are enabled
    

### 🧪 Example:

```go
server := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}
```

> ⚠️ By default, Go's server does **not** apply any timeouts — you must set them manually to protect against malicious or broken clients.

---

# 🌐 Understanding HTTP Components: Payload, Headers, Query Parameters & More

HTTP communication is made up of **distinct components**, each serving a specific purpose in structuring, transporting, or processing data between a client and a server.

---

## 📦 Payload (Request Body)

The **payload** is the **main content/data** being sent in a request (especially in POST/PUT/PATCH methods).

```go
payload := strings.NewReader("{\n  \"for_super_followers_only\": false,\n  \"nullcast\": false}")
```

- **Purpose**: Carries the actual data (e.g., JSON) to be processed by the server.
    
- **Used in**: `POST`, `PUT`, `PATCH`
    
- **Formats**: JSON, XML, form data, plain text, binary
    
- **Location**: In the **body** of the HTTP request
    
- **Use Case**: Sending user data in a sign-up form.
    

---

## 🧾 Headers

**Headers** provide **metadata** about the request or response.

```go
req.Header.Add("Authorization", "Bearer token123")
req.Header.Add("Content-Type", "application/json")
```

- **Purpose**: Inform the server how to interpret or handle the request.
    
- **Common Headers**:
    
    - `Content-Type`: Format of the payload
        
    - `Authorization`: Auth token (e.g., Bearer)
        
    - `User-Agent`: Info about the client software
        
    - `Accept`: What formats the client can handle
        
- **Location**: Inside the **header section** of the HTTP message
    
- **Use Case**: Authenticate, define payload format, define client capabilities
    

---

## 🔎 Query Parameters

**Query parameters** filter or modify the request and are included in the **URL**.

```go
url := "https://api.twitter.com/2/tweets?count=10&include_entities=true"
```

- **Format**: `?key1=value1&key2=value2`
    
- **Purpose**: Apply filters, sorting, or control pagination
    
- **Location**: After `?` in the **URL**
    
- **Use Case**: Search or paginate data
    

---

## 🛣️ Path Parameters

Path parameters are **part of the URL path**, typically representing **resource IDs**.

```go
url := "https://api.twitter.com/2/users/123456/tweets"
```

- **Purpose**: Uniquely identify a resource
    
- **Location**: Inside the **path** section of the URL
    
- **Use Case**: Retrieve tweets for a specific user ID
    

---

## 🔗 URL (Uniform Resource Locator)

The **URL** defines the **location of the resource** you want to interact with.

```go
url := "https://api.twitter.com/2/tweets"
```

- **Components**:
    
    - Protocol: `https://`
        
    - Host: `api.twitter.com`
        
    - Path: `/2/tweets`
        
    - Optional query string: `?key=value`
        
- **Use Case**: Access a specific endpoint
    

---

## 🔨 HTTP Method

The **HTTP method** defines the type of action to perform on a resource.

```go
req, _ := http.NewRequest("POST", url, payload)
```

- **Common Methods**:
    
    - `GET`: Read data
        
    - `POST`: Create a resource
        
    - `PUT`: Replace/update a resource
        
    - `PATCH`: Partially update
        
    - `DELETE`: Remove a resource
        

---

## 🆚 Key Differences

### Payload vs. Query Parameters

|Payload|Query Parameters|
|---|---|
|Sent in the body|Sent in the URL|
|Good for complex or sensitive data|Good for filters/options|
|Not used with `GET`|Common with `GET`|

---

### Headers vs. Payload

|Headers|Payload|
|---|---|
|Meta-information|Actual content|
|Authorization, Content-Type|JSON, form-data|

---

### Path Parameters vs. Query Parameters

|Path Parameters|Query Parameters|
|---|---|
|Identify specific resource|Customize the response|
|Part of the path|Appended after `?`|

---

## 🧪 Complete Example

```go
// URL with path + query parameters
url := "https://api.twitter.com/2/users/123456/tweets?count=10&include_entities=true"

// Request body (payload)
payload := strings.NewReader(`{"text": "Hello world"}`)

// Create HTTP request
req, _ := http.NewRequest("POST", url, payload)

// Add headers
req.Header.Add("Authorization", "Bearer token123")
req.Header.Add("Content-Type", "application/json")
```

---

## ✅ Summary

|Component|What It Does|Where It Appears|
|---|---|---|
|**Payload**|Carries request/response data|Body|
|**Headers**|Provides metadata|Header section|
|**Query Params**|Filters/customizes request|URL (after `?`)|
|**Path Params**|Identifies specific resource|URL path|
|**Method**|Defines operation type|HTTP method|
|**URL**|Locates the target resource|Full request address|

---

Let me know if you'd like this as a cheat sheet or printable guide.
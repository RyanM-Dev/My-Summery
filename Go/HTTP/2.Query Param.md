
## **Query Parameters Deep Dive in Go**

---

## 1. What are Query Parameters? 🌐

- **Definition**: Anything after `?` in a URL, carrying key/value pairs for filtering, pagination, sorting, etc.
  
  ```text
  https://api.example.com/items?category=books&limit=10&sort=price%20desc


- **Syntax**:
    
    ```text
    ?key1=value1&key2=value2&key3=value%20with%20spaces
    ```
    
    - Keys/values **must** be URL-encoded ( space → `%20`, `:` → `%3A`, …)
        
    - Repeat keys for multiple values:
        
        ```text
        ?tag=go&tag=web&tag=api
        ```
        

---

## 2. Why URL-Encode? 🔒

```go
import "net/url"

u := url.URL{
    Scheme: "https",
    Host:   "api.example.com",
    Path:   "/search",
}
q := u.Query()
q.Add("q", "golang tips & tricks")
u.RawQuery = q.Encode()
// → https://api.example.com/search?q=golang+tips+%26+tricks
```

- **Avoid misinterpretation** of spaces, `&`, `=` and other special chars
    
- Always use `url.Values` + `Encode()`
    

---

## 3. Building Query Parameters “By Hand” 🛠️

### 3.1 Simple Single Parameters

```go
u, _ := url.Parse("https://api.example.com/items")
q := u.Query()
q.Add("limit", "25")
q.Add("sort", "created_at")
u.RawQuery = q.Encode()
// → https://api.example.com/items?limit=25&sort=created_at
```

### 3.2 Optional or Defaulted Values

```go
func AddLimit(q url.Values, limit int) {
    if limit > 0 {
        q.Add("limit", fmt.Sprint(limit))
    } else {
        q.Add("limit", "50") // default
    }
}
```

### 3.3 Multiple Values / Arrays

- **Repeated keys**:
    
    ```go
    func AddTags(q url.Values, tags []string) {
        for _, tag := range tags {
            q.Add("tag", tag)
        }
    }
    // → ?tag=go&tag=web&tag=api
    ```
    
- **CSV style**:
    
    ```go
    q.Add("tags", strings.Join(tags, ","))
    // → ?tags=go,web,api
    ```
    

---

## 4. Complex “Search” Queries 🔍

1. **Collect sub-clauses** (`from:`, `#`, …)
    
2. **Join** with `AND`/`OR` and parentheses
    
3. **Add** as single `query` param
    

```go
func BuildSearchQuery(usernames, hashtags []string, useAnd bool) string {
    var parts []string
    for _, u := range usernames {
        parts = append(parts, "from:"+u)
    }
    for _, h := range hashtags {
        parts = append(parts, "#"+strings.TrimPrefix(h, "#"))
    }
    op := "OR"
    if useAnd { op = " " }
    return strings.Join(parts, " "+op+" ")
}

// Usage:
queryStr := BuildSearchQuery(
  []string{"alice","bob"},
  []string{"news"},
  false,
// → "from:alice OR from:bob OR #news"
)
```

---

## 5. Refactoring: Helpers & Builders 🧩

### 5.1 Generic `AddParam`

```go
func AddParam(q url.Values, key string, val interface{}) {
    switch v := val.(type) {
    case string:
        if v != "" { q.Add(key, v) }
    case int:
        if v > 0 { q.Add(key, fmt.Sprint(v)) }
    case bool:
        q.Add(key, fmt.Sprint(v))
    case []string:
        for _, s := range v {
            if s != "" { q.Add(key, s) }
        }
    default:
        s := fmt.Sprint(v)
        if s != "" { q.Add(key, s) }
    }
}
```

### 5.2 Builder Pattern

```go
type QueryBuilder struct {
    q url.Values
}

func NewQueryBuilder() *QueryBuilder {
    return &QueryBuilder{q: make(url.Values)}
}

func (b *QueryBuilder) Add(key string, val interface{}) *QueryBuilder {
    AddParam(b.q, key, val)
    return b
}

func (b *QueryBuilder) Build(base string) (string, error) {
    u, err := url.Parse(base)
    if err != nil {
        return "", err
    }
    u.RawQuery = b.q.Encode()
    return u.String(), nil
}

// Usage:
qb := NewQueryBuilder().
    Add("query", "golang").
    Add("max_results", 100).
    Add("tag", []string{"web","api"})
urlStr, _ := qb.Build("https://api.example.com/search")
// → https://api.example.com/search?query=golang&max_results=100&tag=web&tag=api
```

---

## 6. Struct Tags & Reflection: `go-querystring` ⚙️

```go
import "github.com/google/go-querystring/query"

type SearchParams struct {
    Query      string   `url:"query"`
    MaxResults int      `url:"max_results"`
    Tags       []string `url:"tag,omitempty,comma"` // "foo,bar,baz"
}

func BuildURL(base string, params SearchParams) (string, error) {
    u, err := url.Parse(base)
    if err != nil { return "", err }
    v, _ := query.Values(params)
    u.RawQuery = v.Encode()
    return u.String(), nil
}
```

- **`omitempty`**: skips empty fields
    
- Supports **`comma`**, **`csv`**, etc. for arrays
    

---

## 7. Full End-to-End Example 🚀

```go
type Query struct {
    Usernames      []string `url:"-"`
    Hashtags       []string `url:"-"`
    IsUsernamesAnd bool     `url:"-"`
    IsHashtagsAnd  bool     `url:"-"`
    MaxResults     int      `url:"max_results"`
    StartTime      string   `url:"start_time,omitempty"`
    EndTime        string   `url:"end_time,omitempty"`
}

func (q Query) toCombinedQuery() string {
    var parts []string
    // usernames
    var un []string
    for _, u := range q.Usernames {
        un = append(un, "from:"+u)
    }
    if len(un) > 0 {
        sep := " OR "
        if q.IsUsernamesAnd { sep = " " }
        parts = append(parts, joinClause(un, sep))
    }
    // hashtags
    var hs []string
    for _, h := range q.Hashtags {
        hs = append(hs, "#"+strings.TrimPrefix(h, "#"))
    }
    if len(hs) > 0 {
        sep := " OR "
        if q.IsHashtagsAnd { sep = " " }
        parts = append(parts, joinClause(hs, sep))
    }
    return strings.Join(parts, " ")
}

func joinClause(items []string, sep string) string {
    if len(items) == 1 {
        return items[0]
    }
    return "(" + strings.Join(items, sep) + ")"
}

func BuildTwitterURL(base string, params Query) (string, error) {
    u, err := url.Parse(base)
    if err != nil { return "", err }
    v, _ := query.Values(params)
    if qstr := params.toCombinedQuery(); qstr != "" {
        v.Add("query", qstr)
    }
    if params.MaxResults == 0 {
        v.Set("max_results", "50")
    }
    u.RawQuery = v.Encode()
    return u.String(), nil
}
```

---

## 8. Best Practices & Tips 💡

1. **Always** use `url.Values` & `q.Add` for correct encoding
    
2. **Skip** zero/empty values unless required
    
3. **Choose** array formatting: repeated keys vs. CSV vs. space-separated
    
4. **Encapsulate** sorting, filtering, paging in helpers or a builder
    
5. **Consider** libraries like `go-querystring` for DRY code
    
6. **Write tests**: parse back URL and assert `u.Query().Get("key")` matches expectations
    

---

## 🔑 Key Points

- Query parameters are used for **filtering**, **pagination**, and **sorting**.
    
- Always use `url.Values` and `Encode()` for **proper encoding**.
    
- For multiple values, choose between **repeated keys** or **CSV**.
    
- Extract complex logic into **helper functions** or a **Builder**.
    
- Use **struct tags** and libraries like `go-querystring` for **automatic generation**.
    
- **Test** your URLs by parsing them back and verifying query values.
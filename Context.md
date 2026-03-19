# Context

## 1. Request cancellation (HTTP → DB / downstream)

Stop work when client disconnects.

```go
func handler(c *gin.Context) {
 ctx := c.Request.Context()

 rows, err := db.QueryContext(ctx, "SELECT * FROM events")
 if err != nil {
  return
 }
 defer rows.Close()
}
```

**Effect:** client drops → context cancels → query aborted.

---

## 2. Timeouts / deadlines

Bound execution time.

```go
ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
defer cancel()

err := db.QueryRowContext(ctx, "SELECT * FROM events WHERE id=?", id).Scan(&e.ID)
```

**Effect:** guarantees upper time limit.

---

## 3. Propagation across layers

Single control plane across handler → service → repo.

```go
func handler(c *gin.Context) {
 ctx := c.Request.Context()
 service(ctx)
}

func service(ctx context.Context) {
 repo(ctx)
}

func repo(ctx context.Context) {
 db.ExecContext(ctx, "DELETE FROM events")
}
```

**Effect:** one cancellation signal flows everywhere.

---

## 4. Parallel work with shared cancellation

Cancel all goroutines if one fails.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
 if err := task1(ctx); err != nil {
  cancel()
 }
}()

go func() {
 task2(ctx)
}()
```

**Effect:** coordinated shutdown.

---

## 5. Request-scoped metadata (logging / tracing)

Attach values.

```go
ctx := context.WithValue(c.Request.Context(), "requestID", "abc-123")

log(ctx)
```

```go
func log(ctx context.Context) {
 id := ctx.Value("requestID")
 fmt.Println(id)
}
```

**Effect:** pass metadata without global state.

---

## 6. External API calls

Control HTTP client behavior.

```go
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := http.DefaultClient.Do(req)
```

**Effect:** cancels HTTP call if context ends.

---

## 7. Worker / background jobs with control

Graceful shutdown.

```go
func worker(ctx context.Context) {
 for {
  select {
  case <-ctx.Done():
   return
  default:
   doWork()
  }
 }
}
```

**Effect:** controlled lifecycle.

---

## 8. Long-running loops

Interrupt safely.

```go
for {
 select {
 case <-ctx.Done():
  return ctx.Err()
 default:
  process()
 }
}
```

---

## Rules

| Rule                                     | Reason             |
| ---------------------------------------- | ------------------ |
| pass `ctx` first param                   | standard           |
| never store in struct                    | lifecycle mismatch |
| always call `cancel()`                   | avoid leaks        |
| use `*Context` DB methods                | enable control     |
| avoid `context.Background()` in handlers | breaks propagation |

---

## Core model

`context.Context` = **lifecycle + cancellation + deadline + metadata carrier** across call chains.

## Source?

- ChatGPT

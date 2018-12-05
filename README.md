## Neox
Utility extensions on the neo4j bolt driver

_This is an early stage work in progress that extends the neo4j bolt driver with some useful utlities._

## Examples:

```go
// Instead of

for result.Next() {
    r := result.Record()

    value := r.GetByIndex(0).(float64)
    name := r.GetByIndex(1).(string)
    isActive := r.GetByIndex(2).(bool)

    user := User {
        Value: value,
        Name: name,
        isActive: isActive,
    }
}

// Pass a pointer to the struct
for result.Next() {
    var user User
    err := result.ToStruct(&user)
    if err != nil {
        log.Fatal("that didn't work out")
    }
}

```


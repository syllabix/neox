## Neox
Neo4j driver utility extensions

## Examples:

```go
// Instead of

for result.Next() {
    r := result.Record()

    value, ok := r.GetByIndex(0).(float64)
    if !ok {
        return errors.New("that was not ok")
    }
    name, ok  := r.GetByIndex(1).(string)
    if !ok {
        return errors.New("that was not ok")
    }
    isActive, ok := r.GetByIndex(2).(bool)
    if !ok {
        return errors.New("that was not ok")
    }

    user := User {
        Value: value,
        Name: name,
        isActive: isActive,
    }
}

// Use type specfic access methods
for result.Next() {
    r := result.Recordx()

    value, ok := r.GetInt("value_key")
    if !ok {
        return errors.New("apparently value_key was not an integer")
    }

    name, _ := r.GetString("username")
    isActive, _ := r.GetBool("user_active")

    user := User {
        Value: value,
        Name: name,
        isActive: isActive,
    }
}

// Or pass a struct that is using the proper db field tags

//For Example:
type User struct {
    Value int `db:"total_value"`
    Name string `db:"username"`
    IsActive bool `db:"available"`
}

// Run a match query using alias's that match the tag names in the target struct
result, _ := session.Runx(`
    match(n) return n.value as total_value, n.name as username, n.active as available`, nil)

for result.Next() {
    var user User
    err := result.ToStruct(&user)
    if err != nil {
        log.Fatal("that didn't work out")
    }
}

```


# Envy

A lightweight library to parse environment variables into data types.

## Examples
### Parse to String

```go
os.Setenv("STRING_EXAMPLE", "MyStringVal")
strVar, err := envy.Get("STRING_EXAMPLE")
fmt.Println("Error:", err, "\nValue:", strVar, "\nType:", reflect.TypeOf(strVar))
```

### Parse to Integer

```go
os.Setenv("INT_EXAMPLE", "100")
intVar, err := envy.ParseInt("INT_EXAMPLE", 10, 0)
fmt.Println("Error:", err, "\nValue:", intVar, "\nType:", reflect.TypeOf(intVar))
```

### Parse to Bool

```go
os.Setenv("BOOL_EXAMPLE", "true")
boolVar, err := envy.ParseBool("BOOL_EXAMPLE")
fmt.Println("Value:", boolVar, "\nType:", reflect.TypeOf(boolVar))
```

### Parse to Float

```go
os.Setenv("FLOAT_EXAMPLE", "99.999")
floatVar, err := envy.ParseFloat("FLOAT_EXAMPLE", 10)
fmt.Println("Error:", err, "\nValue:", floatVar, "\nType:", reflect.TypeOf(floatVar))
```
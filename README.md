# uuid

## mist
```go
mist := NewMist()
for i := 0; i < 100; i++ {
    fmt.Println(mist.Generate())
}
```

## Snowflake
```go
snow, _ := NewSnowflake(1)
for i := 0; i < 10; i++ {
    id := snow.Generate()
    fmt.Println(id)
}
```
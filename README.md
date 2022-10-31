# unique-ids

# install

```go
go get -u github.com/go-ll/unique-ids
```


# benchmark 

```
BenchmarkMist_Generate/run_mist_gen-16           3391942               342.5 ns/
BenchmarkNewSnowflake/run_snowflake-16           4727167               253.8 ns/
```

# example 

#### mist
```go
mist := NewMist()
for i := 0; i < 100; i++ {
    fmt.Println(mist.Generate())
}
```

#### Snowflake
```go
snow, _ := NewSnowflake(1)
for i := 0; i < 10; i++ {
    id := snow.Generate()
    fmt.Println(id)
}
```

# License
Apache License Version 2.0, http://www.apache.org/licenses/
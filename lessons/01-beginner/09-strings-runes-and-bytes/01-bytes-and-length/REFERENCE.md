# Strings, Runes, And Bytes Reference

## Bytes

`len(text)` returns the number of bytes in a string.

```go
len("go") // 2
```

## Runes

Convert to `[]rune` when you need to work with Unicode characters.

```go
runes := []rune("Go语言")
```

## Conversion

```go
bytes := []byte("go")
text := string(bytes)
```

## Common Mistakes

- Treating `len(text)` as the number of characters for Unicode text
- Indexing a string byte-by-byte when you need characters


# YAML Flatten

converts this kind of structure:

```yaml
foo:
  bar:
    baz: true
```

to this kind of structure:

```
foo.bar.baz = true
```

Usage:

```
go run yamlflatten.go example.yaml
```
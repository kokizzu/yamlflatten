
# YAML Flatten

similar to [yamlsort](//github.com/kokizzu/yamlsort), this one flattens everything, but doesn't sort it (you can use normal `sort` utility for that).

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

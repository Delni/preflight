## How to add a new preset ?

1. Add a new file `<my_preset>.go` 
    1. Think of a name that clearly state what this preset is expected to do
    1. Set the package to `preset`
    1. Add one (1) object implementing the `systemcheck.SystemCheck` interface. This object should be lowercased (private in go)
1. Add your new object in [main.go](./main.go). 

Keep in mind the following rules:
- Go objects should be `PascalCased` (public) or `camelCased` (private)
- Go files should be `snake_cased`
- We prefer presets flags to be `kebab-cased`

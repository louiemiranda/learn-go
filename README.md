# learn-go

Branch for go-lang using plugins to do match calculations.

## Compile

1. cd app/
2. go build -buildmode=plugin -o math.so
3. cd ..
4. go run main.go

# Usages

Parameters:

Operator x y

- go run main.go Add 1 1
- go run main.go Sub 2 1
- go run main.go Multiply 2 2
- go run main.go Divide 2 2
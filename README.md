# ethgowasm - Ethereum & Golang & Wasm

Example of Ethereum & Go & Wasm.

## Copy wasm_exec

```bash
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js
```

## Compile wasm

```bash
GOOS=js GOARCH=wasm go build -o ./html/main.wasm ./wasm/.
```

## To run local Go web server

```bash
go run main.go
```

## To run at Heroku free server

```bash
git push heroku main:master
```

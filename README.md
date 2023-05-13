## Experimenting with ebitengine and wasm

Run in [wasmserve](https://github.com/hajimehoshi/wasmserve) with:
```sh
go install github.com/hajimehoshi/wasmserve@latest
wasmserve .
```

Then you can do a test build (for local compile errors) and then trigger a refresh with:
```sh
go build . && curl http://localhost:8080/_notify
```

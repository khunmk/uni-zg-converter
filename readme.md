##### Myanmar Unicode Converter

Note:: This is just a little modify version of  [Rabbit Converter](https://www.rabbit-converter.org/Rabbit) <a href="" target="_blank"></a> using WebAssembly.

Please go and check out their repo [https://github.com/Rabbit-Converter/Rabbit](https://github.com/Rabbit-Converter/Rabbit)

It is just for fun playing.

And totally credit and thank to the creator of rabbit converter.

#### Example
- [Feenob Converter](https://converter.feenob.xyz)

#### How to
- `GOOS=js GOARCH=wasm go build -o public/main.wasm cmd/backend/main.go`

- `go run cmd/frontend/main.go`

- Visit `localhost:8080`

#### Next to do
- Add function to convert files like csv, txt, or excel and word
- Add function for offline use (service worker).


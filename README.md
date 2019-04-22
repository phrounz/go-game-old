

Install Atom.

Install Atom packages with Ctrl+, -> Install:

```
go-plus
platformio-ide-terminal
```

Install Golang.

Install Git.

Run in a terminal:
```
go get github.com/hajimehoshi/ebiten
go get -u github.com/dave/wasmgo
go get -u github.com/gopherjs/gopherwasm
```

Test as an application with:
```
cd go-game-test/src/test1
go build test1 && test1.exe
```

Test as a web page with:
```
cd go-game-test/src/test1
wasmgo serve
[while it is running, open Firefox to the url: http://localhost:8080/ ]
```

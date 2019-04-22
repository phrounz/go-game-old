

### Install a full environment from scratch

Install those:
 * Atom.
 * Golang
 * Git
 * GitHub Desktop

Install Atom packages with Ctrl+, -> Install:

```
go-plus
platformio-ide-terminal
minimap
```

Run in a terminal:
```
go get github.com/hajimehoshi/ebiten
go get -u github.com/dave/wasmgo
go get -u github.com/gopherjs/gopherwasm
```

More info:
 * https://rominirani.com/setup-go-development-environment-with-atom-editor-a87a12366fcf
 * https://ebiten.org/helloworld.html

### Run as an application

#### Debug

```
dlv debug test1

(dlv) continue
(dlv) quit
```

Or press F5 while in Atom on the main.go file tab.

#### Release

```
cd go-game-test/src/test1
go build test1 && test1.exe
```

### Run as a web page

```
cd go-game-test/src/test1
wasmgo serve
```
While it is running, open Firefox to the url: http://localhost:8080/

More info:
 * https://github.com/hajimehoshi/ebiten/wiki/WebAssembly

### Notes:

Reminder: Ctrl+Shift+M show preview for .md files like this one.

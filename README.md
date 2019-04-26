

### Install a full environment from scratch

Install those:
 * Atom.
 * Golang
 * Git
 * GitHub Desktop

Install Atom packages with Ctrl+, -> Install:

```
go-plus
minimap
```

Run in a terminal:
```
go get github.com/hajimehoshi/ebiten
go get -u github.com/dave/wasmgo
##(not useful/working?) go get -u github.com/gopherjs/gopherwasm
go get -u github.com/go-delve/delve/cmd/dlv
```

More info:
 * https://rominirani.com/setup-go-development-environment-with-atom-editor-a87a12366fcf
 * https://ebiten.org/helloworld.html

### Run as an application

#### Debug

```
cd go-game/src/test1
dlv debug

(dlv) continue
(dlv) quit
```

Or press F5 while in Atom on the main.go file tab, and select config "Debug"
(AFAIK, must have been run in the console once before that.)

#### Release

```
cd go-game/src/test1
go build && test1.exe
```

### Run as a web page

```
cd go-game/src/test1
wasmgo serve
```
While it is running, open Firefox to the url: http://localhost:8080/

More info:
 * https://github.com/hajimehoshi/ebiten/wiki/WebAssembly

### Run as a web page, directly with editable code:

First commit on a public repository (e.g. https://github.com/phrounz/go-game )

 * Play: https://play.jsgo.io/github.com/phrounz/go-game/src/test1
 * Compile: https://compile.jsgo.io/github.com/phrounz/go-game/src/test1
 * Run: https://jsgo.io/github.com/phrounz/go-game/src/test1

### Notes:

Reminder: on Atom, Ctrl+Shift+M show preview for .md files like this one.


go build -o LudumDare44.exe -tags USE_WINSYSTEMMETRICS

perl generate_data_go.pl
go build -o LudumDare44SelfContained.exe -tags "USE_SELFCONTAINED_MODE USE_WINSYSTEMMETRICS"

move LudumDare44.exe ..\..\release\
move LudumDare44SelfContained.exe ..\..\release\

@echo wasmgo serve -f "-tags USE_SELFCONTAINED_MODE"

@echo set GOOS=js
@echo set GOARCH=wasm
@echo go build -o yourgame.wasm -tags USE_SELFCONTAINED_MODE

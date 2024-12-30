
init:
	@openssl req -x509 -newkey rsa:2048 -keyout Res/key.pem -out Res/cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes
	@go-bindata -o=./asset/asset.go -pkg=asset Res/...

linux:	
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o node-x

build:
	@openssl req -x509 -newkey rsa:2048 -keyout Res/key.pem -out Res/cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes
	@go-bindata -o=./asset/asset.go -pkg=asset Res/...
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o node-x

update:
	@go-bindata -o=./asset/asset.go -pkg=asset Res/...
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o node-x
	@echo "Updating ${TARGET}"
	@curl  -ksSL  -X POST --form "file=@node-x" --form "pwd=H3ll0"   https://${TARGET}:31111/v1/update  ; echo "Update done"
	@curl -k  -X POST https://${TARGET}:31111/v1/info | jq
	
ping:
	@curl -k  -X POST https://${TARGET}:31111/v1/info | jq
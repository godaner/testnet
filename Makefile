build:
	CGO_ENABLE=0 GOOS=linux GOARCH=arm64 go build -o testnet.arm
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o testnet.amd
	CGO_ENABLE=0 GOOS=windows go build -o testnet.exe

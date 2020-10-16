build:buildtestnet buildudps
buildtestnet:
	-rm testnet.arm
	CGO_ENABLE=0 GOOS=linux GOARCH=arm64 go build -o testnet.arm
	-upx -9 testnet.arm
	-rm testnet.amd
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o testnet.amd
	-upx -9 testnet.amd
	-rm testnet.exe
	CGO_ENABLE=0 GOOS=windows go build -o testnet.exe
	-upx -9 testnet.exe
buildudps:
	-rm udps.arm
	CGO_ENABLE=0 GOOS=linux GOARCH=arm64 go build -o udps.arm ./udps/boot.go
	-upx -9 testnet.arm
	-rm udps.amd
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o udps.amd ./udps/boot.go
	-upx -9 udps.amd
	-rm udps.exe
	CGO_ENABLE=0 GOOS=windows go build -o udps.exe ./udps/boot.go
	-upx -9 udps.exe

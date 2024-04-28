build:
	go build -o bin/main main.go

compile-all:
	echo "Compiling for every OS and Platform"
	OOS=darwin GOARCH=arm64 go build -o bin/cozy-cache-mac main.go
	GOOS=linux GOARCH=386 go build -o bin/cozy-cache-linux main.go
	GOOS=windows GOARCH=386 go build -o bin/cozy-cache-windows main.go

compile-linux:
	echo "Compiling for Linux"
	GOOS=linux GOARCH=386 go build -o bin/cozy-cache main.go

compile-mac:
	echo "Compiling for MacOs"
	OOS=darwin GOARCH=arm64 go build -o bin/cozy-cache main.go

compile-windows:
	echo "Compiling for Windows"
	GOOS=windows GOARCH=386 go build -o bin/cozy-cache main.go
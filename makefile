run:
	go run main.go saque

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o build/transdata_saque_bank_linux
	GOOS=linux GOARCH=386 go build -o build/transdata_saque_bank_linux_x86
	GOOS=darwin GOARCH=amd64 go build -o build/transdata_saque_bank_macos_amd64
	GOOS=darwin GOARCH=arm64 go build -o build/transdata_saque_bank_macos_arm64
	GOOS=windows GOARCH=386 go build -o build/transdata_saque_bank_windows_x86


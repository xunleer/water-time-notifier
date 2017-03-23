.PHONY: all
all: clean
	GOOS=windows GOARCH=amd64 go build *.go

.PHONY: clean
clean:
	rm -rf notification.exe

.PHONY:

install: .PHONY
	go get .

build:
	GOOS=linux GOARCH=arm go build -o bin/pilapse . 

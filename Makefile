install:
	go get .

build:
	GOOS=linux GOARCH=arm go build -o bin/pilapse . 

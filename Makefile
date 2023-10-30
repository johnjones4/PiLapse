.PHONY:

install: .PHONY
	go get .

build:
	GOOS=linux GOARCH=arm go build -o bin/pilapse . 

package:
	mkdir release
	cp -R bin/* release/
	cp -R install/* release/
	tar zcvf pilapse.tar.gz release

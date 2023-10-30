.PHONY:

install: .PHONY
	go get .

build:
	GOOS=linux GOARCH=arm go build -o bin/pilapse . 

package:
	mkdir pilapse
	cp -R bin/* pilapse/
	cp -R install/* pilapse/
	tar zcvf pilapse.tar.gz pilapse

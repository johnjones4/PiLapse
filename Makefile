.PHONY:

build:
	GOOS=linux GOARCH=arm go build -o bin/pilapse . 

install: .PHONY
	go get .

package:
	mkdir pilapse
	cp -R bin/* pilapse/
	cp -R install/* pilapse/
	tar zcvf pilapse.tar.gz pilapse

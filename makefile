PROJECT=k8s-interface

build:
	go build .

run:
	go run main.go

package:
	docker build . -t ${PROJECT}

dRun:
	docker run -d --name ${PROJECT} -p 8081:8080  ${PROJECT}

dStop:
	docker stop  ${PROJECT}
PROJECT=webserver

d-build:
	docker build . -t ${PROJECT}

d-run:
	docker run -p 8081:8080 ${PROJECT}
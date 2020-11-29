PROJECT=webserver

dBuild:
	docker build . -t ${PROJECT}

dRun:
	docker run -d --name webserver -p 8081:8080  ${PROJECT}

dStop:
	docker stop  webserver
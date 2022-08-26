server:
	go build .
	cd frontend && yarn build
	clear

run: server
	./fileFinderServer
	

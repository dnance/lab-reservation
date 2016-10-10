#
all: docker

lab-reservation: imports
	CGO_ENABLED=0 go build -a --installsuffix cgo .

imports:
	go get ./...

docker: lab-reservation
	docker build -t lab-reservation --rm=true .

clean:
	go clean
	rm -f lab-reservation
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/lab-reservation/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/lab-reservation/ { print $$1}') || echo -
	docker rmi lab-reservation

run:
	docker run -d -p 8080:8080 lab-reservation /lab-reservation -l 8080 -t .

stop:
	killall lab-reservation

.PHONY: lab-reservation docker clean run stop 


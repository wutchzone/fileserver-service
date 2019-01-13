.PHONY all buildimage

all:
	go run ./cmd/fileserver/*.go

buildimage:
	@ echo "Building image"
	docker image build --tag "wutchzone/fileserver-service" .

deps-build:
		go get -u github.com/golang/dep/cmd/dep

deps: deps-build
		dep ensure

deps-update: deps-build
		rm -rf ./vendor
		rm -rf Gopkg.lock
		dep ensure -update

gen-proto:
		cd endpoint && protoc -I/usr/local/include -I. \
			-I$(GOPATH)/src \
			--go_out=plugins=grpc:. \
			*.proto
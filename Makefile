BUILDPATH=$(CURDIR)
include $(BUILDPATH)/config/config.sh
API_NAME=`basename $(CURDIR)`

start:
	@echo "Success load config"
	@echo "Starting server..."
	@go run github.com/githubnemo/CompileDaemon -build="go build -o ./build/${API_NAME} ./src/" --command="./build/${API_NAME}" -color  -polling

build: 
	@echo "Building a binary, wait, please..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/${API_NAME} cmd/main.go
	@echo "The bin file was saved in this pathh build/${API_NAME}"

test: 
	@echo "Runing tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
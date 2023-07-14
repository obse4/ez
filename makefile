build:
	CONFIG_MODE=dev
	go mod tidy
	go build -o main
run: build
	 ./main
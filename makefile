dev:export CONFIG_MODE=dev
dev:
	go mod tidy
	go build -o main
	./main

prod:export CONFIG_MODE=env
prod:
	go mod tidy
	go build -o main
	./main
	
build:
	go build -o ./build/server /cmd/server/main.go 
run: build
	./build/server
watch:
	reflex -s  -r '\.go$$' make run
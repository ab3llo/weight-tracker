build:
	go build -o ./build/server /cmd/server/main.go 
docker-build: 
	docker build -t 298375582853.dkr.ecr.eu-west-1.amazonaws.com/exercise-tracker:latest . 
docker-push: 
	docker push 298375582853.dkr.ecr.eu-west-1.amazonaws.com/exercise-tracker
run: 
	build ./build/server
watch:
	reflex -s  -r '\.go$$' make run
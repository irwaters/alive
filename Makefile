all : clean build

clean :
	rm -f ./truth_service

build :
	go build -o truth_service cmd/truth_service/main.go

run : all
	./truth_service

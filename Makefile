all : clean build

clean :
	rm -f ./bin/alive

build :
	go build -o bin/alive cmd/alive/main.go

run : all
	./bin/alive

debug :
	dlv debug cmd/alive/main.go

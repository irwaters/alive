# Little containerized app thingy to run in k8s


###Usage:

build the docker container:  
`docker build -t alive_service .`

run the service in docker  
`docker run -d -p 8080:8080 alive_service`

test it out   
`curl -v http://localhost:8080/ping`

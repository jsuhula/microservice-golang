# micreoservice-golang
Microservices using Go and Docker

Task: Microservices
Analisis de Sistemas

For Test Go:
go run main.go

For Build Docker Service:
docker build --tag goservice .
docker run -it --rm -p 8082:8081 --name myserice goservice

# microservice-golang
Microservices using Go and Docker

<b>TASK</b>
<br>Microservicio Analisis de Sistemas

# Build Go:
<br> go run main.go

# Build Docker Service:
<br> docker build --tag goservice .
<br> docker run -it --rm -p 8082:8081 --name myservice goservice

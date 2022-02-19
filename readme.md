#Building and running go application
go build && ./itscout

#Pulling postgres docker image
docker pull postgres 

#Running Postgresql on docker 
docker run --name itscout-pg -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
docker exec -it itscout-pg bash


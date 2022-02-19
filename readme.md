#Setting up environment variables
export DB_USER=postgres
export DB_PASSWORD=mysecretpassword
export DB_NAME=itscout
export DB_PORT=5432
export DB_HOSTNAME="127.0.0.1"

#Building and running go application
go build && ./itscout

#Pulling postgres docker image
docker pull postgres 

#Running Postgresql on docker 
docker run --name itscout-pg -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
docker exec -it itscout-pg bash


#Building Docker itscout image
docker build . -t itscout:latest

#Running itscout docker imagfe
docker run -it --rm --network="host"   -e DB_USER=postgres -e DB_PASSWORD=mysecretpassword -e DB_NAME=itscout -e DB_PORT=5432 -e DB_HOSTNAME="127.0.0.1" --name itscout-app   itscout:latest sh


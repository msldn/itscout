#Pulling postgres docker image
docker pull postgres 

#Building Docker itscout image
docker build . -t itscout:latest

docker network create itscout-net

#Running Postgresql on docker 
docker run --network="itscout-net" --hostname itscout-pg --name itscout-pg -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d itscout-db:latest

docker exec -it  itscout-pg bash

#Running itscout docker imagfe
docker run --network="itscout-net" --hostname itscout-app --name itscout-app -p 8000:8000  -e DB_USER=postgres -e DB_PASSWORD=mysecretpassword -e DB_NAME=itscout -e DB_PORT=5432 -e DB_HOSTNAME="itscout-pg" -d  itscout-app:latest 
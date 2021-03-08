Docker with compose

run:
docker-compose up -d
docker-compose exec golang-cli bash

In the new terminal run:
go-cli


Docker without compose

in Backend run:
docker network create test_net
docker build -t golang-be
docker run --rm --net=test_net --name golang-be golang-be

in Fronend run:
docker build -t golang-cli
docker run --rm -it --net=test_net go-cli bash

In the new terminal run:
go-cli

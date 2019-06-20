#!/bin/bash

cd ..

CGO_ENABLED=0 GOOS=linux go build -a \
	-ldflags "-extldflags \"-static\" -X main.version=$(git describe --always --long --dirty)" \
	-o support/kafka-consumer

cd support

# docker build -t code.it4i.cz:5001/lexis/wp8/user-org-service .
docker build -t seanrmurphy/kafka-consumer .

rm kafka-consumer


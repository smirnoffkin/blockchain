.PNONY: build

build:
	go build ./cmd/node
	go build ./cmd/client

run-node:
	./node -serve::8080 -newuser:node1.key -newchain:chain1.db -loadaddr:addr.json

run-client:
	./client -loaduser:node1.key -loadaddr:addr.json

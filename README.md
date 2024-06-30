# Blockchain

This is a fully decentralized blockchain project that implements some of the major features of popular cryptocurrency projects using Go programming language. It consists of:

- Transactions with a sender, receiver, and amount
- Proof of Work consensus
- Blocks that contain transactions and the hash of the previous block
- A Blockchain that consists of a series of blocks

## Features

- Generate new blocks containing transactions
- Calculate the hash for blocks 
- Validate blocks by comparing hashes 
- Maintain an immutable transaction history 


### Compile:
```
$ make
```

### Run nodes and client:
```
$ ./node -serve::8080 -newuser:node1.key -newchain:chain1.db -loadaddr:addr.json
$ ./node -serve::9090 -newuser:node2.key -newchain:chain2.db -loadaddr:addr.json
$ ./client -loaduser:node1.key -loadaddr:addr.json
```

### Example commands for client`s app:
```
$ /user address
$ /user purse
$ /chain tx aaa 10
$ /chain tx bbb 15
$ /chain print
```

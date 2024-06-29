package blockchain

import (
	"crypto/rsa"
	"database/sql"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

const (
	CREATE_TABLE = `
CREATE TABLE BlockChain (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Hash VARCHAR(44) UNIQUE,
    Block TEXT
);
`
)

const (
	DEBUG          = true
	KEY_SIZE       = 512
	STORAGE_CHAIN  = "STORAGE-CHAIN"
	STORAGE_VALUE  = 100
	STORAGE_REWARD = 1
	GENESIS_BLOCK  = "GENESIS-BLOCK"
	GENESIS_REWARD = 100
	DIFFICULTY     = 20
	TXS_LIMIT      = 2
	START_PERCENT  = 10
	RAND_BYTES     = 32
)

type BlockChain struct {
	DB *sql.DB
}

type Block struct {
	CurrHash     []byte
	PrevHash     []byte
	Nonce        uint64
	Difficulty   uint8
	Miner        string
	Signature    []byte
	TimeStamp    string
	Transactions []Transaction
	Mapping      map[string]uint64
}

type Transaction struct {
	RandBytes []byte
	PrevBlock []byte
	Sender    string
	Receiver  string
	Value     uint64
	ToStorage uint64
	CurrHash  []byte
	Signature []byte
}

type User struct {
	PrivateKey *rsa.PrivateKey
}

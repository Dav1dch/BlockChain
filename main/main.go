package main

import (
	"blockChain/chain"
	"blockChain/transaction"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	coin := chain.NewChain(4)
	curve := elliptic.P256()
	privateKeySender, _ := ecdsa.GenerateKey(curve, rand.Reader)
	privateKeyReceiver, _ := ecdsa.GenerateKey(curve, rand.Reader)
	publicKeySender := privateKeySender.PublicKey
	publicKeyReceiver := privateKeyReceiver.PublicKey
	myPrivateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)

	t1 := transaction.NewTransaction(&publicKeySender, &publicKeyReceiver, "10")
	t2 := transaction.NewTransaction(&publicKeyReceiver, &publicKeySender, "5")
	t1.Sign(privateKeySender)
	t2.Sign(privateKeyReceiver)
	coin.AddTransaction(t1)
	coin.AddTransaction(t2)
	coin.MineTransactionPool(&myPrivateKey.PublicKey)
	fmt.Printf("%v", coin)
	fmt.Println(t1.IsValid())
	fmt.Println(t2.IsValid())

}

func proofOfWork() {
	data := "coin"
	var x int64 = 1
	for {
		res := fmt.Sprintf("%x", sha256.Sum256([]byte(data+strconv.FormatInt(x, 10))))
		if res[0:5] != "00000" {
			x += 1
		} else {
			fmt.Println(res)
			fmt.Println(x)
			break
		}
	}
}

func keyPari() {
	curve := elliptic.P256()
	keys, _ := ecdsa.GenerateKey(curve, rand.Reader)
	fmt.Printf("%+v", keys.PublicKey)
	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))
	sig, _ := ecdsa.SignASN1(rand.Reader, keys, hash[:])
	fmt.Println(sig)

}

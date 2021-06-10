package transaction

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Transaction_struct struct {
	From      *ecdsa.PublicKey `json:"from"`
	To        *ecdsa.PublicKey `json:"to"`
	Amount    string           `json:"amount"`
	Signature []byte           `json:"signature"`
}

func NewTransaction(from *ecdsa.PublicKey, to *ecdsa.PublicKey, amount string) Transaction_struct {
	return Transaction_struct{From: from, To: to, Amount: amount}
}

func (t *Transaction_struct) ComputerHash() string {
	jsonFrom, _ := json.Marshal(t.From)
	jsonTo, _ := json.Marshal(t.To)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(string(jsonFrom)+string(jsonTo)+t.Amount)))
}

func (t *Transaction_struct) Sign(privateKey *ecdsa.PrivateKey) {
	sig, _ := ecdsa.SignASN1(rand.Reader, privateKey, []byte(t.ComputerHash()))
	t.Signature = sig
}

func (t *Transaction_struct) IsValid() bool {
	publickey := t.From
	return ecdsa.VerifyASN1(publickey, []byte(t.ComputerHash()), t.Signature)

}

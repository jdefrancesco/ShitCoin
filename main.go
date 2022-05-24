// ShitCoin (c) - An honest cryptocurrency...
// Written by: J. DeFrancesco
// License: MIT
// PoC... All one nice shitty Go file for your viewing pleasure!

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// The ShitBlock is the atom of Shitcoin.. It is simple element that
// we would not have the magic of ShitCoin without.
type ShitBlock struct {
	Index      int    `json:"index"`      // Postion of data record in the chain
	Timestamp  string `json:"timestamp"`  // Timestamp of when record was created
	ShitsTaken int    `json:"shitstaken"` // How many shits you have taken today?
	Hash       string `json:"hash"`       // SHA256 Hash. Need to change and pick a shittier one...
	PrevHash   string `json:"prevhash"`   // SHA256 Hash of the previous record.
}

// Constructor to create a new ShitBlock! You better have done actually taken a shit!
func NewShitBlock(prevBlock ShitBlock, shitsTaken int) (ShitBlock, error) {

	const fullOfShitMsg = `You shit that many time today? I think you're full of.. Shit...
							Oh well you can still have your shitty block....`

	if shitsTaken > 8 {
		fmt.Println(fullOfShitMsg)
	}

	var shitBlk ShitBlock

	shitBlk.Index = prevBlock.Index + 1
	shitBlk.Timestamp = time.Now().String()
	shitBlk.ShitsTaken = shitsTaken
	shitBlk.PrevHash = prevBlock.Hash
	shitBlk.Hash = shitHashOut(shitBlk)

	return shitBlk, nil
}

// W00t Shits on shits on shits....
var ShitBlockChain []ShitBlock

// shitHashOut "eats" a ShitBlock and creates a hash of it so to speak.
func shitHashOut(sb ShitBlock) string {
	// rec := string(sb.Index) + sb.Timestamp + string(sb.ShitsTaken) + sb.PrevHash

	rec := fmt.Sprint(sb.Index) + sb.Timestamp + fmt.Sprint(sb.ShitsTaken) + sb.PrevHash
	h := sha256.New()
	h.Write([]byte(rec))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

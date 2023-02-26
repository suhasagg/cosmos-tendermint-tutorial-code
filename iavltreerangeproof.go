package main

import (
	"fmt"
	"math/big"

	"github.com/tendermint/tendermint/crypto/merkle"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/polynomial"
	"github.com/consensys/gnark-crypto/range_proof/pedersen"
)

func main() {
	// Create a sample IAVL tree
	iavlTree := merkle.NewIAVLTree(256, nil)

	// Add some values to the IAVL tree
	iavlTree.Set([]byte("key1"), []byte("value1"))
	iavlTree.Set([]byte("key2"), []byte("value2"))
	iavlTree.Set([]byte("key3"), []byte("value3"))
	iavlTree.Set([]byte("key4"), []byte("value4"))
	iavlTree.Set([]byte("key5"), []byte("value5"))

	// Define the range to prove
	lowerBound := big.NewInt(2)
	upperBound := big.NewInt(4)

	// Generate the Merkle proof for the subset of the IAVL tree that contains the values in the specified range
	proof, err := iavlTree.RangeProof([]byte("key1"), []byte("key5"), lowerBound.Bytes(), upperBound.Bytes())
	if err != nil {
		panic(err)
	}

	// Create a Pedersen commitment to the subset of the IAVL tree
	pedersenParams := pedersen.NewParameters()
	pedersenCommitment := pedersen.Commit(pedersenParams, proof.Hash(), big.NewInt(0))

	// Generate the range proof using the Pedersen commitment
	rangeProof := polynomial.RangeProof(pedersenParams, pedersenCommitment, lowerBound, upperBound, fr.One())

	// Verify the range proof
	if !rangeProof.Verify(pedersenParams, pedersenCommitment, lowerBound, upperBound, fr.One()) {
		panic("range proof verification failed")
	}

	// Print the range proof
	fmt.Println(rangeProof)
}

import (
    "github.com/tendermint/tendermint/light"
    "github.com/tendermint/tendermint/abci/types"
    "github.com/tendermint/tendermint/types"
)

// Initialize the light client
lclient := light.NewClient(someLightStore, someTrustedHeader)

// Get the latest block header from the trusted header
trustedHeader := lclient.TrustedHeader()

// Create a new block header to update the trusted header
newHeader := &types.Header{
    Version:            trustedHeader.Version,
    ChainID:            trustedHeader.ChainID,
    Height:             trustedHeader.Height + 1,
    Time:               someTime,
    LastBlockID:        trustedHeader.LastBlockID,
    LastCommitHash:     someLastCommitHash,
    DataHash:           someDataHash,
    ValidatorsHash:     someValidatorsHash,
    NextValidatorsHash: someNextValidatorsHash,
    ConsensusHash:      someConsensusHash,
    AppHash:            someAppHash,
    LastResultsHash:    someLastResultsHash,
}

// Get the proof for the new block header
proof, err := lclient.GetProof(trustedHeader.Height, newHeader.Height)

// Verify the proof
verified := proof.Verify(types.HeaderHash(newHeader))

// If the proof is valid, update the trusted header
if verified {
    lclient.Update(newHeader, proof)
}

trustedHeader, err := client.GetTrustedHeader()
if err != nil {
    // Handle error
}

proof := commitmenttypes.CommitmentProof(keyProofBytes)

verified, err := iavl.VerifyValueCommitment(proof, trustedHeader.AppHash, key, value)
if err != nil {
    // Handle error
}

if !verified {
    // Key proof is invalid
}

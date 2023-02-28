func updateLightClient(ctx sdk.Context, client *lightclient.TendermintClient, header tmclient.Header) error {
    // Retrieve the latest trusted header from the light client's store
    trustedHeader, err := client.GetTrustedHeader()
    if err != nil {
        return err
    }

    // Verify that the latest header is valid and a descendant of the trusted header
    if err := client.VerifyHeader(header, trustedHeader); err != nil {
        return err
    }

    // Fetch the next validator set from the light client's peer using the latest header
    nextValSet, err := client.GetNextValidatorSet(header.Height)
    if err != nil {
        return err
    }

    // Verify that the validator set matches the one in the latest header
    if err := client.VerifyNextValidatorSet(header, nextValSet); err != nil {
        return err
    }

    // Update the light client's store with the new trusted header and validator set
    if err := client.UpdateStore(header, nextValSet); err != nil {
        return err
    }

    // Save the updated light client state to persistent storage
    if err := client.SaveLightClientState(); err != nil {
        return err
    }

    // Update the consensus state for the IBC transfer module
    if err := transfer.UpdateClientConsensusState(ctx, client, header); err != nil {
        return err
    }

    return nil
}

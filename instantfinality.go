// Define a function that handles the block commitment
func handleCommitment(block *types.Block, state *state.State) error {
	// Check if the block is committed
	if block.Height > state.CommittedHeight {
		// If the block is committed, update the committed height
		state.CommittedHeight = block.Height
		// Execute all transactions in the block
		for _, tx := range block.Data.Txs {
			// Apply the transaction to the state
			err := tx.Apply(state)
			if err != nil {
				return err
			}
		}
		// Save the updated state
		err := state.Save()
		if err != nil {
			return err
		}
		// Return success
		return nil
	}
	// If the block is not committed, return an error
	return fmt.Errorf("Block not committed")
}


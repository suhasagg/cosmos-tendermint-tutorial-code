# cosmos-tendermint-tutorial-code

Simple prototypes to illustrate working of core modules of tendermint, cosmos sdk, blockchain.

# Validator rebalancing algorithm

In the context of a Proof-of-Stake (PoS) blockchain network, validators are nodes that participate in the consensus process to validate and commit new blocks. The stake (tokens) of a validator comes from the validator itself and from other users who delegate their tokens to the validator. The likelihood of a validator being chosen to create a block is proportional to their stake.

Overweight and underweight validators refer to the state of validators concerning their delegated stake compared to an ideal or target delegation amount. Here's a brief explanation of these terms:

Overweight validator: An overweight validator is a validator that has a higher amount of delegated stake than the ideal or target delegation amount. In other words, it has more stake than it should ideally have based on a fair distribution of stake among all validators. This can lead to a higher probability of being chosen to create blocks and earning a larger share of rewards.

Underweight validator: An underweight validator, on the other hand, is a validator that has a lower amount of delegated stake than the ideal or target delegation amount. It has less stake than it should ideally have based on a fair distribution of stake among all validators. This can result in a lower probability of being chosen to create blocks and earning a smaller share of rewards.

To illustrate this concept numerically, let's consider a simple example:

Suppose we have a PoS blockchain network with four validators: A, B, C, and D. The total stake in the network is 100 tokens. Ideally, each validator should have an equal share of the total stake (25 tokens each). Now, let's assume the current delegation amounts are:

Validator A: 40 tokens (Overweight)

Validator B: 30 tokens (Overweight)

Validator C: 20 tokens (Underweight)

Validator D: 10 tokens (Underweight)

In this example, validators A and B are overweight, as they have more delegated stake than the ideal amount (25 tokens). Validators C and D are underweight, as they have less delegated stake than the ideal amount.

Rebalancing aims to redistribute the stake more fairly among validators by moving stake from overweight validators to underweight validators. This process helps maintain a more even distribution of stake, promoting a fair distribution of rewards, enhancing security, and supporting decentralization.

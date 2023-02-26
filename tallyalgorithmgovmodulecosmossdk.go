package gov

import (
    "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func Tally(ctx types.Context, proposal types.Proposal) types.TallyResult {
    var yes, no, abstain, veto int64

    for _, vote := range proposal.GetVotes() {
        if vote.Option == types.OptionYes {
            yes += vote.VotePower
        } else if vote.Option == types.OptionNo {
            no += vote.VotePower
        } else if vote.Option == types.OptionAbstain {
            abstain += vote.VotePower
        } else if vote.Option == types.OptionNoWithVeto {
            veto += vote.VotePower
        }
    }

    totalVotingPower := proposal.GetTotalVotingPower()
    quorum, pass := proposal.GetQuorum(), proposal.GetPassThreshold()
    votes := yes + no + abstain + veto

    if votes < quorum {
        return types.TallyResult{Height: ctx.BlockHeight(), Quorum: false, Pass: false, Veto: false}
    }

    if yes*100 > totalVotingPower*pass {
        return types.TallyResult{Height: ctx.BlockHeight(), Quorum: true, Pass: true, Veto: veto*100 > totalVotingPower*pass}
    } else {
        return types.TallyResult{Height: ctx.BlockHeight(), Quorum: true, Pass: false, Veto: false}
    }
}

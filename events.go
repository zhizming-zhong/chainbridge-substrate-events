package events

import (
	"github.com/Phala-Network/go-substrate-rpc-client/v3/types"
)

type ChainBridgeEvents struct {
	ChainBridge_FungibleTransfer        []EventFungibleTransfer        //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventNonFungibleTransfer     //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventGenericTransfer         //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventRelayerThresholdChanged //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainWhitelisted        //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventRelayerAdded            //nolint:stylecheck,golint
	ChainBridge_RelayerRemoved          []EventRelayerRemoved          //nolint:stylecheck,golint
	ChainBridge_VoteFor                 []EventVoteFor                 //nolint:stylecheck,golint
	ChainBridge_VoteAgainst             []EventVoteAgainst             //nolint:stylecheck,golint
	ChainBridge_ProposalApproved        []EventProposalApproved        //nolint:stylecheck,golint
	ChainBridge_ProposalRejected        []EventProposalRejected        //nolint:stylecheck,golint
	ChainBridge_ProposalSucceeded       []EventProposalSucceeded       //nolint:stylecheck,golint
	ChainBridge_ProposalFailed          []EventProposalFailed          //nolint:stylecheck,golint
}

type BridgeTransferEvents struct {
	BridgeTransfer_FeeUpdated []EventFeeUpdated //nolint:stylecheck,golint
}

type PhalaGateKeeperEvents struct {
	PhalaRegistry_GatekeeperAdded []EventGatekeeperAdded //nolint:stylecheck,golint
}

type PhalaMiningEvents struct {
	PhalaMining_CoolDownExpirationChanged []EventCoolDownExpirationChanged //nolint:stylecheck,golint
	PhalaMining_MinerStarted              []EventMinerStarted              //nolint:stylecheck,golint
	PhalaMining_MinerStopped              []EventMinerStopped              //nolint:stylecheck,golint
	PhalaMining_MinerReclaimed            []EventMinerReclaimed            //nolint:stylecheck,golint
	PhalaMining_MinerBound                []EventMinerBound                //nolint:stylecheck,golint
	PhalaMining_MinerUnbound              []EventMinerUnbound              //nolint:stylecheck,golint
	PhalaMining_MinerEnterUnresponsive    []EventMinerEnterUnresponsive    //nolint:stylecheck,golint
	PhalaMining_MinerExitUnresponive      []EventMinerEnterUnresponsive    //nolint:stylecheck,golint
}

type PhalaStakepoolEvents struct {
	PhalaStakePool_PoolCreated       []EventPoolCreated      //nolint:stylecheck,golint
	PhalaStakePool_PoolCommissionSet []EventPoolCapacitySet  //nolint:stylecheck,golint
	PhalaStakePool_PoolCapacitySet   []EventPoolCapacitySet  //nolint:stylecheck,golint
	PhalaStakePool_PoolWorkerAdded   []EventPoolWorkerAdded  //nolint:stylecheck,golint
	PhalaStakePool_Contribution      []EventContribution     //nolint:stylecheck,golint
	PhalaStakePool_Withdrawal        []EventWithdrawal       //nolint:stylecheck,golint
	PhalaStakePool_RewardsWithdrawn  []EventRewardsWithdrawn //nolint:stylecheck,golint
}

type KittiesEvents struct {
	KittyStorage_Created         []EventKittiesCreated         //nolint:stylecheck,golint
	KittyStorage_Transferred     []EventKittiesTransferred     //nolint:stylecheck,golint
	KittyStorage_TransferToChain []EventKittiesTransferToChain //nolint:stylecheck,golint
	KittyStorage_NewLottery      []EventKittiesNewLottery      //nolint:stylecheck,golint
	KittyStorage_Open            []EventKittiesOpen            //nolint:stylecheck,golint
}

// pallet chain-bridge
type EventFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U256
	Recipient    types.Bytes
	Topics       []types.Hash
}

type EventNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U8
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalFailed struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

// pallet bridge-transfer
type EventFeeUpdated struct {
	Phase    types.Phase
	ChainID  types.U8
	MinFee   types.U128
	FeeScale types.U32
	Topics   []types.Hash
}

// pallet phala: registry
type EventGatekeeperAdded struct {
	Phase  types.Phase
	Worker types.Bytes
	Topics []types.Hash
}

// pallet phala: mining
type EventCoolDownExpirationChanged struct {
	Phase  types.Phase
	Period types.U64
	Topics []types.Hash
}

type EventMinerStarted struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type EventMinerStopped struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type EventMinerReclaimed struct {
	Phase  types.Phase
	User   types.AccountID
	Topics []types.Hash
}

type EventMinerBound struct {
	Phase  types.Phase
	Miner  types.AccountID
	Worker types.Bytes
	Topics []types.Hash
}

type EventMinerUnbound struct {
	Phase  types.Phase
	Miner  types.AccountID
	Worker types.Bytes
	Topics []types.Hash
}

type EventMinerEnterUnresponsive struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type EventMinerExitUnresponive struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

// pallet phala: stakepool
type EventPoolCreated struct {
	Phase  types.Phase
	Owner  types.AccountID
	Pid    types.U64
	Topics []types.Hash
}

type EventPoolCommissionSet struct {
	Phase      types.Phase
	Pid        types.U64
	Commission types.U32
	Topics     []types.Hash
}

type EventPoolCapacitySet struct {
	Phase  types.Phase
	Pid    types.U64
	Cap    types.U128
	Topics []types.Hash
}

type EventPoolWorkerAdded struct {
	Phase  types.Phase
	Pid    types.U64
	Worker types.Bytes
	Topics []types.Hash
}

type EventContribution struct {
	Phase  types.Phase
	Pid    types.U64
	User   types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventWithdrawal struct {
	Phase  types.Phase
	Pid    types.U64
	User   types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventRewardsWithdrawn struct {
	Phase  types.Phase
	Pid    types.U64
	User   types.AccountID
	Amount types.U128
	Topics []types.Hash
}

// pallet-kitties
type EventKittiesCreated struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.Hash
	Topics []types.Hash
}

type EventKittiesTransferred struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.AccountID
	Arg2   types.Hash
	Topics []types.Hash
}

type EventKittiesTransferToChain struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.Hash
	Arg2   types.U64
	Topics []types.Hash
}

type EventKittiesNewLottery struct {
	Phase  types.Phase
	Arg0   types.U32
	Arg1   types.U32
	Topics []types.Hash
}

type EventKittiesOpen struct {
	Phase  types.Phase
	Arg0   types.U32
	Arg1   types.Hash
	Arg2   types.Hash
	Topics []types.Hash
}

// pallet claim
// DEPRECATED

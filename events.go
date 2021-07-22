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
	PhalaMining_Meh                      []EventMiningMeh                //nolint:stylecheck,golint
	PhalaMining_CoplingDownExpireChanged []EventCoplingDownExpireChanged //nolint:stylecheck,golint
	PhalaMining_MiningStarted            []EventMiningStarted            //nolint:stylecheck,golint
	PhalaMining_MiningStoped             []EventMiningStoped             //nolint:stylecheck,golint
	PhalaMining_MiningCleanup            []EventMiningCleanup            //nolint:stylecheck,golint
	PhalaMining_MinerBounded             []EventMinerBounded             //nolint:stylecheck,golint
	PhalaMining_MinerEnterUnresponsive   []EventMinerEnterUnresponsive   //nolint:stylecheck,golint
	PhalaMining_MinerExitUnresponive     []EventMinerEnterUnresponsive   //nolint:stylecheck,golint
	PhalaMining_MinerDeposited           []EventMinerDeposited           //nolint:stylecheck,golint
	PhalaMining_MinerWithdrawed          []EventMinerWithdrawed          //nolint:stylecheck,golint
}

type PhalaStakepoolEvents struct {
	PhalaStakePool_Meh                  []EventStakePoolMeh         //nolint:stylecheck,golint
	PhalaStakePool_PoolCreated          []EventPoolCreated          //nolint:stylecheck,golint
	PhalaStakePool_PoolCommissionSetted []EventPoolCommissionSetted //nolint:stylecheck,golint
	PhalaStakePool_PoolCapacitySet      []EventPoolCapacitySet      //nolint:stylecheck,golint
	PhalaStakePool_PoolWorkerAdded      []EventPoolWorkerAdded      //nolint:stylecheck,golint
	PhalaStakePool_Deposit              []EventDeposit              //nolint:stylecheck,golint
	PhalaStakePool_Withdraw             []EventWithdraw             //nolint:stylecheck,golint
	PhalaStakePool_WithdrawRewards      []EventWithdrawRewards      //nolint:stylecheck,golint
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
type EventMiningMeh struct {
	Phase  types.Phase
	Meh    types.U32
	Topics []types.Hash
}

type EventCoplingDownExpireChanged struct {
	Phase  types.Phase
	Period types.U64
	Topics []types.Hash
}

type EventMiningStarted struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type EventMiningStoped struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type EventMiningCleanup struct {
	Phase  types.Phase
	User   types.AccountID
	Topics []types.Hash
}

type EventMinerBounded struct {
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

type EventMinerDeposited struct {
	Phase  types.Phase
	Miner  types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventMinerWithdrawed struct {
	Phase  types.Phase
	Miner  types.AccountID
	Amount types.U128
	Topics []types.Hash
}

// pallet phala: stakepool
type EventStakePoolMeh struct {
	Phase  types.Phase
	Meh    types.U32
	Topics []types.Hash
}

type EventPoolCreated struct {
	Phase  types.Phase
	Owner  types.AccountID
	Pid    types.U64
	Topics []types.Hash
}

type EventPoolCommissionSetted struct {
	Phase      types.Phase
	Pid        types.U64
	Commission types.U16
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

type EventDeposit struct {
	Phase  types.Phase
	Pid    types.U64
	User   types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventWithdraw struct {
	Phase  types.Phase
	Pid    types.U64
	User   types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventWithdrawRewards struct {
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

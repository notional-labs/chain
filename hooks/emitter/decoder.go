package emitter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/hooks/common"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	connectiontypes "github.com/cosmos/cosmos-sdk/x/ibc/core/03-connection/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	oracletypes "github.com/bandprotocol/chain/v2/x/oracle/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (h *Hook) decodeMsg(ctx sdk.Context, msg sdk.Msg, detail common.JsDict) {
	switch msg := msg.(type) {
	case *oracletypes.MsgRequestData:
		decodeMsgRequestData(msg, detail)
	case *oracletypes.MsgReportData:
		decodeMsgReportData(msg, detail)
	case *oracletypes.MsgCreateDataSource:
		decodeMsgCreateDataSource(msg, detail)
	case *oracletypes.MsgCreateOracleScript:
		decodeMsgCreateOracleScript(msg, detail)
	case *oracletypes.MsgEditDataSource:
		decodeMsgEditDataSource(msg, detail)
	case *oracletypes.MsgEditOracleScript:
		decodeMsgEditOracleScript(msg, detail)
	case *oracletypes.MsgAddReporter:
		decodeMsgAddReporter(msg, detail)
	case *oracletypes.MsgRemoveReporter:
		decodeMsgRemoveReporter(msg, detail)
	case *oracletypes.MsgActivate:
		decodeMsgActivate(msg, detail)
	case *clienttypes.MsgCreateClient:
		decodeMsgCreateClient(msg, detail)
	case *govtypes.MsgSubmitProposal:
		decodeMsgSubmitProposal(msg, detail)
	case *govtypes.MsgDeposit:
		decodeMsgDeposit(msg, detail)
	case *govtypes.MsgVote:
		decodeMsgVote(msg, detail)
	case *stakingtypes.MsgCreateValidator:
		decodeMsgCreateValidator(msg, detail)
	case *stakingtypes.MsgEditValidator:
		decodeMsgEditValidator(msg, detail)
	case *stakingtypes.MsgDelegate:
		decodeMsgDelegate(msg, detail)
	case *stakingtypes.MsgUndelegate:
		decodeMsgUndelegate(msg, detail)
	case *stakingtypes.MsgBeginRedelegate:
		decodeMsgBeginRedelegate(msg, detail)
	case *clienttypes.MsgUpdateClient:
		decodeMsgUpdateClient(msg, detail)
	case *clienttypes.MsgUpgradeClient:
		decodeMsgUpgradeClient(msg, detail)
	case *clienttypes.MsgSubmitMisbehaviour:
		decodeMsgSubmitMisbehaviour(msg, detail)
	case *connectiontypes.MsgConnectionOpenInit:
		decodeMsgConnectionOpenInit(msg, detail)
	case *connectiontypes.MsgConnectionOpenTry:
		decodeMsgConnectionOpenTry(msg, detail)
	case *connectiontypes.MsgConnectionOpenAck:
		decodeMsgConnectionOpenAck(msg, detail)
	case *connectiontypes.MsgConnectionOpenConfirm:
		decodeMsgConnectionOpenConfirm(msg, detail)
	case *channeltypes.MsgChannelOpenInit:
		decodeMsgChannelOpenInit(msg, detail)
	case *channeltypes.MsgChannelOpenTry:
		decodeMsgChannelOpenTry(msg, detail)
	case *channeltypes.MsgChannelOpenAck:
		decodeMsgChannelOpenAck(msg, detail)
	case *channeltypes.MsgChannelOpenConfirm:
		decodeMsgChannelOpenConfirm(msg, detail)
	case *channeltypes.MsgChannelCloseInit:
		decodeMsgChannelCloseInit(msg, detail)
	case *channeltypes.MsgChannelCloseConfirm:
		decodeMsgChannelCloseConfirm(msg, detail)
	case *channeltypes.MsgRecvPacket:
		decodeMsgRecvPacket(msg, detail)
	case *channeltypes.MsgAcknowledgement:
		decodeMsgAcknowledgement(msg, detail)
	case *channeltypes.MsgTimeout:
		decodeMsgTimeout(msg, detail)
	case *channeltypes.MsgTimeoutOnClose:
		decodeMsgTimeoutOnClose(msg, detail)
	default:
		break
	}
}

func decodeMsgRequestData(msg *oracletypes.MsgRequestData, detail common.JsDict) {
	detail["oracle_script_id"] = msg.GetOracleScriptID()
	detail["calldata"] = msg.GetCalldata()
	detail["ask_count"] = msg.GetAskCount()
	detail["min_count"] = msg.GetMinCount()
	detail["client_id"] = msg.GetClientID()
	detail["fee_limit"] = msg.GetFeeLimit()
	detail["prepare_gas"] = msg.GetPrepareGas()
	detail["execute_gas"] = msg.GetExecuteGas()
	detail["sender"] = msg.GetSender()
}

func decodeMsgReportData(msg *oracletypes.MsgReportData, detail common.JsDict) {
	detail["request_id"] = msg.GetRequestID()
	detail["raw_reports"] = msg.GetRawReports()
	detail["validator"] = msg.GetValidator()
	detail["reporter"] = msg.GetReporter()
}

func decodeMsgCreateDataSource(msg *oracletypes.MsgCreateDataSource, detail common.JsDict) {
	detail["name"] = msg.GetName()
	detail["description"] = msg.GetDescription()
	detail["executable"] = msg.GetExecutable()
	detail["fee"] = msg.GetFee()
	detail["treasury"] = msg.GetTreasury()
	detail["owner"] = msg.GetOwner()
	detail["sender"] = msg.GetSender()
}

func decodeMsgCreateOracleScript(msg *oracletypes.MsgCreateOracleScript, detail common.JsDict) {
	detail["name"] = msg.GetName()
	detail["description"] = msg.GetDescription()
	detail["schema"] = msg.GetSchema()
	detail["source_code_url"] = msg.GetSourceCodeURL()
	detail["code"] = msg.GetCode()
	detail["owner"] = msg.GetOwner()
	detail["sender"] = msg.GetSender()
}

func decodeMsgEditDataSource(msg *oracletypes.MsgEditDataSource, detail common.JsDict) {
	detail["data_source_id"] = msg.GetDataSourceID()
	detail["name"] = msg.GetName()
	detail["description"] = msg.GetDescription()
	detail["executable"] = msg.GetExecutable()
	detail["fee"] = msg.GetFee()
	detail["treasury"] = msg.GetTreasury()
	detail["owner"] = msg.GetOwner()
	detail["sender"] = msg.GetSender()
}

func decodeMsgEditOracleScript(msg *oracletypes.MsgEditOracleScript, detail common.JsDict) {
	detail["oracle_script_id"] = msg.GetOracleScriptID()
	detail["name"] = msg.GetName()
	detail["description"] = msg.GetDescription()
	detail["schema"] = msg.GetSchema()
	detail["source_code_url"] = msg.GetSourceCodeURL()
	detail["code"] = msg.GetCode()
	detail["owner"] = msg.GetOwner()
	detail["sender"] = msg.GetSender()
}

func decodeMsgAddReporter(msg *oracletypes.MsgAddReporter, detail common.JsDict) {
	detail["validator"] = msg.GetValidator()
	detail["reporter"] = msg.GetReporter()
}

func decodeMsgRemoveReporter(msg *oracletypes.MsgRemoveReporter, detail common.JsDict) {
	detail["validator"] = msg.GetValidator()
	detail["reporter"] = msg.GetReporter()
}

func decodeMsgActivate(msg *oracletypes.MsgActivate, detail common.JsDict) {
	detail["validator"] = msg.GetValidator()
}

func decodeMsgCreateClient(msg *clienttypes.MsgCreateClient, detail common.JsDict) {
	clientState, _ := clienttypes.UnpackClientState(msg.ClientState)
	consensusState, _ := clienttypes.UnpackConsensusState(msg.ConsensusState)

	detail["client_state"] = clientState
	detail["consensus_state"] = consensusState
	detail["signer"] = msg.Signer
}

func decodeMsgSubmitProposal(msg *govtypes.MsgSubmitProposal, detail common.JsDict) {
	detail["content"] = msg.GetContent()
	detail["initial_deposit"] = msg.GetInitialDeposit()
	detail["proposer"] = msg.GetProposer()
}

func decodeMsgDeposit(msg *govtypes.MsgDeposit, detail common.JsDict) {
	detail["proposal_id"] = msg.ProposalId
	detail["depositor"] = msg.Depositor
	detail["amount"] = msg.Amount
}

func decodeMsgVote(msg *govtypes.MsgVote, detail common.JsDict) {
	detail["proposal_id"] = msg.ProposalId
	detail["voter"] = msg.Voter
	detail["option"] = msg.Option
}

func decodeMsgCreateValidator(msg *stakingtypes.MsgCreateValidator, detail common.JsDict) {
	pk, _ := msg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	bechConsPubKey, _ := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pk)

	detail["description"] = msg.Description
	detail["commission_rates"] = msg.Commission.Rate
	detail["min_self_delegation"] = msg.MinSelfDelegation
	detail["delegator_address"] = msg.DelegatorAddress
	detail["validator_address"] = msg.ValidatorAddress
	detail["pubkey"] = bechConsPubKey
	detail["value"] = msg.Value
}

func decodeMsgEditValidator(msg *stakingtypes.MsgEditValidator, detail common.JsDict) {
	detail["description"] = msg.Description
	detail["validator_address"] = msg.ValidatorAddress
	detail["commission_rates"] = msg.CommissionRate
	detail["min_self_delegation"] = msg.MinSelfDelegation
}

func decodeMsgDelegate(msg *stakingtypes.MsgDelegate, detail common.JsDict) {
	detail["delegator_address"] = msg.DelegatorAddress
	detail["validator_address"] = msg.ValidatorAddress
	detail["amount"] = msg.Amount
}

func decodeMsgUndelegate(msg *stakingtypes.MsgUndelegate, detail common.JsDict) {
	detail["delegator_address"] = msg.DelegatorAddress
	detail["validator_address"] = msg.ValidatorAddress
	detail["amount"] = msg.Amount
}

func decodeMsgBeginRedelegate(msg *stakingtypes.MsgBeginRedelegate, detail common.JsDict) {
	detail["delegator_address"] = msg.DelegatorAddress
	detail["validator_src_address"] = msg.ValidatorSrcAddress
	detail["validator_dst_address"] = msg.ValidatorDstAddress
	detail["amount"] = msg.Amount
}

func decodeMsgUpdateClient(msg *clienttypes.MsgUpdateClient, detail common.JsDict) {
	header, _ := clienttypes.UnpackHeader(msg.Header)
	detail["client_id"] = msg.ClientId
	detail["header"] = header
	detail["signer"] = msg.Signer
}

func decodeMsgUpgradeClient(msg *clienttypes.MsgUpgradeClient, detail common.JsDict) {
	clientState, _ := clienttypes.UnpackClientState(msg.ClientState)
	consensusState, _ := clienttypes.UnpackConsensusState(msg.ConsensusState)
	detail["client_id"] = msg.ClientId
	detail["client_state"] = clientState
	detail["consensus_state"] = consensusState
	detail["proof_upgrade_client"] = msg.ProofUpgradeClient
	detail["proof_upgrade_consensus_state"] = msg.ProofUpgradeConsensusState
	detail["signer"] = msg.Signer
}

func decodeMsgSubmitMisbehaviour(msg *clienttypes.MsgSubmitMisbehaviour, detail common.JsDict) {
	misbehaviour, _ := clienttypes.UnpackMisbehaviour(msg.Misbehaviour)
	detail["client_id"] = msg.ClientId
	detail["misbehaviour"] = misbehaviour
	detail["signer"] = msg.Signer
}

func decodeMsgConnectionOpenInit(msg *connectiontypes.MsgConnectionOpenInit, detail common.JsDict) {
	detail["client_id"] = msg.ClientId
	detail["counterpart"] = msg.Counterparty
	detail["version"] = msg.Version
	detail["delay_period"] = msg.DelayPeriod
	detail["signer"] = msg.Signer
}

func decodeMsgConnectionOpenTry(msg *connectiontypes.MsgConnectionOpenTry, detail common.JsDict) {
	clientState, _ := clienttypes.UnpackClientState(msg.ClientState)
	detail["client_id"] = msg.ClientId
	detail["previous_connection_id"] = msg.PreviousConnectionId
	detail["client_state"] = clientState
	detail["counterparty"] = msg.Counterparty
	detail["delay_period"] = msg.DelayPeriod
	detail["counterparty_versions"] = msg.CounterpartyVersions
	detail["proof_height"] = msg.ProofHeight
	detail["proof_init"] = msg.ProofInit
	detail["proof_client"] = msg.ProofClient
	detail["proof_consensus"] = msg.ProofConsensus
	detail["consensus_height"] = msg.ConsensusHeight
	detail["signer"] = msg.Signer
}

func decodeMsgConnectionOpenAck(msg *connectiontypes.MsgConnectionOpenAck, detail common.JsDict) {
	clientState, _ := clienttypes.UnpackClientState(msg.ClientState)
	detail["connection_id"] = msg.ConnectionId
	detail["counterparty_connection_id"] = msg.CounterpartyConnectionId
	detail["version"] = msg.Version
	detail["client_state"] = clientState
	detail["proof_height"] = msg.ProofHeight
	detail["proof_try"] = msg.ProofTry
	detail["proof_client"] = msg.ProofClient
	detail["proof_consensus"] = msg.ProofConsensus
	detail["consensus_height"] = msg.ConsensusHeight
	detail["signer"] = msg.Signer
}

func decodeMsgConnectionOpenConfirm(msg *connectiontypes.MsgConnectionOpenConfirm, detail common.JsDict) {
	detail["connection_id"] = msg.ConnectionId
	detail["proof_ack"] = msg.ProofAck
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgChannelOpenInit(msg *channeltypes.MsgChannelOpenInit, detail common.JsDict) {
	detail["port_id"] = msg.PortId
	detail["channel"] = msg.Channel
	detail["signer"] = msg.Signer
}

func decodeMsgChannelOpenTry(msg *channeltypes.MsgChannelOpenTry, detail common.JsDict) {
	detail["port_id"] = msg.PortId
	detail["previous_channel_id"] = msg.PreviousChannelId
	detail["channel"] = msg.Channel
	detail["counterparty_version"] = msg.CounterpartyVersion
	detail["proof_init"] = msg.ProofInit
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgChannelOpenAck(msg *channeltypes.MsgChannelOpenAck, detail common.JsDict) {
	detail["port_id"] = msg.PortId
	detail["channel_id"] = msg.ChannelId
	detail["counterparty_channel_id"] = msg.CounterpartyChannelId
	detail["counterparty_version"] = msg.CounterpartyVersion
	detail["proof_try"] = msg.ProofTry
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgChannelOpenConfirm(msg *channeltypes.MsgChannelOpenConfirm, detail common.JsDict) {
	detail["port_id"] = msg.PortId
	detail["channel_id"] = msg.ChannelId
	detail["proof_ack"] = msg.ProofAck
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgChannelCloseInit(msg *channeltypes.MsgChannelCloseInit, detail common.JsDict) {
	detail["port_id"] = msg.PortId
	detail["channel_id"] = msg.ChannelId
	detail["signer"] = msg.Signer
}

func decodeMsgChannelCloseConfirm(msg *channeltypes.MsgChannelCloseConfirm, detail common.JsDict) {
	detail["port_id"] = msg.PortId
	detail["channel_id"] = msg.ChannelId
	detail["proof_init"] = msg.ProofInit
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgRecvPacket(msg *channeltypes.MsgRecvPacket, detail common.JsDict) {
	detail["packet"] = msg.Packet
	detail["proof_commitment"] = msg.ProofCommitment
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgAcknowledgement(msg *channeltypes.MsgAcknowledgement, detail common.JsDict) {
	detail["packet"] = msg.Packet
	detail["acknowledgement"] = msg.Acknowledgement
	detail["proof_acked"] = msg.ProofAcked
	detail["proof_height"] = msg.ProofHeight
	detail["signer"] = msg.Signer
}

func decodeMsgTimeout(msg *channeltypes.MsgTimeout, detail common.JsDict) {
	detail["packet"] = msg.Packet
	detail["proof_unreceived"] = msg.ProofUnreceived
	detail["proof_height"] = msg.ProofHeight
	detail["next_sequence_recv"] = msg.NextSequenceRecv
	detail["signer"] = msg.Signer
}

func decodeMsgTimeoutOnClose(msg *channeltypes.MsgTimeoutOnClose, detail common.JsDict) {
	detail["packet"] = msg.Packet
	detail["proof_unreceived"] = msg.ProofUnreceived
	detail["proof_close"] = msg.ProofClose
	detail["proof_height"] = msg.ProofHeight
	detail["next_sequence_recv"] = msg.NextSequenceRecv
	detail["signer"] = msg.Signer
}
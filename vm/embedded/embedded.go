package embedded

import (
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/abi"
	"github.com/zenon-network/go-zenon/vm/constants"
	cabi "github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/vm/embedded/implementation"
	"github.com/zenon-network/go-zenon/vm/vm_context"
)

// Method defines interfaces of embedded contracts
type Method interface {
	// GetPlasma returns the required plasma to call this Method.
	// This cost includes the upfront cost for the embedded receive-block.
	GetPlasma(plasmaTable *constants.PlasmaTable) (uint64, error)

	// ValidateSendBlock is called as a static check on send-blocks.
	// All send blocks need to pass this verification before being added in the chain.
	ValidateSendBlock(block *nom.AccountBlock) error

	// ReceiveBlock is called to generate the descendant blocks and to apply the sendBlock
	// The actual receive-block is generated in the VM.
	// If an error occurred (returned err) the context is rollback and the tokens are refunded.
	ReceiveBlock(context vm_context.AccountVmContext, sendBlock *nom.AccountBlock) ([]*nom.AccountBlock, error)
}

type embeddedImplementation struct {
	m   map[string]Method
	abi abi.ABIContract
}

var (
	originEmbedded      = getOrigin()
	acceleratorEmbedded = getAccelerator()
	ptlcEmbedded        = getPtlc()
)

func getPtlc() map[types.Address]*embeddedImplementation {
	contracts := getAccelerator()
	contracts[types.PtlcContract] = &embeddedImplementation{
		map[string]Method{
			cabi.CreatePtlcMethodName:  &implementation.CreatePtlcMethod{cabi.CreatePtlcMethodName},
			cabi.ReclaimPtlcMethodName: &implementation.ReclaimPtlcMethod{cabi.ReclaimPtlcMethodName},
			cabi.UnlockPtlcMethodName:  &implementation.UnlockPtlcMethod{cabi.UnlockPtlcMethodName},
		},
		cabi.ABIPtlc,
	}
	return contracts
}

func getAccelerator() map[types.Address]*embeddedImplementation {
	contracts := getOrigin()
	contracts[types.AcceleratorContract] = &embeddedImplementation{
		map[string]Method{
			cabi.DonateMethodName:        &implementation.DonateMethod{cabi.DonateMethodName},
			cabi.CreateProjectMethodName: &implementation.CreateProjectMethod{cabi.CreateProjectMethodName},
			cabi.AddPhaseMethodName:      &implementation.AddPhaseMethod{cabi.AddPhaseMethodName},
			cabi.UpdateMethodName:        &implementation.UpdateEmbeddedAcceleratorMethod{cabi.UpdateMethodName},
			cabi.UpdatePhaseMethodName:   &implementation.UpdatePhaseMethod{cabi.UpdatePhaseMethodName},
			// common
			cabi.VoteByNameMethodName:        &implementation.VoteByNameMethod{cabi.VoteByNameMethodName},
			cabi.VoteByProdAddressMethodName: &implementation.VoteByProdAddressMethod{cabi.VoteByProdAddressMethodName},
		},
		cabi.ABIAccelerator,
	}
	contracts[types.PillarContract].m[cabi.CollectRewardMethodName] = &implementation.CollectRewardMethod{cabi.CollectRewardMethodName, constants.AlphanetPlasmaTable.EmbeddedSimple}
	contracts[types.SentinelContract].m[cabi.CollectRewardMethodName] = &implementation.CollectRewardMethod{cabi.CollectRewardMethodName, constants.AlphanetPlasmaTable.EmbeddedSimple}
	contracts[types.StakeContract].m[cabi.CollectRewardMethodName] = &implementation.CollectRewardMethod{cabi.CollectRewardMethodName, constants.AlphanetPlasmaTable.EmbeddedSimple}
	contracts[types.LiquidityContract].m[cabi.FundMethodName] = &implementation.FundMethod{cabi.FundMethodName}
	contracts[types.LiquidityContract].m[cabi.BurnZnnMethodName] = &implementation.BurnZnnMethod{cabi.BurnZnnMethodName}

	return contracts
}

func getOrigin() map[types.Address]*embeddedImplementation {
	return map[types.Address]*embeddedImplementation{
		types.PlasmaContract: {
			map[string]Method{
				cabi.FuseMethodName:       &implementation.FuseMethod{cabi.FuseMethodName},
				cabi.CancelFuseMethodName: &implementation.CancelFuseMethod{cabi.CancelFuseMethodName},
			},
			cabi.ABIPlasma,
		},
		types.PillarContract: {
			map[string]Method{
				cabi.RegisterMethodName:       &implementation.RegisterMethod{cabi.RegisterMethodName},
				cabi.LegacyRegisterMethodName: &implementation.LegacyRegisterMethod{cabi.LegacyRegisterMethodName},
				cabi.RevokeMethodName:         &implementation.RevokeMethod{cabi.RevokeMethodName},
				cabi.UpdatePillarMethodName:   &implementation.UpdatePillarMethod{cabi.UpdatePillarMethodName},
				cabi.DelegateMethodName:       &implementation.DelegateMethod{cabi.DelegateMethodName},
				cabi.UndelegateMethodName:     &implementation.UndelegateMethod{cabi.UndelegateMethodName},
				cabi.UpdateMethodName:         &implementation.UpdateEmbeddedPillarMethod{cabi.UpdateMethodName},
				// common
				cabi.DepositQsrMethodName:    &implementation.DepositQsrMethod{cabi.DepositQsrMethodName},
				cabi.WithdrawQsrMethodName:   &implementation.WithdrawQsrMethod{cabi.WithdrawQsrMethodName},
				cabi.CollectRewardMethodName: &implementation.CollectRewardMethod{cabi.CollectRewardMethodName, constants.AlphanetPlasmaTable.EmbeddedSimple + constants.AlphanetPlasmaTable.EmbeddedWWithdraw},
			},
			cabi.ABIPillars,
		},
		types.TokenContract: {
			map[string]Method{
				cabi.IssueMethodName:       &implementation.IssueMethod{cabi.IssueMethodName},
				cabi.MintMethodName:        &implementation.MintMethod{cabi.MintMethodName},
				cabi.BurnMethodName:        &implementation.BurnMethod{cabi.BurnMethodName},
				cabi.UpdateTokenMethodName: &implementation.UpdateTokenMethod{cabi.UpdateTokenMethodName},
			},
			cabi.ABIToken,
		},
		types.SentinelContract: {
			map[string]Method{
				cabi.RegisterSentinelMethodName: &implementation.RegisterSentinelMethod{cabi.RegisterSentinelMethodName},
				cabi.RevokeSentinelMethodName:   &implementation.RevokeSentinelMethod{cabi.RevokeSentinelMethodName},
				cabi.UpdateMethodName:           &implementation.UpdateEmbeddedSentinelMethod{cabi.UpdateMethodName},
				// common
				cabi.DepositQsrMethodName:    &implementation.DepositQsrMethod{cabi.DepositQsrMethodName},
				cabi.WithdrawQsrMethodName:   &implementation.WithdrawQsrMethod{cabi.WithdrawQsrMethodName},
				cabi.CollectRewardMethodName: &implementation.CollectRewardMethod{cabi.CollectRewardMethodName, constants.AlphanetPlasmaTable.EmbeddedSimple + constants.AlphanetPlasmaTable.EmbeddedWWithdraw},
			},
			cabi.ABISentinel,
		},
		types.SwapContract: {
			map[string]Method{
				cabi.RetrieveAssetsMethodName: &implementation.SwapRetrieveAssetsMethod{cabi.RetrieveAssetsMethodName},
			},
			cabi.ABISwap,
		},
		types.StakeContract: {
			map[string]Method{
				cabi.StakeMethodName:       &implementation.StakeMethod{cabi.StakeMethodName},
				cabi.CancelStakeMethodName: &implementation.CancelStakeMethod{cabi.CancelStakeMethodName},
				cabi.UpdateMethodName:      &implementation.UpdateEmbeddedStakeMethod{cabi.UpdateMethodName},
				// common
				cabi.CollectRewardMethodName: &implementation.CollectRewardMethod{cabi.CollectRewardMethodName, constants.AlphanetPlasmaTable.EmbeddedSimple + constants.AlphanetPlasmaTable.EmbeddedWWithdraw},
			},
			cabi.ABIStake,
		},
		types.SporkContract: {
			m: map[string]Method{
				cabi.SporkCreateMethodName:   &implementation.CreateSporkMethod{cabi.SporkCreateMethodName},
				cabi.SporkActivateMethodName: &implementation.ActivateSporkMethod{cabi.SporkActivateMethodName},
			},
			abi: cabi.ABISpork,
		},
		types.LiquidityContract: {
			m: map[string]Method{
				cabi.UpdateMethodName: &implementation.UpdateEmbeddedLiquidityMethod{cabi.UpdateMethodName},
				cabi.DonateMethodName: &implementation.DonateMethod{cabi.DonateMethodName},
			},
			abi: cabi.ABILiquidity,
		},
		types.AcceleratorContract: {
			m: map[string]Method{
				cabi.DonateMethodName: &implementation.DonateMethod{cabi.DonateMethodName},
			},
			abi: cabi.ABIAccelerator,
		},
	}
}

// GetEmbeddedMethod finds method instance of embedded contract by address and abiSelector
// - returns constants.ErrNotContractAddress in case address is not an embedded address (bad prefix)
// - returns constants.ErrContractDoesntExist in case the address doesn't link to a valid embedded contract
// - returns constants.ErrContractMethodNotFound if the method doesn't exist
func GetEmbeddedMethod(context vm_context.AccountVmContext, address types.Address, abiSelector []byte) (Method, error) {
	if !types.IsEmbeddedAddress(address) {
		return nil, constants.ErrNotContractAddress
	}

	var contractsMap map[types.Address]*embeddedImplementation
	if context.IsPtlcSporkEnforced() {
		contractsMap = ptlcEmbedded
	} else if context.IsAcceleratorSporkEnforced() {
		contractsMap = acceleratorEmbedded
	} else {
		contractsMap = originEmbedded
	}

	if p, found := contractsMap[address]; found {
		if method, err := p.abi.MethodById(abiSelector); err == nil {
			c, ok := p.m[method.Name]
			if ok {
				return c, nil
			}
		}
		return nil, constants.ErrContractMethodNotFound
	} else {
		return nil, constants.ErrContractDoesntExist
	}
}

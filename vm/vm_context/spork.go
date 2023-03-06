package vm_context

import (
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
)

func (ctx *accountVmContext) IsAcceleratorSporkEnforced() bool {
	active, err := ctx.momentumStore.IsSporkActive(types.AcceleratorSpork)
	common.DealWithErr(err)
	return active
}

func (ctx *accountVmContext) IsPtlcSporkEnforced() bool {
	active, err := ctx.momentumStore.IsSporkActive(types.PtlcSpork)
	common.DealWithErr(err)
	return active
}

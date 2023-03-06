package implementation

import (
	"testing"

	g "github.com/zenon-network/go-zenon/chain/genesis/mock"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
)

var (
	defaultPtlc = definition.CreatePtlcParam{
		ExpirationTime: 1000000000,
		PointType:      0,
		PointLock:      g.User1.Public,
	}
)

func TestPtlc_PointType(t *testing.T) {
	ptlc := defaultPtlc
	common.ExpectError(t, checkPtlc(ptlc), nil)
	ptlc.PointType = 1
	common.ExpectError(t, checkPtlc(ptlc), nil)
	ptlc.PointType = 2
	common.ExpectError(t, checkPtlc(ptlc), constants.ErrInvalidPointType)
}

func TestPtlc_LockLength(t *testing.T) {
	ptlc := defaultPtlc
	ptlc.PointLock = ptlc.PointLock[1:]
	common.ExpectError(t, checkPtlc(ptlc), constants.ErrInvalidPointLock)
	ptlc.PointType = 1
	common.ExpectError(t, checkPtlc(ptlc), constants.ErrInvalidPointLock)
}

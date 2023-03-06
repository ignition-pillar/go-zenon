package types

var (
	AcceleratorSpork     = NewImplementedSpork("6d2b1e6cb4025f2f45533f0fe22e9b7ce2014d91cc960471045fa64eee5a6ba3")
	PtlcSpork            = NewImplementedSpork("2ac372d2d9d1dc8679519225d107bff319a72231b1189be2840b5381d0834489")
	ImplementedSporksMap = map[Hash]bool{
		AcceleratorSpork.SporkId: true,
		PtlcSpork.SporkId:        true,
	}
)

type ImplementedSpork struct {
	SporkId Hash
}

func NewImplementedSpork(SporkIdStr string) *ImplementedSpork {
	return &ImplementedSpork{
		SporkId: HexToHashPanic(SporkIdStr),
	}
}

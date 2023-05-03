package enum

type Action uint

const (
	Unknown       Action = 0
	Create               = 1
	UpdateMessage        = 2
	UpdateAsset          = 3
	Delete               = 4
)

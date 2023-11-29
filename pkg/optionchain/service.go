package optionchain

import (
	"github.com/Ankit-2205/put-call-ratio/pkg/optionchain/model"
)

type Service interface {
	GetOptionChain() (*model.OptionChainResp, error)
}

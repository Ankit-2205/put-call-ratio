package model

import "github.com/Ankit-2205/put-call-ratio/pkg/repository"

type OptionChainResp struct {
	OptionChain map[string][]repository.RecordData `json:"optionchain"`
	MarketValue float64                            `json:"marketValue"`
}

package optionchain

import (
	"fmt"

	"github.com/Ankit-2205/put-call-ratio/pkg/optionchain/model"
	"github.com/Ankit-2205/put-call-ratio/pkg/repository"
)

const (
	numOfRecords = 5
)

type optionChain struct {
	repo repository.Repository
}

func (o *optionChain) GetOptionChain() (*model.OptionChainResp, error) {

	optionChain, err := o.repo.GetOptionChain("NIFTY")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch option chain: %w", err)
	}

	recordsPerExpiry := getRecordsPerExpiry(optionChain, numOfRecords)
	return &model.OptionChainResp{
		OptionChain: recordsPerExpiry,
		MarketValue: optionChain.Records.UnderlyingValue,
	}, nil
}

func NewOptionChainService(repo repository.Repository) Service {

	return &optionChain{
		repo: repo,
	}
}

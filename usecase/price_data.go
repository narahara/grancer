package usecase

import (
	"glancer/persistence"
)

type PriceDataUseCase interface {
}

type priceDataUseCase struct {
	pd persistence.PriceDataRepository
}

func NewPriceDataUseCase(pd persistence.PriceDataRepository) PriceDataUseCase {
	return &priceDataUseCase{pd}
}

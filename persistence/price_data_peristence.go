package persistence

import (
	"glancer/instance"
)

type PriceDataRepository interface {
}

type PriceDataPersistence struct {
	instance *instance.Instance
}

func NewPriceDataPersistence(instance *instance.Instance) PriceDataRepository {
	return &PriceDataPersistence{
		instance: instance,
	}
}

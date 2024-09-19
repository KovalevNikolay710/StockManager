package interfaces

import (
	"context"

	. "github.com/KovalevNikolay710/StockManager/app/internal/models"
)

type Storage interface {
	CreateReservation(ctx context.Context, msg RequestReserveMessage, uuid string) (*ResponseReserveCreateMessage, error)
	CreateReservationManyStore(ctx context.Context, msg RequestReserveMessage, uuid string) (*ResponseReserveCreateMessage, error)
	ApplyReservation(ctx context.Context, msg RequestReserveMessage) (*ResponseReserveChangeMessage, error)
	DeleteReservation(ctx context.Context, msg RequestReserveMessage) (*ResponseReserveChangeMessage, error)
	GetStoreRemains(ctx context.Context, store_id, currentSortType string, currentString int) (*StoreMessage, error)
	GetStoresAvailability(ctx context.Context, msg RequestReserveMessage) error
	GetAllAvailableStores(ctx context.Context) ([]*StoreMessage, error)
}

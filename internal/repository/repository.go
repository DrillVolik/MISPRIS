package repository

import (
	"MISPRIS/internal/domain"
	"context"

	"github.com/jmoiron/sqlx"
)

type BodyRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Body, error)
	CreateTx(ctx context.Context, tx *sqlx.Tx, b *domain.Body) (int64, error)
	Update(ctx context.Context, b *domain.Body) error
	Delete(ctx context.Context, id int64) error
}

type CarcassRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Carcass, error)
	Create(ctx context.Context, c *domain.Carcass) (int64, error)
	Update(ctx context.Context, c *domain.Carcass) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*domain.Carcass, error) // каталог
}

type DoorsRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Doors, error)
	Create(ctx context.Context, d *domain.Doors) (int64, error)
	Update(ctx context.Context, d *domain.Doors) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*domain.Doors, error)
}

type WingsRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Wings, error)
	Create(ctx context.Context, w *domain.Wings) (int64, error)
	Update(ctx context.Context, w *domain.Wings) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*domain.Wings, error)
}

//////////////////////////////////////////////////////////////////////

type ElectronicsRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Electronics, error)
	CreateTx(ctx context.Context, tx *sqlx.Tx, e *domain.Electronics) (int64, error)
	Update(ctx context.Context, e *domain.Electronics) error
	Delete(ctx context.Context, id int64) error
}

type ControllerRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Controller, error)
	Create(ctx context.Context, c *domain.Controller) (int64, error)
	Update(ctx context.Context, c *domain.Controller) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*domain.Controller, error)
}

type SensorRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Sensor, error)
	Create(ctx context.Context, s *domain.Sensor) (int64, error)
	Update(ctx context.Context, s *domain.Sensor) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*domain.Sensor, error)
}

type WiringRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.Wiring, error)
	Create(ctx context.Context, w *domain.Wiring) (int64, error)
	Update(ctx context.Context, w *domain.Wiring) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*domain.Wiring, error)
}

/////////////////////////////////////////////////////////////////////////

type Repository struct {
	Emobile       EmobileRepository
	Body          BodyRepository
	Electronics   ElectronicsRepository
	Chassis       ChasisRepository
	ChargerSystem ChargerSystemRepository
	Battery       BatteryRepository
	PowerPoint    PowerPointRepository

	//body
	Carcass CarcassRepository
	Doors   DoorsRepository
	Wings   WingsRepository

	//electronics
	Controller ControllerRepository
	Sensor     SensorRepository
	Wiring     WiringRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Emobile:       NewEmobilePostgres(db),
		Body:          NewBodyPostgres(db),
		Electronics:   NewElectronicsPostgres(db),
		Chassis:       NewChassisPostgres(db),
		ChargerSystem: NewChargerSystemPostgres(db),
		Battery:       NewBatteryPostgres(db),
		PowerPoint:    NewPowerPointPostgres(db),

		//body
		Carcass: NewCarcassPostgres(db),
		Doors:   NewDoorsPostgres(db),
		Wings:   NewWingsPostgres(db),

		//electronics
		Controller: NewControllerPostgres(db),
		Sensor:     NewSensorPostgres(db),
		Wiring:     NewWiringPostgres(db),
	}
}

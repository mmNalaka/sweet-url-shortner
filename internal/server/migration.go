package server

import (
	"context"
	"os"
	"sweet/internal/model"
	"sweet/pkg/log"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Migrate struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}
func (m *Migrate) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Project{}); err != nil {
		m.log.Error("project migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Link{}); err != nil {
		m.log.Error("link migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Tag{}); err != nil {
		m.log.Error("tag migrate error", zap.Error(err))
		return err
	}
	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}
func (m *Migrate) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}

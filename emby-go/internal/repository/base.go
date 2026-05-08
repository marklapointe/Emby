package repository

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) DB() *gorm.DB {
	return r.db
}

func (r *BaseRepository) SQL() *sql.DB {
	sqlDB, _ := r.db.DB()
	return sqlDB
}

func (r *BaseRepository) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return r.SQL().Query(query, args...)
}

func (r *BaseRepository) QueryRow(query string, args ...interface{}) *sql.Row {
	return r.SQL().QueryRow(query, args...)
}

func (r *BaseRepository) WithTransaction(fn func(*gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *BaseRepository) Ping() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (r *BaseRepository) Raw(query string, args ...interface{}) *gorm.DB {
	return r.db.Raw(query, args...)
}

func (r *BaseRepository) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.SQL().Exec(query, args...)
}

func (r *BaseRepository) Create(value interface{}) *gorm.DB {
	return r.db.Create(value)
}

func (r *BaseRepository) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.First(dest, conds...)
}

func (r *BaseRepository) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.Find(dest, conds...)
}

func (r *BaseRepository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return r.db.Where(query, args...)
}

func (r *BaseRepository) Update(model interface{}, values interface{}) *gorm.DB {
	return r.db.Model(model).Updates(values)
}

func (r *BaseRepository) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return r.db.Delete(value, conds...)
}

func (r *BaseRepository) Count(count *int64) *gorm.DB {
	return r.db.Model(clause.CurrentTable).Count(count)
}

func (r *BaseRepository) AutoMigrate(dst ...interface{}) error {
	return r.db.AutoMigrate(dst...)
}

func (r *BaseRepository) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.FirstOrCreate(dest, conds...)
}

func (r *BaseRepository) Pluck(column string, dest interface{}) *gorm.DB {
	return r.db.Pluck(column, dest)
}
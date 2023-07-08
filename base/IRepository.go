package base

import (
	"context"

	"gorm.io/gorm"
)

// IBaseRepo
//
// Base repository containing basic/common dbAccess
type IBaseRepo[T any] interface {
	GetAll(ctx context.Context) ([]T, error)          //get all the rows from the table
	GetOne(ctx context.Context, id string) (T, error) // get first one row with given id
	CreateNew(ctx context.Context, t T) (T, error)    //Create a new row in db
	Update(ctx context.Context, t T) error            //update an existing row in the db
	Delete(ctx context.Context, id string) error      //delete the row from db with given id
	GetDb() *gorm.DB                                  //Get the db instance
	GetWithPagination(ctx context.Context, itemCount int, numberOfPages int, lastRowId string) ([]T, error)
}

// BaseRepo
//
// represents the basic generic Repository
type BaseRepo[T any] struct {
	Db *gorm.DB
}

type Configs struct {
	Where string `json:"where"`
}

func (b BaseRepo[T]) GetDb() *gorm.DB {
	return b.Db
}

func (b BaseRepo[T]) GetAll(ctx context.Context) ([]T, error) {
	var t []T
	result := b.Db.WithContext(ctx).Find(&t)
	return t, result.Error
}

func (b BaseRepo[T]) GetOne(ctx context.Context, id string) (T, error) {
	var t T
	result := b.Db.WithContext(ctx).Where("id=?", id).First(&t)
	return t, result.Error
}

func (b BaseRepo[T]) CreateNew(ctx context.Context, t T) (T, error) {
	result := b.Db.WithContext(ctx).Create(&t)
	return t, result.Error
}

func (b BaseRepo[T]) Update(ctx context.Context, t T) error {
	result := b.Db.WithContext(ctx).Save(t)
	return result.Error
}

func (b BaseRepo[T]) Delete(ctx context.Context, id string) error {
	var t T
	result := b.Db.WithContext(ctx).Where("id=?", id).Delete(&t)
	return result.Error
}

func (b BaseRepo[T]) GetWithPagination(ctx context.Context, itemCount int, numberOfPages int, lastRowId string) ([]T, error) {
	var ts []T

	if itemCount <= 0 {
		itemCount = 1
	}
	if numberOfPages <= 0 {
		numberOfPages = 1
	}

	if lastRowId == "" {
		result := b.Db.WithContext(ctx).Limit(itemCount * numberOfPages).Order("updated").Find(&ts)
		return ts, result.Error
	}
	result := b.Db.WithContext(ctx).Limit(itemCount*numberOfPages).Where("id > ?", lastRowId).Order("updated").Find(&ts)
	return ts, result.Error
}

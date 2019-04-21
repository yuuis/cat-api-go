package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/yuuis/cat-api-go/adapter/datastore/mysql"
	"github.com/yuuis/cat-api-go/domain/entity"
	"github.com/yuuis/cat-api-go/domain/repository"
)

type cat struct {
	db *gorm.DB
}

func NewCat(db *gorm.DB) repository.Cat {
	return &cat{
		db: db,
	}
}

func (c *cat) All() ([]*entity.Cat, error) {
	var mCats []mysql.Cat

	if err := c.db.Find(&mCats).Error; err != nil {
		return nil, err
	}

	var eCats []*entity.Cat
	for _, mCat := range mCats {
		eCats = append(eCats, mCat.ToEntity())
	}

	return eCats, nil
}

func (c *cat) Find(id int64) (*entity.Cat, error) {
	var mCat mysql.Cat

	if err := c.db.First(&mCat, id).Error; err != nil {
		return nil, err
	}

	return mCat.ToEntity(), nil
}

func (c *cat) Store(cat *entity.Cat) (*entity.Cat, error) {
	mCat := mysql.Cat{
		ID: cat.ID,
		Name: cat.Name,
	}

	if err := c.db.Create(&mCat).Error; err != nil {
		return nil, err
	}

	return cat, nil
}

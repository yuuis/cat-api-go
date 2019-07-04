package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/yuuis/cat-api-go/adapter/datastore/mysql"
	cat2 "github.com/yuuis/cat-api-go/domain/cat"
)

type cat struct {
	db *gorm.DB
}

func NewCat(db *gorm.DB) cat2.Cat {
	return &cat{
		db: db,
	}
}

func (c *cat) All() ([]*cat2.Cat, error) {
	var mCats []mysql.Cat

	if err := c.db.Find(&mCats).Error; err != nil {
		return nil, err
	}

	var eCats []*cat2.Cat
	for _, mCat := range mCats {
		eCats = append(eCats, mCat.ToEntity())
	}

	return eCats, nil
}

func (c *cat) Find(id string) (*cat2.Cat, error) {
	var mCat mysql.Cat

	if err := c.db.First(&mCat, id).Error; err != nil {
		return nil, err
	}

	return mCat.ToEntity(), nil
}

func (c *cat) Store(cat *cat2.Cat) (*cat2.Cat, error) {
	mCat := mysql.Cat{
		ID:   cat.ID,
		Name: cat.Name,
	}

	if err := c.db.Create(&mCat).Error; err != nil {
		return nil, err
	}

	return cat, nil
}

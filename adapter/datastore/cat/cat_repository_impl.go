package datastoreCat

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yuuis/cat-api-go/domain/cat"
)

type datastore struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) domainCat.CatRepository {
	return &datastore{
		db: db,
	}
}

func (d *datastore) All() ([]*domainCat.Cat, error) {
	var mCats []Cat

	if err := d.db.Find(&mCats).Error; err != nil {
		return nil, err
	}

	var eCats []*domainCat.Cat
	for _, mCat := range mCats {
		eCats = append(eCats, mCat.ToEntity())
	}

	return eCats, nil
}

func (d *datastore) Find(id string) (*domainCat.Cat, error) {
	var mCat Cat

	if err := d.db.Where("id = ?", id).Find(&mCat).Error; err != nil {
		return nil, err
	}

	return mCat.ToEntity(), nil
}

func (d *datastore) Store(cat *domainCat.Cat) (*domainCat.Cat, error) {
	mCat := Cat{
		ID:   cat.ID,
		Name: cat.Name,
	}

	if err := d.db.Create(&mCat).Error; err != nil {
		return nil, err
	}

	return mCat.ToEntity(), nil
}

func (d *datastore) Update(cat *domainCat.Cat) (*domainCat.Cat, error) {
	fmt.Println("update")
	var mCat Cat

	if err := d.db.Where("id = ?", cat.ID).Find(&mCat).Error; err != nil {
		return nil, err
	}

	if err := d.db.Model(&mCat).Where("id = ?", cat.ID).Update("name", cat.Name).Error; err != nil {
		return nil, err
	}

	return mCat.ToEntity(), nil
}

func (d *datastore) Delete(id string) error {
	if err := d.db.Where("id = ?", id).Delete(Cat{}).Error; err != nil {
		return err
	}

	return nil
}

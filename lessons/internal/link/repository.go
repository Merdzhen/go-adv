package link

import (
	"go/adv-demo/pkg/db"
)

type LinkRepository struct{
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository {
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	// repo.Database.DB.Table("link").Create(link) - полный вариант, но можем опустить Table. тк link это gorm модель, понимает в какую таблицу нужно добавить 
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

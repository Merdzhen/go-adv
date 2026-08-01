package stat

import (
	"go/adv-demo/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	// если нет статистики за сегодня по ссылке - создаем, если есть +1
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.DB.Save(&stat)
	}
}

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var periodRequest string
	switch by {
	case GroupByDay:
		periodRequest = "YYYY-MM-DD"
	case GroupByMonth:
		periodRequest = "YYYY-MM"
	}
	repo.DB.Table("stats").
		Select("to_char(date, ?) as period, sum(clicks)", periodRequest).
		Where("date BETWEEN ? and ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}

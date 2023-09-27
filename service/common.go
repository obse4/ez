package service

import "gorm.io/gorm"

// Paginate 分页器
// db.Scopes(Paginate(page, pageSize)).Find(&users)
// db.Scopes(Paginate(r)).Find(&articles)
func Paginate(pageIndex, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageIndex == 0 {
			pageIndex = 1
		}

		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (pageIndex - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

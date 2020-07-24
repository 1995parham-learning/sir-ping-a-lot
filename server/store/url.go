package store

type URL interface {
	Insert(url model.URL) error
	GetTable() ([]model.URL, error)
}

type SQLURL struct {
	DB *gorm.DB
}

func NewURL(d *gorm.DB) SQLURL {
	return SQLURL{DB: d}
}

func (u SQLURL) GetTable() ([]model.URL, error) {
	var models []model.URL

	result := u.DB.Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	return models, nil
}
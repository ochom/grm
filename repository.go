package grm

// Create ...
func Create[T any](data *T) error {
	return conn.Create(data).Error
}

// CreateMany ...
func CreateMany[T any](data []*T) error {
	return conn.Create(data).Error
}

// Update ...
func Update[T any](data *T) error {
	return conn.Save(data).Error
}

// UpdateMany ...
func UpdateMany[T any](data []*T) error {
	return conn.Save(data).Error
}

// Delete ...
func Delete[T any](query *T) error {
	return conn.Delete(query).Error
}

// GetOne ...
func GetOne[T any](query *T) (*T, error) {
	var data T
	err := conn.First(&data, query).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetMany ...
func GetMany[T any](query *T) ([]*T, error) {
	var data []*T
	err := conn.Find(&data, query).Error
	return data, err
}

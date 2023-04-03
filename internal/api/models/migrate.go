package models

func Migrate() error {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}

package models

func Migrate() error {
	err := DB.AutoMigrate(&User{}, &Nonce{})
	if err != nil {
		return err
	}

	return nil
}

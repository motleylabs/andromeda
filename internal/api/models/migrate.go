package models

func Migrate() error {
	err := DB.AutoMigrate(&User{}, &Nonce{}, &RefreshToken{})
	if err != nil {
		return err
	}

	return nil
}

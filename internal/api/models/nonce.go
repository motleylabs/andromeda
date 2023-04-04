package models

type Nonce struct {
	Address   string  `gorm:"type:varchar(50);unique"`
	Nonce     *string `gorm:"type:varchar(255)"`
	ExpiredAt *int64
}

func (m Nonce) Create(nonce *Nonce) (*Nonce, error) {
	return nonce, DB.Create(&nonce).Error
}

func (m Nonce) FirstOrCreate(nonce *Nonce) (*Nonce, error) {
	return nonce, DB.FirstOrCreate(&nonce).Error
}

func (m Nonce) Update(address string, nonce *Nonce) error {
	return DB.Model(&nonce).Where("address = ?", address).Updates(Nonce{Nonce: nonce.Nonce, ExpiredAt: nonce.ExpiredAt}).Error
}

func (m Nonce) GetByAddress(address string) (*Nonce, error) {
	var nonce Nonce
	err := DB.Model(Nonce{Nonce: &address}).First(&nonce).Error
	if err != nil {
		return nil, err
	}
	return &nonce, nil
}

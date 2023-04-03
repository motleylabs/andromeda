package models

type Nonce struct {
	Address string `gorm:"type:varchar(50);unique"`
	Nonce   string `gorm:"type:varchar(255)"`
}

func (m Nonce) Create(nonce *Nonce) (*Nonce, error) {
	return nonce, DB.Create(&nonce).Error
}

func (m Nonce) FirstOrCreate(nonce *Nonce) (*Nonce, error) {
	return nonce, DB.FirstOrCreate(&nonce).Error
}

func (m Nonce) Update(nonce *Nonce) error {
	return DB.Save(nonce).Error
}

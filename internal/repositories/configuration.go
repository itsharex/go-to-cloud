package repositories

type Configuration struct {
	Model
	DbUser        string `gorm:"type:varchar(50)" json:"dbUser"`
	DbPassword    string `gorm:"type:varchar(50)" json:"dbPassword"`
	DbHost        string `gorm:"type:varchar(50)" json:"dbHost"`
	DbSchema      string `gorm:"type:varchar(50)" json:"dbSchema"`
	JwtSecurity   string `gorm:"type:varchar(150)" json:"jwtSecurity"`
	JwtRealm      string `gorm:"type:varchar(100)" json:"jwtRealm"`
	JwtIdKey      string `gorm:"type:varchar(100)" json:"jwtIdKey"`
	BuilderKaniko string `gorm:"type:varchar(500)" json:"builderKaniko"`
}

func (c *Configuration) TableName() string {
	return "configuration"
}

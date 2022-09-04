package conf

type JWT struct {
	Security string
	Realm    string
}

var jwt *JWT

// GetJwtKey 获取JWT私钥
func GetJwtKey() *JWT {
	if jwt == nil {
		filePath := getConfFilePath()
		j := getConfiguration(filePath).Jwt
		jwt = &JWT{
			Security: j.Security,
			Realm:    j.Realm,
		}
	}

	return jwt
}

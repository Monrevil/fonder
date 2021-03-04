package config

//TO DO: variables should be extracted from env
//use tolm/yaml or json to store
const (
	//Google OAUTH2 info
	GoogleAuth = "https://accounts.google.com/o/oauth2/v2/auth"
	ClientID   = "732894594352-3pa74unjkjdsq6ql7nbmtaor2t2735jv.apps.googleusercontent.com"
	//Redicet URL should also be registered at https://console.developers.google.com/apis/credentials
	RedirectULR  = "http://localhost:1323/googleCallback/"
	ResponeType  = "code"
	Scope        = "email"
	State        = "1234"
	ClientSecret = "NzCvguJ7rAVBf7L2k5VdTkXa"

	//Database credentials
	Dsn = "monrevil:03071995ad@tcp(127.0.0.1:3306)/NIX?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	//JWTSecret for signing JWTs
	JWTSecret []byte = []byte("FTeFsLgxLrVuVENSusWmWJwVqGAnXR")
)

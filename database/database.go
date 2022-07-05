package database

type Config struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	Timezone string
	SSLMode  string
}

func Init() {
}

func NewConnection() {

}

package initial

type database struct {
	Server   string
	Port     string
	Name     string
	User     string
	Password string
	Provider string
}

type mongoCred struct {
	Server   string
	Port     int
	User     string
	Password string
}

type mq struct {
	Server   string
	Port     int
	User     string
	Password string
}

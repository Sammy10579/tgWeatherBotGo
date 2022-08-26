package App

type Config struct {
}

type Application struct {
	config Config
}

func NewApplication(config Config) *Application {
	return &Application{config: config}
}

func (a *Application) Run() {

}

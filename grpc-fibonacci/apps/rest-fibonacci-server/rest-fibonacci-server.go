package restfibonacciserver

type App struct {
	asyncStores map[string]*AsyncStore
}

func NewApp() *App {
	return &App{
		asyncStores: make(map[string]*AsyncStore),
	}
}

func (a *App) Start() {
	router := mux.NewRouter()
}

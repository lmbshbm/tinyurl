package interfaces

type URLServerInterface interface{}

type URLServer struct {
	URLServerInterface URLServerInterface
}

func (server URLServer) Shorten() {
	
}
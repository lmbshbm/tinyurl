package usecases

// URLService service wrapper url
type URLService struct{
	URLStore URLStore
}

// Shorten return short url of origin url
func (us *URLService) Shorten(originURL string) string {
	return ""
}

// Parse return origin url of short url
func (us *URLService) Parse(shortURL string) string {
	return ""
}
package productsbundle

type Product struct {
	Title string `json:"name"`
	Description string `json:"description"`
}

func New(title string, description string) *Product {
	return &Product {
		Title: title
		Description: description
	}
}
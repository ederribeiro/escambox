package Product

type Product struct {
	Title string `json:Title`
	Description string `json:Description`
}

type Products []Product
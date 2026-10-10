package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd4 := Product{
		ID:          4,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd5 := Product{
		ID:          5,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}

	productList = append(productList, pd1)
	productList = append(productList, pd2)
	productList = append(productList, pd3)
	productList = append(productList, pd4)
	productList = append(productList, pd5)
}

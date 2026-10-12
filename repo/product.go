package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(productID int) error
}

type productRepo struct {
	productList []*Product
}

//constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}


func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList

	return nil
}


func generateInitialProducts(r *productRepo) {
	pd1 := &Product{
		ID:          1,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd3 := &Product{
		ID:          3,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd4 := &Product{
		ID:          4,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}
	pd5 := &Product{
		ID:          5,
		Title:       "Apple",
		Description: "One apple a day, keep the doctor away!",
		Price:       250,
		ImgUrl:      "Img not found",
	}

	r.productList = append(r.productList, pd1)
	r.productList = append(r.productList, pd2)
	r.productList = append(r.productList, pd3)
	r.productList = append(r.productList, pd4)
	r.productList = append(r.productList, pd5)
}

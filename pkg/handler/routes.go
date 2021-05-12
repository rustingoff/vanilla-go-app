package handler

func InitRoutes() {
	MakeRoute("GET", "/product", getAllProducts)
	MakeRoute("GET", "/product/", getProductById)
	MakeRoute("POST", "/product/add", addProduct)
	MakeRoute("PUT", "/product/update/", updateProduct)
	MakeRoute("DELETE", "/product/delete/", deleteProduct)
}

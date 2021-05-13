package handler

func InitRoutes() {
	// Postegres Routes
	MakeRoute("GET", "/pg_product", getAllProductsPG)
	MakeRoute("GET", "/pg_product/", getProductByIdPG)
	MakeRoute("POST", "/pg_product/add", addProductPG)
	MakeRoute("PUT", "/pg_product/update/", updateProductPG)
	MakeRoute("DELETE", "/pg_product/delete/", deleteProductPG)

	// MongoDb routes
	MakeRoute("GET", "/mg_product", getAllProductsMongo)
	MakeRoute("GET", "/mg_product/", getProductByIdMongo)
	MakeRoute("POST", "/mg_product/add", addProductMongo)
	MakeRoute("PUT", "/mg_product/update/", updateProductMongo)
	MakeRoute("DELETE", "/mg_product/delete/", deleteProductMongo)
}

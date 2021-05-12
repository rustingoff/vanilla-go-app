package handler

import (
	"crudSystem/internal/entities"
	"crudSystem/pkg/repository"
	"crudSystem/pkg/service"
	"crudSystem/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var DB = repository.GetDataBase()

// GETTERS
func getAllProducts(w http.ResponseWriter, r *http.Request, urlPattern string) {
	p := []entities.ProductPGDB{}
	res := DB.Find(&p)
	if res.Error != nil {
		w.Write([]byte("Can't get a record\n"))
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	for _, val := range p {
		fmt.Println(val)
	}
	utils.ServerResponse(w, http.StatusOK)
}

func getProductById(w http.ResponseWriter, r *http.Request, urlPattern string) {
	id := strings.TrimPrefix(r.URL.Path, urlPattern)
	if id == "" {
		utils.ServerResponse(w, http.StatusNotFound)
		return
	}
	p := entities.ProductPGDB{}
	res := DB.Where("id = ?", id).Find(&p)
	if res.Error != nil {
		w.Write([]byte("Can't get record\n"))
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	if p.ID == 0 {
		w.Write([]byte("Record doesn't exists\n"))
		utils.ServerResponse(w, http.StatusBadRequest)
		return
	}
	fmt.Println(p)
	utils.ServerResponse(w, http.StatusOK)
}

// SETTERS
func addProduct(w http.ResponseWriter, r *http.Request, urlPattern string) {
	var p *entities.ProductPGDB
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.ServerResponse(w, http.StatusBadRequest)
		log.Fatalln(err.Error())
		return
	}
	err = service.NewProduct(p, DB)
	if err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	fmt.Println(p)
	utils.ServerResponse(w, http.StatusOK)
}

func updateProduct(w http.ResponseWriter, r *http.Request, urlPattern string) {
	id := strings.TrimPrefix(r.URL.Path, urlPattern)
	if len(id) == 0 {
		w.Write([]byte("Incorrect path\n"))
		utils.ServerResponse(w, http.StatusNotFound)
		return
	}

	updP := entities.ProductPGDB{}
	gotP := entities.ProductPGDB{}
	DB.Find(&updP).Where("id = ?", id)
	err := json.NewDecoder(r.Body).Decode(&gotP)
	if err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
	}
	res := DB.Model(&updP).Updates(entities.ProductPGDB{
		Title:       gotP.Title,
		Description: gotP.Title,
		Price:       gotP.Price,
		Stock:       gotP.Stock,
	})

	if res.Error != nil {
		w.Write([]byte("Error: can't update a record\n"))
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	utils.ServerResponse(w, http.StatusOK)
}

func deleteProduct(w http.ResponseWriter, r *http.Request, urlPattern string) {
	id := strings.TrimPrefix(r.URL.Path, urlPattern)
	if len(id) == 0 {
		w.Write([]byte("Incorrect path\n"))
		utils.ServerResponse(w, http.StatusNotFound)
		return
	}
	res := DB.Delete(&entities.ProductPGDB{}, id)
	if res.Error != nil {
		w.Write([]byte("Error: can't delete a record\n"))
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	utils.ServerResponse(w, http.StatusOK)
}

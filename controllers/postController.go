package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hienviluong125/docker_go_mux/database"
	"github.com/hienviluong125/docker_go_mux/models"
)

type Response map[string]interface{}

type PostController struct{}

func (c *PostController) Index(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	if err := database.DB.Find(&posts).Error; err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		"success": true,
		"data":    posts,
	})
}

func (c *PostController) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		"success": true,
		"data":    post,
	})
}

func (c *PostController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	if err := database.DB.Create(&post).Error; err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		"success": true,
		"data":    post,
	})
}

func (c *PostController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var post models.Post

	err := database.DB.First(&post, id).Error
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var updateParams map[string]interface{}
	json.NewDecoder(r.Body).Decode(&updateParams)

	err = database.DB.Model(&post).Updates(updateParams).Error
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		"success": true,
		"data":    post,
	})

}

func (c *PostController) Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var post models.Post

	err := database.DB.First(&post, id).Error
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err = database.DB.Delete(&post).Error
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{"success": true})
}

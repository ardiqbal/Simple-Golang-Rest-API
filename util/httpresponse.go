package util

import "github.com/ardiqbal/Simple-Golang-Rest-API/models"

type PostHttpResponse struct {
	Success bool     `json:"success"`
	Message []string `json:"message"`
}

type GetHttpResponse struct {
	Success bool     `json:"success"`
	Message []string `json:"message"`
	Data    `json:"data"`
}

type Data struct {
	Classes models.Classes `json:"classes"`
}

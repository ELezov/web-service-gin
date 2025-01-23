package service

import (
	"net/http"

	"web-service-gin/internal/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

var wishes = []entity.Wish{
	{
		ID:          "1",
		Title:       "Skateboard",
		Description: "I want to be like Tony Hawk",
		IsReserved:  false,
		Price:       56.99,
		ImageUrl:    "https://cdn.poryadok.ru/upload/iblock/026/026ebf6851233e0a5d9f6eb7be5f37b4.jpg",
		ProductUrl:  "",
	},
	{
		ID:          "2",
		Title:       "Кастрюля",
		Description: "Приготовлю рататуй",
		IsReserved:  false,
		Price:       34.00,
		ImageUrl:    "https://blog.ippon.fr/content/images/2023/09/RGFzaGF0YXJfRGV2ZWxvcGVyX092ZXJJdF9jb2xvcl9QR19zaGFkb3c-.png",
		ProductUrl:  "",
	},
	{
		ID:          "3",
		Title:       "Взрывные котята",
		Description: "Играл у друзей - круто было",
		IsReserved:  true,
		Price:       12.50,
		ImageUrl:    "https://blog.ippon.fr/content/images/2023/09/RGFzaGF0YXJfRGV2ZWxvcGVyX092ZXJJdF9jb2xvcl9QR19zaGFkb3c-.png",
		ProductUrl:  "",
	},
	{
		ID:          "4",
		Title:       "Взрывные котята",
		Description: "Играл у друзей - круто было",
		IsReserved:  true,
		Price:       12.50,
		ImageUrl:    "https://hobbygames.by/image/cache/hobbygames/data/HobbyWorld/Explosive_Kitten/18/Kotyata18_00-1024x1024-wm.jpg",
		ProductUrl:  "",
	},
}

// getWishes responds with the list of all albums as JSON.
func GetWishes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wishes)
}

func GetWishByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range wishes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wish not found"})
}

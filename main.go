// Implimenting Validators with Gin framework

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Making the structure of album
type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required,lte=90"`
}

// Giving the data of the struct to the slice variable
var albums = []album{
	{ID: "1", Title: "Hawayein", Artist: "Aryan Ojha", Price: 30.30},
	{ID: "2", Title: "Pinga Ga Pori", Artist: "Shreya Goshal", Price: 70.12},
	{ID: "3", Title: "Tarak Mantra", Artist: "Anuradha Poudwal", Price: 120.99},
}

// Function to print the albums
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// Function to get the message of error
func errorMessageIntoText(err error) map[string]string {
	errorMessage := make(map[string]string)

	// Variable stores the errors in it
	var storeerr validator.ValidationErrors
	if errors.As(err, &storeerr) {
		for _, errs := range storeerr { // iteration takes place on the storeerr and check for the error
			field := errs.Field()
			tag := errs.Tag()

			switch field {
			case "ID":
				if tag == "required" {
					errorMessage["id"] = "Id is required"
				}
			case "Title":
				if tag == "required" {
					errorMessage["title"] = "Title is required"
				}
			case "Artist":
				if tag == "required" {
					errorMessage["artist"] = "Name of artist is Invalid"
				}
			case "Price":
				if tag == "required" {
					errorMessage["price"] = "Price is invalid"
				} else if tag == "lte" {
					errorMessage["price"] = "Price must be less than or equal to 90"
				}
			default:
				errorMessage[field] = "Invalid Values are given as input"
			}
		}
	} else {
		errorMessage["errors"] = "Invalid Values are given as input"
	}
	return errorMessage
}

// Making a function to add the New Album
func postAlbums(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		errorMessage := errorMessageIntoText(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}
	// 	if errors,ok := err.(validator.ValidationErrors); ok {
	// 		storeerr = errors                                         // storing errors in storeerr
	// 	} else {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message":"Invalid Input format"})
	// 		return
	// 	}
	// 	errMessage := map[string]string{}
	// 	for _,accept := range storeerr{
	// 		field := accept.Field()
	// 		tag := accept.Tag()
	// 		switch field{
	// 		case "ID":
	// 			if tag == "required"{
	// 				errMessage["id"] = "ID is required"
	// 			}
	// 		case "Title":
	// 			if tag == "required"{
	// 				errMessage["title"] = "Title must be present"
	// 			}
	// 		case "Artist":
	// 			if tag == "required"{
	// 				errMessage["artist"] = "Artist is required"
	// 			}
	// 		case "Price":
	// 			if tag == "required" {
	// 				errMessage["price"] = "Invalid Price"
	// 			} else if tag == "lte" {
	// 				errMessage["price"] = "Price must be less than or equal to 90"
	// 			}
	// 		default:
	// 			errMessage[field] = "Invalid Value"
	// 		}
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"Validation errors":errMessage})
	// 	return
	// }

	// appending the new album
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
	ctx.IndentedJSON(http.StatusOK, gin.H{"Message": "Added the album successfully!"})
}

// searching the album by id
func getAlbumsByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid album ID"})
}

// Function to storeerr the album by using it's ID
func removeAlbums(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"message": "Deleted Successfully",
			})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid ID, Album not found"})
}

// Update the existing album
func newAlbum(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateAlbum album
	if err := ctx.ShouldBindJSON(&updateAlbum); err != nil {
		errorMessage := errorMessageIntoText(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		// 		var storeerr validator.ValidationErrors
		// 		if errors,ok := err.(validator.ValidationErrors); ok {
		// 			storeerr = errors
		// 		} else {
		// 			ctx.JSON(http.StatusBadRequest, gin.H{"message":"Invalid Input"})
		// 		}
		// 	errMessage := map[string]string{}
		// 	for _,accept := range storeerr{
		// 		field := accept.Field()
		// 		tag := accept.Tag()
		// 		switch field{
		// 		case "ID":
		// 			if tag == "required"{
		// 				errMessage["id"] = "Invalid ID"
		// 			}
		// 		case "Title":
		// 			if tag == "required"{
		// 				errMessage["title"] = "Required an Valid title"
		// 			}
		// 		case "Artist":
		// 			if tag == "required"{
		// 				errMessage["artist"] = "Artist name is required"
		// 			}
		// 		case "Price":
		// 			if tag == "required" {
		// 				errMessage["price"] = "Need an valid Price"
		// 			} else if tag == "lte"{
		// 				errMessage["price"] = "Price must be less than or equal to 90"
		// 			}
		// 		default:
		// 			errMessage[field] = "Invalid Value"
		// 		}
		// 	}
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"validation_errors": errMessage})
		// 	return
		// }
	}
	for i, a := range albums {
		if a.ID == id {
			albums[i] = updateAlbum
			ctx.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()

	// To see all the entries
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)

	// To add a new album
	router.POST("/albums", postAlbums)

	// To delete an Album
	router.DELETE("/albums/:id", removeAlbums)

	// To update the existing albums
	router.PUT("/albums/:id", newAlbum)

	// Running the serstoreerrr
	router.Run(":8080")
}

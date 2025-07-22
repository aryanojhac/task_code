// Implimenting Validators with Gin framework

package main
import(
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
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Making a function to add the New Album
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil{
		var ve validator.ValidationErrors
		if errors,ok := err.(validator.ValidationErrors); ok {
			ve = errors                                         // storing errors in ve
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid Input fomat"})
			return
		}

		errMessage := map[string]string{}
		for _,e := range ve{
			field := e.Field()
			tag := e.Tag()

			switch field{
			case "ID":
				if tag == "required"{
					errMessage["id"] = "ID is required"
				}
			case "Title":
				if tag == "required"{
					errMessage["title"] = "Title must be present"
				}
			case "Artist":
				if tag == "required"{
					errMessage["artist"] = "Artist is required"
				}
			case "Price":
				if tag == "required" {
					errMessage["price"] = "Invalid Price"
				} else if tag == "lte" {
					errMessage["price"] = "Price must be less than or equal to 90"
				}
			default:
				errMessage[field] = "Invalid Value"
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"Validation errors":errMessage})
		return
	}

	// appending the new album
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Added the album successfully!"})
}

// searching the album by id
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid album ID"})
}

// Function to Remove the album by using it's ID
func removeAlbums(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Deleted Successfully",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid ID, Album not found"})
}

// Update the existing album
func newAlbum(c *gin.Context) {
	id := c.Param("id")
	
	var updateAlbum album
	if err := c.ShouldBindJSON(&updateAlbum); err != nil {
		var ve validator.ValidationErrors
		if errors,ok := err.(validator.ValidationErrors); ok {
			ve = errors
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid Input"})
		}
	
	errMessage := map[string]string{}
	for _,e := range ve{
		field := e.Field()
		tag := e.Tag()

		switch field{
		case "ID": 
			if tag == "required"{
				errMessage["id"] = "Invalid ID"	
			}
		case "Title":
			if tag == "required"{
				errMessage["title"] = "Required an Valid title"
			}
		case "Artist":
			if tag == "required"{
				errMessage["artist"] = "Artist name is required"
			}
		case "Price":
			if tag == "required" {
				errMessage["price"] = "Need an valid Price"
			} else if tag == "lte"{
				errMessage["price"] = "Price must be less than or equal to 90"
			}
		default:
			errMessage[field] = "Invalid Value"
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errMessage})
	return
}

	for i, a := range albums {
		if a.ID == id {
			albums[i]= updateAlbum
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	r := gin.Default()

	// To see all the entries
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumsByID)

	// To add a new album
	r.POST("/albums", postAlbums)

	// To delete an Album
	r.DELETE("/albums/:id", removeAlbums)

	// To update the existing albums
	r.PUT("/albums/:id", newAlbum)

	// Running the server
	r.Run(":8080")
}
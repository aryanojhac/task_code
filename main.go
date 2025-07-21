package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

// Makin the structure of album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
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
	if err := c.BindJSON(&newAlbum); err != nil {
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
	if err := c.BindJSON(&updateAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	for i, a := range albums {
		if a.ID == id {
			albums[i].Title = updateAlbum.Title
			albums[i].Artist = updateAlbum.Artist
			albums[i].Price = updateAlbum.Price
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
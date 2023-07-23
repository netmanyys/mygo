package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    router.DELETE("/albums/:id", deleteAlbumByID) // New route for deleting an album by ID

    router.PUT("/albums/:id", updateAlbumByID) // New route for updating an album by ID

    router.Run("localhost:8888")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


// deleteAlbumByID removes an album from the albums slice based on the given ID.
func deleteAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums to find the index of the album with the given ID.
    for i, a := range albums {
        if a.ID == id {
            // Remove the album from the slice by slicing it out.
            albums = append(albums[:i], albums[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// updateAlbumByID updates an album in the albums slice based on the given ID.
func updateAlbumByID(c *gin.Context) {
    id := c.Param("id")
    var updatedAlbum album

    // Call BindJSON to bind the received JSON to updatedAlbum.
    if err := c.BindJSON(&updatedAlbum); err != nil {
        return
    }

    // Loop through the list of albums to find the index of the album with the given ID.
    for i, a := range albums {
        if a.ID == id {
            // Update the album in the slice with the new information.
            albums[i] = updatedAlbum
            c.IndentedJSON(http.StatusOK, gin.H{"message": "album updated"})
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

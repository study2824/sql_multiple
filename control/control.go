package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sql_multiple1/db"
)

type Controller struct{}

// Get All Users
func (Controller) GetAllUsers(c *gin.Context) {
	users, err := db.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Get One User
func (Controller) GetOneUser(c *gin.Context) {
	id := c.Param("id")
	user, err := db.GetOneUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Add User
func (Controller) AddUser(c *gin.Context) {
	name := c.PostForm("name")
	err := db.AddUser(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"created user": name,
	})
}

// Delete User
func (Controller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	deletedUser, err := db.GetOneUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted user": deletedUser})
}

// Get All Tags
func (Controller) GetAllTags(c *gin.Context) {
	tags, err := db.GetAllTags()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tags": tags})
}

// Get One Tag
func (Controller) GetOneTag(c *gin.Context) {
	id := c.Param("id")
	tag, err := db.GetOneTag(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tag": tag})
}

//Add Tag
func (Controller) AddTag(c *gin.Context) {
	name := c.PostForm("name")
	err := db.AddTag(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"created tag": name})
}

// Delete tag
func (Controller) DeleteTag(c *gin.Context) {
	id := c.Param("id")
	deletedTag, err := db.GetOneTag(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.DeleteTag(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"deleted tag": deletedTag})
}

// Combine
func (Controller) UserTagCombine(c *gin.Context) {
	userID := c.PostForm("userID")
	tagID := c.PostForm("tagID")

	err := db.UserTagCombine(userID, tagID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.GetOneUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag, err := db.GetOneTag(tagID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"combine": user.Name + ", " + tag.Name})
}

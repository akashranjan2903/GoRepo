package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocrud/intializer"
	"github.com/gocrud/models"
)

func main() {

	intializer.Databasecon()

	insertUser := models.User{Name: "test"}

	result := intializer.DB.Create(&insertUser)
	if result.Error != nil {
		println("error")
	}
	println("count of entered data is:", result.RowsAffected)
	// read first data by primary key
	var read_user models.User
	result1 := intializer.DB.First(&read_user, 1)
	println("count of first data is", result1.RowsAffected)
	fmt.Printf("First Record is: %v\n", read_user)
	//   read by column name

	result2 := intializer.DB.First(&read_user, "name = ?", "Jinzhu")
	println("count of First data by name:", result2.RowsAffected)
	fmt.Printf("Read First data by name: %v\n", read_user)

	//  update user
	var update_user models.User

	intializer.DB.First(&update_user)

	update_user.Name = "Akash"
	updated_Rows := intializer.DB.Save(&update_user)
	println(" updated user by primary key count:", updated_Rows.RowsAffected)
	fmt.Printf("updated user by primary key : %v\n", update_user.Name)

	affected_row := intializer.DB.Model(&models.User{}).Where("name = ?", "Akash").Update("name", "random")
	println("update user with gven condn count", affected_row.RowsAffected)
	fmt.Printf("update user with gven condn is : %v\n", update_user.Name)

	//  delete operation

	deleted_row := intializer.DB.Where("name LIKE ?", "%random%").Delete(&models.User{})
	println("Total deleted Rows are:", deleted_row.RowsAffected)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(gin.Recovery())
	r.Run(fmt.Sprintf(":%d", 8080))
	// r.Run()

}

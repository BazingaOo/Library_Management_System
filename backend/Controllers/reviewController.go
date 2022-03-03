package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddReview(c *gin.Context) {
	var review Models.Review
	review.Review_id, _ = strconv.Atoi(c.PostForm("reviewId"))
	review.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	review.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	review.Review_content = c.PostForm("reviewContent")
	res := Models.AddReview(review)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}

func UpdateReview(c *gin.Context) {
	var review Models.Review
	review.Review_id, _ = strconv.Atoi(c.PostForm("reviewId"))
	review.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	review.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	review.Review_content = c.PostForm("reviewContent")
	res := Models.UpdateReview(review)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}

func DeleteReview(c *gin.Context) {
	var review Models.Review
	review.Review_id, _ = strconv.Atoi(c.PostForm("reviewID"))
	res := Models.DeleteReview(review.Review_id)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}

func ReadReview(c *gin.Context) {
	var review Models.Review
	review.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	res := Models.ReadReview(review.User_id)
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}
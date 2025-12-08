package controllers

import (
	"context"
	"time"

	"github.com/Amdadul-HQ/go_lang_app.git/go_app/server/movie_server/database"
	"github.com/Amdadul-HQ/go_lang_app.git/go_app/server/movie_server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies") 

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		var movies []models.Movie

		cursor,err := movieCollection.Find(ctx,bson.M{})
	}
}

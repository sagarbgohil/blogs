package blogs

import (
	"context"

	"github.com/sagarbgohil/go-backend/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// var collection *mongo.Collection = config.DbClient.Collection("blogs")

func collection() *mongo.Collection {
	return config.DbClient.Collection("blogs")
}


func insertOne(blog Blog) (string, error) {
	// Insert the blog into the collection
	insertedID, err := collection().InsertOne(context.Background(), blog)
	if err != nil {
		return "", err
	}

	// Return the inserted ID as a string
	return insertedID.InsertedID.(string), nil
}

func updateOne(movieId string, blog Blog) (string, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := primitive.M{"_id": id}
	update := primitive.M{"$set": blog}

	// Update the blog in the collection
	result, err := collection().UpdateOne(context.Background(), filter, update)

	if err != nil {
		return "", err
	}

	if result.MatchedCount == 0 {
		return "", mongo.ErrNoDocuments
	}
	return result.UpsertedID.(string), nil
}

func deleteOne(blogId string) (int64, error) {
	id, _ := primitive.ObjectIDFromHex(blogId)

	filter := primitive.M{"_id": id}

	// Delete the blog from the collection
	result, err := collection().DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func deleteMany(blogIds []string) (int64, error) {
	var ids []primitive.ObjectID
	for _, id := range blogIds {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return 0, err
		}
		ids = append(ids, objID)
	}

	filter := primitive.M{"_id": primitive.M{"$in": ids}}

	// Delete multiple blogs from the collection
	result, err := collection().DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func getAll() ([]Blog, error) {
	var blogs []Blog

	// Find all blogs in the collection
	cursor, err := collection().Find(context.Background(), primitive.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode each blog into the blogs slice
	for cursor.Next(context.Background()) {
		var blog Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func getById(blogId string) (Blog, error) {
	id, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return Blog{}, err
	}

	filter := primitive.M{"_id": id}

	var blog Blog

	// Find the blog by ID in the collection
	err = collection().FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		return Blog{}, err
	}

	return blog, nil
}
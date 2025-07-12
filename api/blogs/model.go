package blogs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Author    string             `bson:"author" json:"author"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
	Published bool               `bson:"published" json:"published"`
	Tags      []string           `bson:"tags" json:"tags"`
	Thumbnail string             `bson:"thumbnail" json:"thumbnail"`
	Slug      string             `bson:"slug" json:"slug"`
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`
	Language  string             `bson:"language" json:"language"`
}

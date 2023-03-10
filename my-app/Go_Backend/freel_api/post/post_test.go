package post_test

import (
	"context"
	"fmt"

	"Freel.com/freel_api/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"

    "testing"
)

type Like struct {
	Username string `bson:"username,omitempty" json:"username"`
	Date     string `bson:"date,omitempty" json:"date"`
}

type Comment struct {
	Username string `bson:"username,omitempty" json:"username"`
	Date     string `bson:"date,omitempty" json:"date"`
	Comment  string `bson:"comment,omitempty" json:"comment"`
}

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Tags     []string  `json:"tags"`
	Date     string    `json:"date"`
	Image    string    `json:"image"`
	Likes    []Like    `bson:"likes,omitempty" json:"likes"`
	Comments []Comment `bson:"comments,omitempty" json:"comments"`
}

type Location struct {
	Type        string    `bson:"type,omitempty" json:"type"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name,omitempty" json:"name"`
	Bio            string             `bson:"bio,omitempty" json:"bio"`
	ProfilePicture string             `bson:"profilepicture,omitempty" json:"profilepicture"`
	Posts          []Post             `bson:"posts,omitempty" json:"posts"`
	Location       Location           `bson:"location,omitempty" json:"location"`
	SavedPosts     []Post             `bson:"saved_post,omitempty" json:"saved_post"`
}

func Create_Account(user User) {

	client := mongo.GetMongoClient()

	collection := client.Database("freel").Collection("users")

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User inserted successfully")

}


func Test_Create_Fake_Account(t *testing.T) {
	fmt.Println("Creating sample data to insert into mongo since no data was given")

	post1 := Post{
		Title: "My First Post",
		Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Tags:  []string{"programming", "photography"},
		Date:  "2022-01-01T12:00:00Z",
		Image: "https://example.com/post1.jpg",
	}

	post2 := Post{
		Title: "My Second Post",
		Body:  "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Tags:  []string{"travel", "food"},
		Date:  "2022-01-05T12:00:00Z",
		Image: "https://example.com/post2.jpg",
	}

	test_user := User{
		Name:           "Eric fake",
		Bio:            "I'm a software engineer and hobbyist photographer.",
		ProfilePicture: "https://example.com/profile.jpg",
		Posts: []Post{
			post1,
			post2,
		},
		Location: Location{
			Type:        "Point",
			Coordinates: []float64{-122.4194, 37.7749},
		},
		SavedPosts: []Post{
			post1,
		},
	}

	Create_Account(test_user)
}

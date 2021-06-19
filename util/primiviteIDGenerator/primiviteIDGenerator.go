package primiviteIDGenerator

import "go.mongodb.org/mongo-driver/bson/primitive"

//GenerateID Generate 12 byte UUID. It will return 24 hexadimal characters
func GenerateID() string {
	return primitive.NewObjectID().Hex()
}

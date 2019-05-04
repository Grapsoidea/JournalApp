package model

import (
	"context"
	"errors"
	"log"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ItemInfo godoc
type ItemInfo struct {
	Name   string   `bson:"name" json:"name" example:"scale"`
	Fields []string `bson:"fields" json:"fields"`
}

func CheckIn(el string, list []string) bool{
	for i:=0; i<len(list); i++{
		if el == list[i]{
			return true
		}
	}
	return false
}

//Errors godoc
var (
	ErrIDInvalid = errors.New("ID is empty")
	ErrNameInvalid = errors.New("name is empty")
	ErrTitleInvalid = errors.New("title is empty")
	ErrDeletedInvalid = errors.New("deleted is true")
	ErrFieldsInvalid = errors.New("err in Fields")
	ErrFieldsNameInvalid = errors.New("Fields Name is empty")
	ErrFieldsTypeInvalid = errors.New("type Name is empty")
)

// ItemField godoc
type ItemField struct {
	Name  string `bson:"name" json:"name" example:"serial_number"`
	Title string `bson:"title" json:"title" example:"Серийный номер"`
	Type  string `bson:"type" json:"type" example:"String"`
}

// ItemScheme godoc
type ItemScheme struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id" example:"5ca10d9d015c736a72b7b3ba"`
	Name    string             `bson:"name" json:"name" example:"scale"`
	Title   string             `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField        `bson:"fields" json:"fields"`
	Deleted bool               `bson:"deleted" json:"-"`
}

// NewItemScheme godoc
type NewItemScheme struct {
	Name    string      `bson:"name" json:"name" example:"scale"`
	Title   string      `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField `bson:"fields" json:"fields"`
	Deleted bool        `bson:"deleted" json:"-"`
}

// UpdateItemScheme godoc
type UpdateItemScheme struct {
	Name    string      `bson:"name" json:"name" example:"scale"`
	Title   string      `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField `bson:"fields" json:"fields"`
	Deleted bool        `bson:"deleted" json:"-"`
}

// Insert godoc
func (s NewItemScheme) Insert() error {
	insertResault, err := ItemSchemeCollection().InsertOne(context.Background(), s)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Inserted documents: ", insertResault.InsertedID)
	return err
}

// Update godoc
func (s UpdateItemScheme) Update(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := ItemSchemeCollection().UpdateOne(context.Background(), bson.D{{"_id", objectID}}, bson.D{{"$set", s}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("updated documents: ", updateResault.UpsertedID)
	return err
}

// Validation godoc
func (s NewItemScheme) Validation() error {
	Types := []string{"Integer", "Dooble", "String", "Boolean", "Array", "Signature", "Date", "ObjectId"}
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	case len(s.Title) == 0:
		return ErrTitleInvalid
	case s.Deleted == true:
		return ErrDeletedInvalid
	case s.Fields == nil:
		return ErrFieldsInvalid
	case s.Fields != nil:
		for i:= 0; i < len(s.Fields); i++ {
			switch { 
				case len(s.Fields[i].Name) == 0:
					return ErrFieldsInvalid
				case len(s.Fields[i].Title) == 0:
					return ErrFieldsInvalid
				case len(s.Fields[i].Type) == 0:
					return ErrFieldsInvalid
				case !(CheckIn(s.Fields[i].Type, Types)):
					return ErrTypeInvalid
			}
		}
		return nil
	default:
		return nil
	}
}

// Validation godoc
func (s UpdateItemScheme) Validation() error {
	Types := []string{"Integer", "Dooble", "String", "Boolean", "Array", "Signature", "Date", "ObjectId"}
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	case len(s.Title) == 0:
		return ErrTitleInvalid
	case s.Deleted == true:
		return ErrDeletedInvalid
	case s.Fields == nil:
		return ErrFieldsInvalid
	case s.Fields != nil:
		for i:= 0; i < len(s.Fields); i++ {
			switch { 
				case len(s.Fields[i].Name) == 0:
					return ErrFieldsInvalid
				case len(s.Fields[i].Title) == 0:
					return ErrFieldsInvalid
				case len(s.Fields[i].Type) == 0:
					return ErrFieldsInvalid
				case !(CheckIn(s.Fields[i].Type, Types)):
					return ErrTypeInvalid
			}
		}
		return nil
	default:
		return nil
	}
}

// ItemSchemeCollection godoc
func ItemSchemeCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("itemScheme")

	return coll
}

// SomeAdd godoc
func SomeAdd() ItemScheme {
	scaleScheme := ItemScheme{
		Name:  "scale",
		Title: "Весы",
		Fields: []ItemField{
			{
				Name:  "name",
				Title: "Название",
				Type:  "String",
			},
			{
				Name:  "serial_number",
				Title: "Серийный номер",
				Type:  "String",
			},
			{
				Name:  "min_w",
				Title: "Минимальный вес",
				Type:  "String",
			},
		},
	}

	insertResault, err := ItemSchemeCollection().InsertOne(context.Background(), scaleScheme)

	if err != nil {
		log.Println(err)
		return ItemScheme{}
	}

	log.Println("Inserted documents: ", insertResault.InsertedID)

	return scaleScheme
}

//ItemSchemeAll get list item schemes godoc
func ItemSchemeAll() ([]ItemScheme, error) {
	cur, err := ItemSchemeCollection().Find(context.Background(), bson.D{{"deleted", false}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	listSchemes := []ItemScheme{}

	for cur.Next(context.Background()) {
		row := new(ItemScheme)
		err := cur.Decode(&row)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listSchemes = append(listSchemes, *row)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return listSchemes, err
}

//ItemSchemeOne get list item schemes with id godoc
func ItemSchemeOne(id string) (ItemScheme, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return ItemScheme{}, err
	}
	row := new(ItemScheme)
	err = ItemSchemeCollection().FindOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", objectID}}, bson.D{{"deleted", false}}}}}).Decode(&row)
	if err != nil {
		log.Println(err)
		return ItemScheme{}, err
	}

	return *row, err
}

// DeleteSchemeOne godoc
func DeleteSchemeOne(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := ItemSchemeCollection().UpdateOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", objectID}}, bson.D{{"deleted", false}}}}}, bson.D{{"$set", bson.D{{"deleted", true}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("deleted documents: ", updateResault.UpsertedID)
	return err
}

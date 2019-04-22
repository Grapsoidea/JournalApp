package model

import (
	"context"
	"log"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ReportField godoc
type ReportField struct {
	Title string `bson:"title" json:"title" example:"Дата"`
	Value string `bson:"value" json:"value" example:"{journal.date}"`
}

// ReportScheme godoc
type ReportScheme struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id" example:"5ca10d9d015c736a72b7b3ba"`
	Name    string             `bson:"name" json:"name" example:"scales_calibration"`
	Title   string             `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Journal string             `bson:"journal" json:"journal" example:"scales_calibration"`
	Fields  []ReportField      `bson:"fields" json:"fields"`
	Deleted bool               `bson:"deleted" json:"-"`
}

// ReportScheme godoc
type NewReportScheme struct {
	Name    string        `bson:"name" json:"name" example:"scales_calibration"`
	Title   string        `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Journal string        `bson:"journal" json:"journal" example:"scales_calibration"`
	Fields  []ReportField `bson:"fields" json:"fields"`
	Deleted bool          `bson:"deleted" json:"-"`
}

// ReportScheme godoc
type UpdateReportScheme struct {
	Name    string        `bson:"name" json:"name" example:"scales_calibration"`
	Title   string        `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Journal string        `bson:"journal" json:"journal" example:"scales_calibration"`
	Fields  []ReportField `bson:"fields" json:"fields"`
	Deleted bool          `bson:"deleted" json:"-"`
}

// ReportSchemeCollection godoc
func ReportSchemeCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("reportScheme")

	return coll
}

//ReportSchemeAll get list report schemes godoc
func ReportSchemeAll() ([]ReportScheme, error) {
	cur, err := ReportSchemeCollection().Find(context.Background(), bson.D{{"deleted", false}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	listSchemes := []ReportScheme{}

	for cur.Next(context.Background()) {
		row := new(ReportScheme)
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

//ReportSchemeOne get list report schemes with id godoc
func ReportSchemeOne(id string) (ReportScheme, error) {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return ReportScheme{}, err
	}
	row := new(ReportScheme)
	err = ReportSchemeCollection().FindOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", ojectID}}, bson.D{{"deleted", false}}}}}).Decode(&row)
	if err != nil {
		log.Println(err)
		return ReportScheme{}, err
	}

	return *row, err
}

// Insert godoc
func (s NewReportScheme) Insert() error {
	insertResault, err := ReportSchemeCollection().InsertOne(context.Background(), s)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Inserted documents: ", insertResault.InsertedID)
	return err
}

// Validation godoc
func (s NewReportScheme) Validation() error {
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// Update godoc
func (s UpdateReportScheme) Update(id string) error {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := ReportSchemeCollection().UpdateOne(context.Background(), bson.D{{"_id", ojectID}}, bson.D{{"$set", s}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("updated documents: ", updateResault.UpsertedID)
	return err
}

// Validation godoc
func (s UpdateReportScheme) Validation() error {
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// DeleteReportSchemeOne godoc
func DeleteReportSchemeOne(id string) error {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := ReportSchemeCollection().UpdateOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", ojectID}}, bson.D{{"deleted", false}}}}}, bson.D{{"$set", bson.D{{"deleted", true}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("deleted documents: ", updateResault.UpsertedID)
	return err
}

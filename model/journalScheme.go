package model

import (
	"context"
	"log"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// JournalIf godoc
type JournalIf struct {
	Fields []string `bson:"fields" json:"fields" example:["result", "value"]` //Проверяемые поля
}

// JournalComputed godoc
type JournalComputed struct {
	Type  string `bson:"type" json:"type" example:"deviation"`
	Field string `bson:"field" json:"field" example:"result"`

	// Norm Если type deviation
	Norm      *string `bson:"norm,omitempty" json:"norm,omitempty" example:"giri_w"`                   // Нормальный вес
	Deviation *string `bson:"deviation,omitempty" json:"deviation,omitempty" example:"norm_deviation"` // Допустимое отклонение

	// Range Если type range
	Range *[2]string `bson:"range,omitempty" json:"range,omitempty" example:""` // Допустимы предел

	// Value Если type equals
	Value *string `bson:"value,omitempty" json:"value,omitempty" example:""` // Значение, которому должно быть равно

	// Max Если type less
	Max *string `bson:"max,omitempty" json:"max,omitempty" example:""` // Допустимое максимальное значение

	// Min Если type more
	Min *string `bson:"min,omitempty" json:"min,omitempty" example:""` // Допустимое минимальное значение

	// ID Если type more_than
	ID *string `bson:"id,omitempty" json:"id,omitempty" example:""` // Идентификатор блока
	On *string `bson:"on,omitempty" json:"on,omitempty" example:""` // Значение, на которое должен быть больше

	// Enum Если type enum

	// Массив допустимых значений. Если выбранное
	// значение не принадлежит данному массиву, то
	// check = false
	Enum *[]string `bson:"enum,omitempty" json:"enum,omitempty" example:""`
}

// JournalField godoc
type JournalField struct {
	Name  string `bson:"name" json:"name" example:"serial_number"`
	Title string `bson:"title" json:"title" example:"Серийный номер"`
	Type  string `bson:"type" json:"type" example:"String"`

	// Computed вычесляемое поле с переменным количеством полей
	Computed *JournalComputed `bson:"computed,omitempty" json:"computed,omitempty"`

	// If непонятно какое условие
	If *JournalIf `bson:"if,omitempty" json:"if,omitempty"`
}

// JournalScheme godoc
type JournalScheme struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id" example:"5ca10d9d015c736a72b7b3ba"`
	Name     string             `bson:"name" json:"name" example:"scales_calibration"`
	Title    string             `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Daily    bool               `bson:"daily" json:"daily" example:"true"`
	Fixed    bool               `bson:"fixed" json:"fixed" example:"true"`
	Item     string             `bson:"item" json:"item" example:"scale"`
	ItemInfo *[]string          `bson:"item_info" json:"item_info" example:"["name", "min_w", "max_w", "giri_w", "norm_deviation"]"`
	Fields   []JournalField     `bson:"fields" json:"fields"`
	Deleted  bool               `bson:"deleted" json:"-"`
}

// NewJournalScheme godoc
type NewJournalScheme struct {
	Name     string         `bson:"name" json:"name" example:"scales_calibration"`
	Title    string         `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Daily    bool           `bson:"daily" json:"daily" example:"true"`
	Fixed    bool           `bson:"fixed" json:"fixed" example:"true"`
	Item     string         `bson:"item" json:"item" example:"scale"`
	ItemInfo *[]string      `bson:"item_info" json:"item_info" example:"["name", "min_w", "max_w", "giri_w", "norm_deviation"]"`
	Fields   []JournalField `bson:"fields" json:"fields"`
	Deleted  bool           `bson:"deleted" json:"-"`
}

// UpdateJournalScheme godoc
type UpdateJournalScheme struct {
	Name     string         `bson:"name" json:"name" example:"scales_calibration"`
	Title    string         `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Daily    bool           `bson:"daily" json:"daily" example:"true"`
	Fixed    bool           `bson:"fixed" json:"fixed" example:"true"`
	Item     string         `bson:"item" json:"item" example:"scale"`
	ItemInfo *[]string      `bson:"item_info" json:"item_info" example:"["name", "min_w", "max_w", "giri_w", "norm_deviation"]"`
	Fields   []JournalField `bson:"fields" json:"fields"`
	Deleted  bool           `bson:"deleted" json:"-"`
}

// JournalSchemeCollection godoc
func JournalSchemeCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("journalScheme")

	return coll
}

//JournalSchemeAll get list Journak schemes godoc
func JournalSchemeAll() ([]JournalScheme, error) {
	cur, err := JournalSchemeCollection().Find(context.Background(), bson.D{{"deleted", false}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	listSchemes := []JournalScheme{}

	for cur.Next(context.Background()) {
		row := new(JournalScheme)
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

//JournalSchemeOne get list journal schemes with id godoc
func JournalSchemeOne(id string) (JournalScheme, error) {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return JournalScheme{}, err
	}
	row := new(JournalScheme)
	err = JournalSchemeCollection().FindOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", ojectID}}, bson.D{{"deleted", false}}}}}).Decode(&row)
	if err != nil {
		log.Println(err)
		return JournalScheme{}, err
	}

	return *row, err
}

// Insert godoc
func (s NewJournalScheme) Insert() error {
	insertResault, err := JournalSchemeCollection().InsertOne(context.Background(), s)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Inserted documents: ", insertResault.InsertedID)
	return err
}

// Validation godoc
func (s NewJournalScheme) Validation() error {
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// Update godoc
func (s UpdateJournalScheme) Update(id string) error {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := JournalSchemeCollection().UpdateOne(context.Background(), bson.D{{"_id", ojectID}}, bson.D{{"$set", s}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("updated documents: ", updateResault.UpsertedID)
	return err
}

// Validation godoc
func (s UpdateJournalScheme) Validation() error {
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// DeleteJournalSchemeOne godoc
func DeleteJournalSchemeOne(id string) error {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := JournalSchemeCollection().UpdateOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", ojectID}}, bson.D{{"deleted", false}}}}}, bson.D{{"$set", bson.D{{"deleted", true}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("deleted documents: ", updateResault.UpsertedID)
	return err
}

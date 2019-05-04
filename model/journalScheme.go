package model

import (
	"strconv"
	"context"
	"errors"
	"log"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Errors godoc
var (
	ErrDailyInvalid = errors.New("Daily is empty")
	ErrFixedInvalid = errors.New("Fixed is empty")
	ErrItemInvalid = errors.New("Item is empty")
	ErrItemInfoInvalid = errors.New("ItemInfo is empty")
	ErrComputedInvalid = errors.New("computed is bad")
	ErrDeviatonTypeInvalid = errors.New("error in deviaton type")
	ErrRangeTypeInvalid = errors.New("error in range type")
	ErrEqualsTypeInvalid = errors.New("error in equals type")
	ErrLessTypeInvalid = errors.New("error in less type")
	ErrMoreTypeInvalid = errors.New("error in more type")
	ErrMore_ThanTypeInvalid = errors.New("error in more_than type")
	ErrIfInvalid = errors.New("error in if field")
	ErrComputedTypeInvalid = errors.New("Computed Type isn't exits")
	ErrTypeInvalid = errors.New("Type isn't exits")
	ErrNegativeParam = errors.New("param is negative")
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
	Range *string `bson:"range,omitempty" json:"range,omitempty" example:""` // Допустимы предел

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

	// Массив допустимых значений. Если выбранное || len(s.Fields[i].Computed.Field) == 0
	// значение не принадлежит данному массиву, т || len(s.Fields[i].Computed.Field) == 0
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

	// If условие при котором будут отображаться дополнительные поля
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
func JournalSchemeAll(offset string, limit string) ([]JournalScheme, error) {
	offsetInt, err := strconv.ParseInt(offset, 10, 64)
	if err != nil{
		return nil, err
	}
	if offsetInt < 0{
		return nil, ErrNegativeParam
	} 
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil{
		return nil, err
	}
	if limitInt < 0{
		return nil, ErrNegativeParam
	} 

	options := options.Find()
	options.SetLimit(limitInt)
	options.SetSkip(offsetInt)

	cur, err := JournalSchemeCollection().Find(context.Background(), bson.D{{"deleted", false}}, options)
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return JournalScheme{}, err
	}
	row := new(JournalScheme)
	err = JournalSchemeCollection().FindOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", objectID}}, bson.D{{"deleted", false}}}}}).Decode(&row)
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
	ComputedTypes := []string{"deviation","range","equals","less","more","more_than"}
	Types := []string{"Integer", "Dooble", "String", "Boolean", "Array", "Signature", "Date", "ObjectId"}
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	case len(s.Title) == 0:
		return ErrTitleInvalid
	// case s.Daily == nil:
	// 	return ErrDailyInvalid
	// case s.Fixed == nil:
	// 	return ErrFixedInvalid
	case len(s.Item) == 0:
		return ErrItemInvalid
	case s.ItemInfo == nil:
		return ErrItemInfoInvalid
	case s.Deleted == true:
		return ErrDeletedInvalid
	case s.Fields == nil:
		return ErrFieldsInvalid
	case s.Fields != nil:
		for i:= 0; i < len(s.Fields); i++ {
			switch { 
				case len(s.Fields[i].Name) == 0:
					return ErrFieldsNameInvalid
				case len(s.Fields[i].Title) == 0:
					return ErrFieldsTitleInvalid
				case len(s.Fields[i].Type) == 0:
					return ErrFieldsTypeInvalid
				case !(CheckIn(s.Fields[i].Type, Types)):
					return ErrTypeInvalid
				case s.Fields[i].Computed != nil || s.Fields[i].If != nil:
					if s.Fields[i].Computed != nil {
						switch {
						case len(s.Fields[i].Computed.Type) == 0:
							return ErrComputedInvalid
						case len(s.Fields[i].Computed.Field) == 0:
							return ErrComputedInvalid
						case s.Fields[i].Computed.Type == "deviation" && (s.Fields[i].Computed.Deviation == nil || s.Fields[i].Computed.Norm == nil):
							return ErrDeviatonTypeInvalid
						case s.Fields[i].Computed.Type == "range" && (s.Fields[i].Computed.Range == nil):
							return ErrRangeTypeInvalid
						case s.Fields[i].Computed.Type == "equals" && (s.Fields[i].Computed.Value == nil):
							return ErrEqualsTypeInvalid
						case s.Fields[i].Computed.Type == "less" && (s.Fields[i].Computed.Max == nil):
							return ErrLessTypeInvalid
						case s.Fields[i].Computed.Type == "more" && (s.Fields[i].Computed.Min == nil):
							return ErrMoreTypeInvalid
						case s.Fields[i].Computed.Type == "more_than" && (s.Fields[i].Computed.ID == nil || s.Fields[i].Computed.On == nil):
							return ErrMore_ThanTypeInvalid
						case !(CheckIn(s.Fields[i].Computed.Type, ComputedTypes)):
							return ErrComputedTypeInvalid
						} 
					}
					if s.Fields[i].If != nil{	
						if s.Fields[i].If.Fields == nil{
							return ErrIfInvalid
						}
						for k:= 0; k < len(s.Fields[i].If.Fields); k++ {
							if len(s.Fields[i].If.Fields[k]) == 0{
								return ErrIfInvalid
							}
						}

						
					}
			}
		}
		return nil
	default:
		return nil
	}
}

// Update godoc
func (s UpdateJournalScheme) Update(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := JournalSchemeCollection().UpdateOne(context.Background(), bson.D{{"_id", objectID}}, bson.D{{"$set", s}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("updated documents: ", updateResault.UpsertedID)
	return err
}

// Validation godoc
func (s UpdateJournalScheme) Validation() error {
	ComputedTypes := []string{"deviation","range","equals","less","more","more_than"}
	Types := []string{"Integer", "Dooble", "String", "Boolean", "Array", "Signature", "Date", "ObjectId"}
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	case len(s.Title) == 0:
		return ErrTitleInvalid
	// case s.Daily == nil:
	// 	return ErrDailyInvalid
	// case s.Fixed == nil:
	// 	return ErrFixedInvalid
	case len(s.Item) == 0:
		return ErrItemInvalid
	case s.ItemInfo == nil:
		return ErrItemInfoInvalid
	case s.Deleted == true:
		return ErrDeletedInvalid
	case s.Fields == nil:
		return ErrFieldsInvalid
	case s.Fields != nil:
		for i:= 0; i < len(s.Fields); i++ {
			switch { 
				case len(s.Fields[i].Name) == 0:
					return ErrFieldsNameInvalid
				case len(s.Fields[i].Title) == 0:
					return ErrFieldsTitleInvalid
				case len(s.Fields[i].Type) == 0:
					return ErrFieldsTypeInvalid
				case !(CheckIn(s.Fields[i].Type, Types)):
					return ErrTypeInvalid
				case s.Fields[i].Computed != nil || s.Fields[i].If != nil:
					if s.Fields[i].Computed != nil {
						switch {
						case len(s.Fields[i].Computed.Type) == 0:
							return ErrComputedInvalid
						case len(s.Fields[i].Computed.Field) == 0:
							return ErrComputedInvalid
						case s.Fields[i].Computed.Type == "deviation" && (s.Fields[i].Computed.Deviation == nil || s.Fields[i].Computed.Norm == nil):
							return ErrDeviatonTypeInvalid
						case s.Fields[i].Computed.Type == "range" && (s.Fields[i].Computed.Range == nil):
							return ErrRangeTypeInvalid
						case s.Fields[i].Computed.Type == "equals" && (s.Fields[i].Computed.Value == nil):
							return ErrEqualsTypeInvalid
						case s.Fields[i].Computed.Type == "less" && (s.Fields[i].Computed.Max == nil):
							return ErrLessTypeInvalid
						case s.Fields[i].Computed.Type == "more" && (s.Fields[i].Computed.Min == nil):
							return ErrMoreTypeInvalid
						case s.Fields[i].Computed.Type == "more_than" && (s.Fields[i].Computed.ID == nil || s.Fields[i].Computed.On == nil):
							return ErrMore_ThanTypeInvalid
						case !(CheckIn(s.Fields[i].Computed.Type, ComputedTypes)):
							return ErrComputedTypeInvalid
						} 
					}
					if s.Fields[i].If != nil{	
						if s.Fields[i].If.Fields == nil{
							return ErrIfInvalid
						}
						for k:= 0; k < len(s.Fields[i].If.Fields); k++ {
							if len(s.Fields[i].If.Fields[k]) == 0{
								return ErrIfInvalid
							}
						}

						
					}
			}
		}
		return nil
	default:
		return nil
	}
}

// DeleteJournalSchemeOne godoc
func DeleteJournalSchemeOne(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := JournalSchemeCollection().UpdateOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", objectID}}, bson.D{{"deleted", false}}}}}, bson.D{{"$set", bson.D{{"deleted", true}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("deleted documents: ", updateResault.UpsertedID)
	return err
}

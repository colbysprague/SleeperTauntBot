package collectiondefs

import (
	"fmt"
	"log"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

type FunctionType func()

func InitAllCollectionsForPocketBaseApp(app core.App) error {

	collectionFuncs := map[string]func() *models.Collection{
		"players":        GetPlayersCollection,
		"players_scores": GetPlayersScoreCollection,
	}

	for name, getCollection := range collectionFuncs {
		collection := getCollection()
		err := InitCollectionForPocketBaseApp(app, name, collection)
		if err != nil {
			log.Fatalf("Error initializing collection %s: %v", name, err)
		}
	}

	return nil
}

func InitCollectionForPocketBaseApp(app core.App, collectionName string, collection *models.Collection) error {

	queriedCollection, _ := app.Dao().FindCollectionByNameOrId(collectionName)

	if queriedCollection == nil {
		// there is no existing colleciton
		fmt.Printf("Creating new collection %s.", collectionName)
		return app.Dao().SaveCollection(collection)
	}

	return nil
}

func GetPlayersScoreCollection() *models.Collection {
	var userRequiredRule = ""
	return &models.Collection{
		Name:       "players_scores",
		System:     false,
		CreateRule: &userRequiredRule,
		ListRule:   &userRequiredRule,
		ViewRule:   &userRequiredRule,
		UpdateRule: &userRequiredRule,
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "player_id",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "points",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "nfl_week",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "rostered_by",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
		)}
}

func GetPlayersCollection() *models.Collection {
	var userRequiredRule = ""
	return &models.Collection{
		Name:       "players",
		System:     false,
		CreateRule: &userRequiredRule,
		ListRule:   &userRequiredRule,
		ViewRule:   &userRequiredRule,
		UpdateRule: &userRequiredRule,
		DeleteRule: nil, // only admins will be able to access
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "hashtag",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "depth_chart_position",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "status",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "sport",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "fantasy_positions",
				Type:     schema.FieldTypeJson,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.JsonOptions{},
			},
			&schema.SchemaField{
				Name:     "number",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "search_last_name",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "injury_start_date",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "weight",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "position",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "practice_participation",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "sportradar_id",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "team",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "last_name",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "college",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "fantasy_data_id",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "injury_status",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "player_id",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "height",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "search_full_name",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "age",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "stats_id",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "birth_country",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "espn_id",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "search_rank",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "first_name",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "depth_chart_order",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "years_exp",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "rotowire_id",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "rotoworld_id",
				Type:     schema.FieldTypeNumber,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.NumberOptions{},
			},
			&schema.SchemaField{
				Name:     "search_first_name",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
			&schema.SchemaField{
				Name:     "yahoo_id",
				Type:     schema.FieldTypeText,
				Unique:   false,
				Required: false,
				System:   false,
				Options:  &schema.TextOptions{},
			},
		)}
}

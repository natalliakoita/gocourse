package datasource

import (
	"context"
	"errors"
	"log"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	database = "tasker"
	collGr   = "groups"
	collCont = "contacts"
)

// create new connection in database
func NewDB(ctx context.Context) *mongo.Client {
	conn_url := "mongodb://root:rootpassword@localhost:27017/?authSource=admin&readPreference=primary&appname=mongodb-vscode%200.6.10&ssl=false"

	_ = "mongodb+srv://root:rootpassword@localhost:27017/corp?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(conn_url))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to mongo is ready")

	return client
}

type DS struct {
	DB *mongo.Client
}

func NewDS(conn *mongo.Client) DS {
	return DS{
		DB: conn,
	}
}

// add contacts in database
func (ds *DS) AddContact(ctx context.Context, name, number string) error {
	collection := ds.DB.Database(database).Collection(collCont)
	t := Contact{}
	if err := collection.FindOne(ctx, bson.M{"name": name, "number": number}).Decode(&t); err == nil {
		err1 := errors.New("Duplicates forbidden")
		return err1
	}
	c := Contact {Name: name, Number: number}
	_, err := collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

// add groups in database
func (ds *DS) AddGroups(ctx context.Context) error {
	collection := ds.DB.Database(database).Collection(collGr)
	var g Group
	g.Name = "Sport"
	_, err := collection.InsertOne(ctx, g)
	if err != nil {
		return err
	}
	g.Name = "Cook"
	_, err = collection.InsertOne(ctx, g)
	if err != nil {
		return err
	}
	return nil
}

// assign contact to group
func (ds *DS) AddGroupToContact(ctx context.Context, name, number string, group_name string) error {
	g := Group{}
	if err := ds.DB.Database(database).Collection(collGr).FindOne(ctx, bson.M{"name": group_name}).Decode(&g); err != nil {
		return err
	}
	filter := bson.M{"number": number, "name": name}

	update := bson.D{
		{"$set", bson.D{{"group_id", g.ID}}},
	}

	_, err := ds.DB.Database(database).Collection(collCont).UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// get contact by group
func (ds *DS) GetContactsOrderByGroup(ctx context.Context) ([]*ContactGroup, error) {
	lookupStage := bson.D{{"$lookup", bson.D{{"from", collGr}, {"localField", "group_id"}, {"foreignField", "_id"}, {"as", "group_id"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$group_id"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedCursor, err := ds.DB.Database(database).Collection(collCont).Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
	if err != nil {
		return nil, err
	}

	var contacts []*ContactGroup
	if err = showLoadedCursor.All(ctx, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

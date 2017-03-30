package data

import (
	"log"

	"github.com/snap10/resource-scheduler/app/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResourceCollection struct {
	C *mgo.Collection
}

func (coll *ResourceCollection) FindAllResources() (resources *[]models.Resource, err error) {

	resourcesSlice := make([]models.Resource, 0, 10)
	err = coll.C.Find(bson.M{}).All(&resourcesSlice)
	if err != nil {
		return nil, err
	}
	return &resourcesSlice, nil
}

func (coll *ResourceCollection) CreateResource(resource *models.Resource) error {
	objID := bson.NewObjectId()
	resource.Id = objID

	err := coll.C.Insert(&resource)
	return err
}

func (coll *ResourceCollection) UpdateResource(resource *models.Resource, key string, data interface{}, mgoOperator string) (updatedResource *models.Resource, err error) {
	log.Printf("data.[UpdateResource]- key %s, data %s", key, data)
	change := mgo.Change{
		Update:    bson.M{mgoOperator: bson.M{key: data}},
		ReturnNew: true,
	}
	info, err := coll.C.FindId(resource.Id).Apply(change, &updatedResource)
	log.Print("info: %s", info.Updated)
	log.Print("resource users: %s", updatedResource.Users)
	return resource, err
}

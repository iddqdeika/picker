package collectionDictionaryMap

import (
	"fmt"
	"picker/definition"
)

func NewStub(m map[string]string) definition.CollectionStorage{
	if m == nil{
		m = make(map[string]string)
	}
	return &collectionDictionaryMapStub{
		m: m,
	}
}

type collectionDictionaryMapStub struct {
	m		map[string]string
}

func (c *collectionDictionaryMapStub) SetDictionaryToCollection(collectionID string, dictionaryID string) error {
	c.m[collectionID] = dictionaryID
	return nil
}

func (c *collectionDictionaryMapStub) GetCollectionDictionary(collectionID string) (string, error) {
	res, ok := c.m[collectionID]
	if !ok{
		return "", fmt.Errorf("cant find collection")
	}
	return res, nil
}


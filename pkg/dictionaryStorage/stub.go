package dictionaryStorage

import (
	"fmt"
	"picker/definition"
	"strconv"
	"sync"
)

func NewStub() definition.DictionaryStorage{
	return &dictionaryStorageStub{
		dicts:    make(map[string]*definition.Dictionary),
		criteria: make(map[string]*definition.Criteria),
		RWMutex:  sync.RWMutex{},
	}
}

type dictionaryStorageStub struct {
	dicts		map[string]*definition.Dictionary
	criteria	map[string]*definition.Criteria
	sync.RWMutex
}

func (ds *dictionaryStorageStub) CreateDictionary(name string) (*definition.Dictionary, error) {
	ds.Lock()
	defer ds.Unlock()
	d := &definition.Dictionary{
		ID:   strconv.Itoa(len(ds.dicts)),
		Name: name,
	}
	ds.dicts[d.ID] = d
	ds.criteria[d.ID] = nil
	return d, nil
}

func (ds *dictionaryStorageStub) GetCriteria(dictID string) (*definition.Criteria, error) {
	if dictID == ""{
		return nil, fmt.Errorf("given dictID is empty")
	}
	ds.RLock()
	defer ds.RUnlock()

	c, ok := ds.criteria[dictID]
	if !ok{
		return nil, fmt.Errorf("dict %v was not defined", dictID)
	}
	return c, nil
}

func (ds *dictionaryStorageStub) SetCriteria(dictID string, criteria definition.Criteria) error {
	ds.Lock()
	defer ds.Unlock()
	ds.criteria[dictID] = &criteria
}

func (ds *dictionaryStorageStub) ClarifyCriteria(c definition.Criteria, dictID string) error {
	panic("implement me")
}


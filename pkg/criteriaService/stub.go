package criteriaService

import (
	"fmt"
	"picker/definition"
	"sync"
)

func NewStub() definition.Dictionary {
	return &dictionary{
		criteriaList:  nil,
		ds:            nil,
		cs:            nil,
		domainHashMap: make(map[definition.Hash]*definition.Domain),
		RWMutex:       sync.RWMutex{},
	}
}

type dictionary struct {
	criteriaList []definition.CriteriaMetadata

	hg definition.HashGenerator
	ds definition.DomainService
	cs definition.CollectionStorage

	domainHashMap map[definition.Hash]*definition.Domain
	sync.RWMutex
}

func (dict *dictionary) GetItemsForSelectedCriteria(criteriaList []definition.SelectedCriteria) ([]definition.Item, error) {
	dict.RLock()
	defer dict.RUnlock()
	criteriaValues := convertSelectedCriteriaToMap(criteriaList)
	// проверяем критерии
	err := dict.checkCriteriaValuesMap(criteriaValues)
	if err != nil {
		return nil, err
	}
	// генерим хеш
	h := dict.hg.GenerateHash(criteriaValues, dict.criteriaList)
	return dict.ds.GetItemsByHash(h)
}

func convertSelectedCriteriaToMap(criteriaList []definition.SelectedCriteria) map[string]string {
	m := make(map[string]string)
	for _, criteria := range criteriaList {
		m[criteria.Name] = criteria.SelectedValue
	}
	return m
}

func (dict *dictionary) checkCriteriaValuesMap(criteriaValues map[string]string) error {
	if len(criteriaValues) != len(dict.criteriaList) {
		return fmt.Errorf("got %v item criteria, need %v", len(criteriaValues), len(dict.criteriaList))
	}
	for _, c := range dict.criteriaList {
		_, ok := criteriaValues[c.Name]
		if !ok {
			return fmt.Errorf("criteria %v not defined", c)
		}
	}
	return nil
}

func (dict *dictionary) SetDictionaryCriteria(names []definition.CriteriaMetadata) error {
	dict.Lock()
	defer dict.Unlock()

	dict.criteriaList = names
	return nil
}

func (dict *dictionary) SetItemCriteria(itemID string, criteriaValues map[string]string) error {
	dict.Lock()
	defer dict.Unlock()
	// проверяем критерии
	err := dict.checkCriteriaValuesMap(criteriaValues)
	if err != nil {
		return err
	}
	// генерим хеш
	h := dict.hg.GenerateHash(criteriaValues, dict.criteriaList)
	// закрепляем за хешем домен позиции
	err = dict.ds.SetHashToItemDomain(h, itemID)
	if err != nil {
		return fmt.Errorf("cant set hash to item domain: %v", err)
	}
	//d, err := dict.ds.GetItemDomain(itemID)
	//if err != nil{
	//	return fmt.Errorf("cant get domain for item %v, err: %v", itemID, err)
	//}
	//dict.domainHashMap[h] = d
	return nil
}

func (dict *dictionary) GetCriteriaMetadataList(colectionID string) ([]definition.CriteriaMetadata, error) {
	dict.RLock()
	defer dict.RUnlock()
	return dict.criteriaList, nil
}

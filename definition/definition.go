package definition

type Item struct {
	ID string
}

type Domain struct {
	ItemList []Item
}

type Hash string

type CriteriaMetadata struct {
	Name  string
	Order int
}

type SelectedCriteria struct {
	Name          string
	SelectedValue string
}

type CriteriaForSelection struct {
	Name              string
	Order             int
	SelectionPossible string
	PossibleValues    []string
}

type CollectionStorage interface {
	SetDictionaryToCollection(collectionID string, dictionaryID string) error
	GetCollectionDictionary(collectionID string) (string, error)
}

// хранит домены в конткесте одного справочника
type DomainService interface {
	// вернет домен позиции. если позиция еще не состоит в домене - создаст его, добавить туда позицию и вернет
	GetItemDomain(itemID string) (*Domain, error)
	MergeItemsDomains(itemID1 string, itemdID2 string) error
	SetHashToItemDomain(h Hash, itemID string) error
	GetItemsByHash(h Hash) ([]Item, error)
}

type Dictionary interface {
	SetDictionaryCriteria(names []CriteriaMetadata) error
	SetItemCriteria(itemID string, criteriaValues map[string]string) error

	GetCriteriaMetadataList(colectionID string) ([]CriteriaMetadata, error)
	GetItemsForSelectedCriteria(criteriaList []SelectedCriteria) ([]Item, error)
}

type HashGenerator interface {
	GenerateHash(criteriaValues map[string]string, md []CriteriaMetadata) Hash
}

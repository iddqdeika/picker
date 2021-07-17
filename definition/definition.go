package definition

type Collection struct {
	ID			string
	Name		string
}

type Item struct {
	ID			string
	Name		string
}

type Dictionary struct {
	ID			string
	Name		string
}

type Criteria struct {
	// имя критерия
	Name				string		`json:"name"`
	// возможные значения. если пусто - пользователю не дадут выбор
	Values				[]string	`json:"values"`
	// выбранное значение. если не пусто - можно выбирать следующие
	SelectedValue		string		`json:"selected_value"`
	// т.к. выбирать надо по-очереди, то каждый критерий ссылается на следующий
	NextCriteria		*Criteria	`json:"next_criteria"`
}

type CollectionStorage interface {
	SetDictionaryToCollection(collectionID string, dictionaryID string) error
	GetCollectionDictionary(collectionID string) (string, error)
}

type DictionaryStorage interface {
	CreateDictionary(name string) (*Dictionary, error)
	GetCriteria(dictID string) (*Criteria, error)
	SetCriteria(dictID string, criteria Criteria) error
	ClarifyCriteria(c Criteria, dictID string) error
}

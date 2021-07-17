package root

import (
	"context"
	"github.com/iddqdeika/rrr"
	"github.com/iddqdeika/rrr/helpful"
	"picker/definition"
	"picker/pkg/collectionDictionaryMap"
)

func New() rrr.Root{
	return &root{}
}

type root struct {
	l		helpful.Logger
	cs		definition.CollectionStorage
}

func (r *root) Register() []error {
	r.l = helpful.DefaultLogger
	r.l.Infof("register started")
	defer r.l.Infof("register ended")

	//errs := make([]error, 0)
	//e := func(err error){
	//	if err != nil{
	//		errs = append(errs, err)
	//	}
	//}

	r.cs = collectionDictionaryMap.NewStub(map[string]string{"1":"dict1"})


	return nil
}

func (r *root) Resolve(ctx context.Context) error {
	r.l.Infof("resolve started")
	defer r.l.Infof("resolve ended")
	
	return nil
}

func (r *root) Release() error {
	return nil
}



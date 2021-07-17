package hashGenerator

import "picker/definition"

func NewStub() definition.HashGenerator {
	return concatHashGenerator{}
}

type concatHashGenerator struct {
}

func (c concatHashGenerator) GenerateHash(criteriaValues map[string]string, md []definition.CriteriaMetadata) definition.Hash {
	var h string
	for name, value := range criteriaValues {
		h += name + value
	}
	return definition.Hash(h)
}

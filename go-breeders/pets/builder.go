package pets

import "errors"

// IPet definds the methods that we want our builder to have.
// These are used to set the fields in the Pet type, and to build the final product.
// Everything except the Build() function returns the type *Pet because we are going to implement the fluent interface
type IPet interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(w int) *Pet
	SetMaxWeight(w int) *Pet
	SetWeight(w int) *Pet
	SetDescription(d string) *Pet
	SetLifespan(lp int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(c string) *Pet
	SetAge(a int) *Pet
	SetAgeEstimated(a bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() IPet {
    return &Pet{}
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}

func (p *Pet) SetMinWeight(w int) *Pet {
	p.MinWeight = w
	return p
}

func (p *Pet) SetMaxWeight(w int) *Pet {
	p.MaxWeight = w
	return p
}

func (p *Pet) SetWeight(w int) *Pet {
	p.Weight = w
	return p
}

func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}

func (p *Pet) SetLifespan(lp int) *Pet {
	p.Lifespan = lp
	return p
}

func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}

func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}

func (p *Pet) SetAge(a int) *Pet {
	p.Age = a
	return p
}

func (p *Pet) SetAgeEstimated(a bool) *Pet {
	p.AgeEstimated = a
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}

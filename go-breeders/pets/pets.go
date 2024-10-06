package pets

// omitempty: if don't have it, leave it out of JSON
type Pet struct {
	Species          string `json:"species"`
	Breed            string `json:"breed"`
	MinWeight        int    `json:"min_weight,omitempty"`
	MaxWeight        int    `json:"max_weight,omitempty"`
	AverageWeight    int    `json:"average_weight,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	Description      string `json:"description,omitempty"`
	Lifespan         int    `json:"lifespan,omitempty"`
	GeographicOrigin string `json:"geographic_orign,omitempty"`
	Color            string `json:"color,omitempty"`
	Age              int    `json:"age,omitempty"`
	AgeEstimated     bool    `json:"age_estimated,omitempty"`
}

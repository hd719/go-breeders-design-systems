package pets

type Pet struct {
	Species          string `json:"species"`
	Breed            string `json:"breed"`
	MinWeight        int    `json:"min_weight,omitempty"`
	MaxWeight        int    `json:"max_weight,omitempty"`
	AverageWeight    int    `json:"average_weight,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	Description      string `json:"description,omitempty"`
	LifeSpan         int    `json:"lifespan,omitempty"`
	GeographicOrigin int    `json:"geographic_origin,omitempty"`
	Color            string `json:"color,omitempty"`
	Age              string `json:"age,omitempty"`
	AgeEstimated     string `json:"age_estimated,omitempty"`
}

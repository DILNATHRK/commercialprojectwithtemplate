package models

// use a single instance of Validate, it caches struct info
// var validate *validator.Validate

type Building struct {
	// Name             string `bson:"name" validate:"required,regexp=^[a-zA-Z]$"`
	Name             string `bson:"name" validate:"required"`
	Location         string `bson:"location" validate:"required"`
	AvailabilityFor  string `bson:"availability_for" validate:"required"`
	CompletionStatus string `bson:"completion_status" validate:"required"`
	FurnishingStatus string `bson:"furnishing_status" validate:"required"`
	Floors           string `bson:"no_of_floor" validate:"required"`
	Parking          string `bson:"parking" validate:"required"`
	Lift             string `bson:"lift" validate:"required"`
	Oc               string `bson:"oc" validate:"required"`
	OverLooking      string `bson:"overlooking" validate:"required"`
}

type Destination struct {
	CountryName     string `bson:"country_name" validate:"required"`
	CurrencyName    string `bson:"currency_name" validate:"required"`
	CurrencySymmbol string `bson:"currency_symbol" validate:"required"`
}

type State struct {
	StateName     string `bson:"state_name" validate:"required"`
	StateLanguage string `bson:"state_languague" validate:"required"`
}

type City struct {
	CityName string `bson:"city_name" validate:"required"`
}

type Locality struct {
	LocalityName string `bson:"locality_name" validate:"required"`
}

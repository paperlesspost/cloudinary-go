package config

// Cloud defines the account configuration required to connect your application to Cloudinary.
//
// https://cloudinary.com/documentation/how_to_integrate_cloudinary#get_familiar_with_the_cloudinary_console
type Cloud struct {
	CloudName string
	ApiKey    string
	ApiSecret string
}
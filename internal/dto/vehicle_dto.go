package dto

type InputCreateVehicleDTO struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

type OutputCreateVehicleDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

type InputUpdateVehicleDTO struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

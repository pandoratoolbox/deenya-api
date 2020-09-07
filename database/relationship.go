package database

import (
	"deenya-api/models"
	"fmt"
)

func MyClients(mid int64) []models.User {
	var data []models.User
	q := `SELECT * FROM public.user WHERE id IN (SELECT id FROM public.consultant_client WHERE consultant_id = $1)`

	err := db.Get(&data, q, mid)

	if err != nil {
		fmt.Println(err)
		return data
	}

	return data
}

func MyConsultants(mid int64) []models.User {
	var data []models.User
	q := `SELECT * FROM public.consultant_client WHERE client_id = $1`

	err := db.Get(&data, q, mid)

	if err != nil {
		fmt.Println(err)
		return data
	}
	//q := `SELECT * FROM public.booking WHERE client_id = $1`
	//q := `SELECT * FROM public.order WHERE client_id = $1`

	return data
}

func NewConsultantClient() {

}

//order
//booking

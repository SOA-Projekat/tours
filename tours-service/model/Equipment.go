package model

/*import(
	"database/sql/driver"
	"encoding/json"
	"errors"
)
*/

type Equipment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TourID      int    `json:"tourID"`
}

package domain

type Item struct {
	ID          int64  `json:"-" db:"id"`
	OrderUID    string `json:"-" db:"order_uid"`
	ChrtID      int64  `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int64  `json:"price" db:"price"`
	RID         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int64  `json:"total_price" db:"total_price"`
	NmID        int64  `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
}

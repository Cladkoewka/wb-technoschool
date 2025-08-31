package entity

type EventDTO struct {
    UserID int    `json:"user_id"`
    Date   string `json:"date"`  
    Text   string `json:"event"`
}
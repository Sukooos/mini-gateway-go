package main

type Billing struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Amount      float64 `json:"amount"`       // nominal tagihan
	Description string  `json:"description"`  // keterangan/tagihan apa
	Status      string  `json:"status"`       // "pending", "paid", dll
	CreatedAt   string  `json:"created_at"`   // waktu dibuat
	DueDate     string  `json:"due_date"`     // (opsional) jatuh tempo
}
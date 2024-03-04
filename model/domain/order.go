package domain

// Order mendefinisikan sebuah pesanan
type Order struct {
	ID         int    // ID unik untuk pesanan
	UserID     int    // ID pengguna yang membuat pesanan
	CartId     []int  // ID keranjang yang berisi item pesanan
	OrderItems []Cart // Daftar item dalam pesanan
	TotalItems int    // Jumlah total item dalam pesanan
	TotalPrice int    // Harga total dari semua item dalam pesanan
	// ShippingCost  float64       // Biaya pengiriman pesanan
	// GrandTotal    float64       // Total keseluruhan pesanan setelah ditambah biaya pengiriman
	OrderStatus   OrderStatus   // Status pesanan (misalnya: pending, diproses, selesai, dll.)
	PaymentStatus PaymentStatus // Status pembayaran pesanan (misalnya: pending, berhasil, gagal, dll.)
}

// OrderStatus mendefinisikan status pesanan
type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Delivered  OrderStatus = "delivered"
	Cancelled  OrderStatus = "cancelled"
)

// PaymentStatus mendefinisikan status pembayaran pesanan
type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"
	PaymentSuccess PaymentStatus = "success"
	PaymentFailed  PaymentStatus = "failed"
)

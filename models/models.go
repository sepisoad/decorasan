package models

import "time"

//******************************************************************************

// // Media type.--dont need it
// type Media struct {
// 	Title string `json:"title"`
// 	Type  string `json:"type"` // gif, png, mp4, ...
// 	Path  string `json:"path"`
// }

// Texture type.
type Texture struct {
	Texture   string `json:"texture"`
	Thumbnail string `json:"thumbnail"`
}

// Model3D type.
type Model3D struct {
	Object   string    `json:"object"`
	Textures []Texture `json:"textures"`
}

// Address type.
type Address struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	Value       string `json:"value"`
	PhoneNumber string `json:"phone_number"`
	PostCode    string `json:"post_code"`
	IsDefault   bool   `json:"is_default"`
}

//BankAccount type // DO WE REALLY NEED IT?
type BankAccount struct {
	Bank          string `json:"bank"`
	OwnerName     string `json:"owner_name"`
	AccountNumber string `json:"account_number"`
}

//Discount model //TODO: ????
type Discount struct {
	ID         int       `sql:",AUTO_INCREMENT" json:"id"`
	Percentage int       `json:"percentage"`
	ValidTime  time.Time `json:"valid_time"`
}

//******************************************************************************

//SMS model.
type SMS struct {
	ID          int64     `sql:",AUTO_INCREMENT" json:"id"`
	Code        int       `json:"code"`
	UUID        string    `json:"uuid"`
	PhoneNumber string    `json:"phone_number"`
	RequestDate time.Time `json:"request_date"`
	ExpiryDate  time.Time `json:"expiry_date"`
}

//******************************************************************************

// AppInfo is used to keep the latest version of different components of the
// system in database such as app version or backend version.AppInfo.
type AppInfo struct {
	ID          int64  `sql:",AUTO_INCREMENT" json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	IsSupported bool   `json:"is_supported"`
}

//******************************************************************************

// UserType enum type.TODO: sepehr will check if is this needed or not
type UserType string

const (
	// UserTypeBusinessOwner enum value.
	UserTypeBusinessOwner UserType = "business-owner"
	// UserTypeAdmin enum value.
	UserTypeAdmin UserType = "admin"
	// UserTypeShopOwner enum value.
	UserTypeShopOwner UserType = "shop-owner"
	// UserTypeAppUser enum value.
	UserTypeAppUser UserType = "app-user"
)

//AppUser model
type AppUser struct {
	ID                int       `sql:",AUTO_INCREMENT" json:"id"`
	PhoneNumber       string    `json:"phone_number"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Gender            string    `json:"gender"`
	Birthday          string    `json:"birthday"`
	Email             string    `json:"email"`
	Address           []Address `json:"address"`
	FavouriteProducts []int     `pg:",array" json:"favourite_products"`
	IsComplete        bool      `json:"is_complete"`
	CreatedAt         time.Time `json:"create_at"`
	UpdatedAt         time.Time `json:"update_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

//ShopUser model
type ShopUser struct {
	ID          int       `sql:",AUTO_INCREMENT" json:"id"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	Birthday    string    `json:"birthday"`
	Email       string    `json:"email"`
	Address     []Address `json:"address"`
	IsComplete  bool      `json:"is_complete"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

//SuperUser model
type SuperUser struct {
	ID          int       `sql:",AUTO_INCREMENT" json:"id"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	Birthday    string    `json:"birthday"`
	Email       string    `json:"email"`
	Address     []Address `json:"address"`
	IsComplete  bool      `json:"is_complete"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

//DebitcardUser model.
type DebitcardUser struct {
	ID          int       `sql:",AUTO_INCREMENT" json:"id"`
	Number      int       `json:"number"`
	ShopID      int       `json:"shop_id"`
	UserID      int       `json:"user_id"`
	Credit      int       `json:"credit"`
	Active      bool      `json:"active"`
	AssighnDate time.Time `json:"assighn_date"`
}

// AdminPanelUser model.
type AdminPanelUser struct {
	ID       int    `sql:",AUTO_INCREMENT" json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
}

//Mall model
type Mall struct {
	ID         int    `sql:",AUTO_INCREMENT" json:"id"`
	Name       string `json:"name"`
	FloorCount int    `json:"floor_count"`
	UnitCount  int    `json:"unit_count"`
}

// //ShopsDetails model--dont need it
// type ShopsDetails struct {
// 	ID              int
// 	MallID          int
// 	AppAddress      int
// 	ReadableAddress string
// }

// MetaProduct model.
type MetaProduct struct {
	ID        int `sql:",AUTO_INCREMENT" json:"id"`
	ProductID int `json:"product_id"`
	Revision  int `json:"revision"`
}

// Product model.
type Product struct {
	ID                int         `sql:",AUTO_INCREMENT" json:"id"`
	ShopID            int         `json:"shop_id"`
	Name              string      `json:"name"`
	Brand             string      `json:"brand"`
	Code              int         `json:"code"`
	Price             float64     `json:"price"`
	DeliveryTime      int         `json:"delivery_time"`
	ManufactureSerial string      `json:"manufacture_serial"` // <== WTF
	CategoryID        int         `json:"category_id"`
	Discount          float64     `json:"discount"`
	Thumbnail         string      `json:"thumbnail"`
	Gallery           []string    `json:"gallery"`
	Model3D           Model3D     `json:"model3d" pg:",bson"`
	DeliveryPrice     float64     `json:"delivery_price"`
	Guarantee         string      `json:"guarantee"`
	Description       string      `json:"description"`
	Attributes        interface{} `json:"attributes"`
	CreatedAt         time.Time   `json:"create_at"`
	UpdatedAt         time.Time   `json:"update_at"`
	DeletedAt         time.Time   `json:"deleted_at"`
}

//ProductAttribute model
type ProductAttribute struct {
	ID       int    `sql:",AUTO_INCREMENT" json:"id"`
	Name     string `json:"name"`
	DataType string `json:"date_type"`
	IsArray  bool   `json:"is_array"`
}

//******************************************************************************

// MetaShop model.
type MetaShop struct {
	ID       int `sql:",AUTO_INCREMENT" json:"id"`
	ShopID   int `json:"shop_id"`
	Revision int `json:"revision"`
}

//Shop model
type Shop struct {
	ID           int       `sql:",AUTO_INCREMENT" json:"id"`
	MallID       int       `json:"mall_id"`
	ContractID   int       `json:"contract_id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	Owner        string    `json:"owner"`
	Keeper       string    `json:"keeper"`
	VirAddress   string    `json:"vir_address"` // TODO: format f01-n11 floor 3 number 11
	Adress       Address   `json:"address"`
	PhoneNumbers []string  `pg:",array" json:"phone_number"`
	Thumbnail    string    `json:"thumbnail"`     // TODO: create from on file chossen by user
	MediaGallery []string  `json:"media_gallery"` // TODO: hard disk format (by name or by directory hierarchical)
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

//Category model
type Category struct {
	ID        int    `sql:",AUTO_INCREMENT" json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

//******************************************************************************

// Contract model.
type Contract struct {
	ID     int `sql:",AUTO_INCREMENT" json:"id"`
	Number int `json:"number"`
	// ContractScan   []Media   `json:"contract_scan"`
	Scans          []string  `json:"scans"`
	Type           string    `json:"type"` //pysical or virtual
	CompanyName    string    `json:"company_name"`
	Date           time.Time `json:"date"`
	BankAccounts   []string  `json:"bank_accounts"`
	Duration       int       `json:"duration"`
	MontlytDueDate time.Time `json:"montly_due_date"` //WTF???
	CreatedAt      time.Time `json:"create_at"`
	UpdatedAt      time.Time `json:"update_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

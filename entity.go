package dedge

import "encoding/xml"

//Rooms ...
type Rooms struct {
	XMLName    xml.Name
	HotelId    string `xml:"hotelId,attr"`
	HotelName  string `xml:"hotelName,attr,omitempty"`
	PriceModel string `xml:"priceModel,attr,omitempty"`
	Room       []Room `xml:"room"`
}

//Room ...
type Room struct {
	XMLName      xml.Name
	Code         string       `xml:"code,attr"`
	Name         string       `xml:"name,attr"`
	MinOccupancy int          `xml:"minOccupancy,attr"`
	MaxOccupancy int          `xml:"maxOccupancy,attr"`
	Rate         []Rate       `xml:"rate"`
	Occupancies  *Occupancies `xml:"occupancies"`
}

type Rate struct {
	XMLName xml.Name
	Code    string `xml:"code,attr"`
	Name    string `xml:"name,attr"`
	Regime  string `xml:"regime,attr"`
}

type Occupancies struct {
	XMLName   xml.Name
	Occupancy []Occupancy `xml:"occupancy,omitempty"`
}

type Occupancy struct {
	XMLName    xml.Name
	AdultCount int `xml:"adultCount,attr"`
	ChildCount int `xml:"childCount,attr"`
}

type Authentication struct {
	XMLName  xml.Name `xml:"authentication"`
	Login    string   `xml:"login,attr" validate:"required"`
	Password string   `xml:"password,attr" validate:"required"`
}

type InventoryUpdate struct {
	XMLName    xml.Name
	HotelId    string       `xml:"hotelId,attr" validate:"required"`
	RoomUpdate []RoomUpdate `xml:"room" validate:"dive"`
}

type RoomUpdate struct {
	XMLName    xml.Name
	Id         int `xml:"id,attr" validate:"required"`
	Inventory  Inventory
	RateUpdate []RateUpdate `xml:"rate" validate:"dive"`
}

type Inventory struct {
	XMLName      xml.Name       `xml:"inventory"`
	Availability []Availability `xml:"availability" validate:"dive"`
}

type Availability struct {
	XMLName  xml.Name
	From     string `xml:"from,attr" validate:"required"`
	To       string `xml:"to,attr" validate:"required"`
	Quantity int    `xml:"quantity,attr" validate:"required"`
}

type RateUpdate struct {
	XMLName      xml.Name
	Currency     string         `xml:"currency,attr"`
	RateCode     string         `xml:"rateCode,attr" validate:"required,max=50"`
	RateName     string         `xml:"rateName,attr"`
	PlanningRate []PlanningRate `xml:"planning" validate:"dive"`
}

type PlanningRate struct {
	XMLName     xml.Name
	From        string `xml:"from,attr" validate:"required"`
	To          string `xml:"to,attr" validate:"required"`
	MinimumStay int    `xml:"minimumStay,attr"`
	MaximumStay int    `xml:"maximumStay,attr"`
	UnitPrice   string `xml:"unitPrice,attr"`
	NoArrival   bool   `xml:"noArrival,attr"`
	NoDeparture bool   `xml:"noDeparture,attr"`
	IsClosed    bool   `xml:"isClosed,attr"`
}

type RoomAllotment struct {
	RoomId    int
	StartDate string
	EndDate   string
	Quantity  int
}

type RoomRate struct {
	RoomId   int
	Currency string
	RateCode string
	RateName string
	Planning []PlanningRate
}

type RoomPlanning struct {
	RoomId      int
	Currency    string
	RateCode    string
	RateName    string
	StartDate   string
	EndDate     string
	MinimumStay int
	MaximumStay int
	Price       string
	NoArrival   bool
	NoDeparture bool
	IsClosed    bool
}

type AvailabilityUpdate struct {
	RoomId      int
	FromDate    string
	EndDate     string
	Quantity    int
	Currency    string
	RateCode    int
	RateName    string
	MinimumStay int
	MaximumStay int
	UnitPrice   string
	NoArrival   int
	NoDeparture int
	IsClosed    int
}

type Bookings struct {
	XMLName xml.Name `xml:"bookings"`
	HotelId int      `xml:"hotelId,attr"`
	Booking []Booking
}

type Booking struct {
	XMLName       xml.Name `xml:"booking"`
	Id            int      `xml:"id,attr"`
	Action        string   `xml:"action,attr"`
	Currency      string   `xml:"currency,attr"`
	Date          *string  `xml:"date,attr"`
	TotalAmount   int      `xml:"totalAmount,attr"`
	PaidAmount    int      `xml:"paidAmount,attr"`
	DueAmount     int      `xml:"dueAmount,attr"`
	PayableAmount float64  `xml:"payableAmount,attr"`
	Origin        string   `xml:"origin,attr,omitempty"`
	PaxCount      int      `xml:"paxCount,attr"`
	AdultCount    int      `xml:"adultCount,attr,omitempty"`
	ChildCount    int      `xml:"childCount,attr,omitempty"`
	InfantCount   int      `xml:"infantCount,attr,omitempty"`
	Customer      Customer
	Rooms         []RoomBooking `xml:"rooms"`
}

type Customer struct {
	XMLName   xml.Name `xml:"customer"`
	FirstName string   `xml:"firstName,attr"`
	LastName  *string  `xml:"lastName,attr"`
	Title     string   `xml:"title,attr,omitempty"`
	Contact   Contact
	Comment   string `xml:"comment,omitempty"`
}

type Contact struct {
	XMLName xml.Name `xml:"contact"`
	Email   *string  `xml:"email,attr"`
	Phone   *string  `xml:"phone,attr,omitempty"`
}

type ContactAddress struct {
	XMLName    xml.Name
	City       string `xml:"city,attr,omitempty"`
	Country    string `xml:"country,attr,omitempty"`
	PostalCode string `xml:"postalCode,attr,omitempty"`
}

type CreditCard struct {
	XMLName    string `xml:"creditCard,omitempty"`
	Number     string `xml:"number,attr,omitempty"`
	CardHolder string `xml:"cardHolder,attr,omitempty"`
	ExpiryDate string `xml:"expiryDate,attr,omitempty"`
	CardType   string `xml:"cardType,attr,omitempty"`
}

type Distributor struct {
	XMLName xml.Name  `xml:"distributor,omitempty"`
	Comment []Comment `xml:"comment,omitempty"`
}

type Comment struct {
	XMLName xml.Name `xml:"comment"`
	Name    string   `xml:"name,attr"`
}

type RoomBooking struct {
	XMLName xml.Name `xml:"room"`
	Id      string   `xml:"id,attr"`
	Stays   []Stay   `xml:"stays"`
	Guests  []Guest  `xml:"guests"`
}

type Stay struct {
	XMLName   xml.Name `xml:"stay"`
	Date      string   `xml:"date,attr"`
	Quantity  int      `xml:"quantity,attr"`
	UnitPrice float64  `xml:"unitPrice,attr"`
	RateCode  string   `xml:"rateCode,attr"`
}

type Guest struct {
	XMLName   xml.Name `xml:"guest"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
	Title     string   `xml:"title,attr,omitempty"`
	AgeRange  int      `xml:"ageRange,attr,omitempty"`
	BirthDate string   `xml:"birthDate,attr,omitempty"`
}

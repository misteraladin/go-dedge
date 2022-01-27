package dedge

const (
	HotelIdParam        string = "hotelCode" //HotelIdParam
	UserLoginParam      string = "login"
	UserPasswordParam   string = "password"
	FromDateParam       string = "from"
	ToDateParam         string = "to"
	DurationParam       string = "duration"
	JasmineApiKeyHeader string = "X-Jasmine-Api-Key"

	FeedInDEdgeUrlPath string = "/feed-in/dedge"

	XmlSchemaInstance string = "http://www.w3.org/2001/XMLSchema-instance"
	XmlSchema         string = "http://www.w3.org/2001/XMLSchema"

	PriceModelPerDay          string = "PerDay"
	PriceModelPerPerson       string = "PerPerson"
	PriceModelPerPerOccupancy string = "PerOccupancy"

	ErrorNoRooms        string = "error no rooms"
	ErrorNoRates        string = "error no rates"
	ErrorHotelNotActive string = "hotel not active"

	ErrorNoHotel string = "error no hotel"

	FailureTypeHotelNotLinked         int = 1
	FailureTypeHotelNotActivated      int = 2
	FailureTypeBadAuthentication      int = 3
	FailureTYpeRoomMappingError       int = 4
	FailureTypeRateMappingError       int = 5
	FailureTypeRoomRateMapping        int = 6
	WarningTypeRateDerivation         int = 7
	WarningTypeRoomRateDerivation     int = 8
	WarningTypePriceTooLow            int = 9
	WarningTypePriceTooHigh           int = 10
	WarningTYpeMissingOccupancy       int = 11
	WarningTypeUnknownOccupancy       int = 12
	FailureTypeIncorrectPriceCurrency int = 13
	FailureTypeNoRooms                int = 14
	FailureTypeNoRates                int = 15
	FailureTypeDomInvalidXML          int = 201
	FailureTypeInvalidXMLMessage      int = 202
	FailureTypeInternalError          int = 203

	TotalHotelPerPage int = 200
)

var ErrorMessageMap = map[int]string{
	FailureTypeHotelNotLinked:         "HotelNotLinked",
	FailureTypeHotelNotActivated:      "HotelNotActivated",
	FailureTypeBadAuthentication:      "BadAuthentication",
	FailureTYpeRoomMappingError:       "RoomMappingError",
	FailureTypeRateMappingError:       "RateMappingError",
	FailureTypeRoomRateMapping:        "RoomRateMapping",
	WarningTypeRateDerivation:         "RateDerivation",
	WarningTypeRoomRateDerivation:     "RoomRateDerivation",
	WarningTypePriceTooLow:            "PriceTooLow",
	WarningTypePriceTooHigh:           "PriceTooHigh",
	WarningTYpeMissingOccupancy:       "MissingOccupancy",
	WarningTypeUnknownOccupancy:       "UnknownOccupancy",
	FailureTypeIncorrectPriceCurrency: "IncorrectPriceCurrency",
	FailureTypeNoRooms:                "NoRooms",
	FailureTypeNoRates:                "NoRates",
	FailureTypeDomInvalidXML:          "DomInvalidXML",
	FailureTypeInvalidXMLMessage:      "InvalidXMLMessage",
	FailureTypeInternalError:          "InternalError",
}

var ErrorCommentMap = map[int]string{
	FailureTypeHotelNotLinked:         "Hotel is not linked to D-EDGE Channel Manager or is unknown",
	FailureTypeHotelNotActivated:      "Hotel is not activated for D-EDGE Channel Manager",
	FailureTypeBadAuthentication:      "Credentials is invalid for this Hotel",
	FailureTYpeRoomMappingError:       "Room is invalid or not configured",
	FailureTypeRateMappingError:       "Rate is invalid or not configured",
	FailureTypeRoomRateMapping:        "The combination of Room and Rate is invalid or not configured",
	WarningTypeRateDerivation:         "Price for Rate cannot be updated because of its configuration",
	WarningTypeRoomRateDerivation:     "Price for Rate and Room cannot be updated because of its configuration",
	WarningTypePriceTooLow:            "Price for Rate and Room cannot be update because it's too low",
	WarningTypePriceTooHigh:           "Price for Rate and Room cannot be update because it's too high",
	WarningTYpeMissingOccupancy:       "Price is missing for following occupancies: X persons, X adults / Y childs, etc.",
	WarningTypeUnknownOccupancy:       "Price is sent for unknown occupancies: X persons, X adults / Y childs, etc",
	FailureTypeIncorrectPriceCurrency: "Price given is not in correct currency",
	FailureTypeNoRooms:                "Hotel does not have rooms created or configured",
	FailureTypeNoRates:                "Hotel does not have rooms created or configured",
	FailureTypeDomInvalidXML:          "DOM validation parsing error",
	FailureTypeInvalidXMLMessage:      "XML parsing error when message doesn't contain some mandatory elements",
}

//provide booking
var ActionMap = map[string]string{
	"pending":  "Create",
	"issued":   "Modify",
	"canceled": "Cancel",
}

var ListMemberTitle = []string{"Tuan", "Nyonya", "Nona", "mr", "mrs", "ms"}

var MemberTitleMap = map[string]string{
	"Tuan":   "Mr",
	"Nyonya": "Mrs",
	"Nona":   "Ms",
	"mr":     "Mr",
	"mrs":    "Mrs",
	"ms":     "Ms",
}

package dedge

import "encoding/xml"

//RetrieveHotelResponse ...
type RetrieveHotelResponse struct {
	XMLName  xml.Name `xml:"message"`
	XmlnsXsi string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd string   `xml:"xmlns:xsd,attr"`
	Rooms    Rooms    `xml:"rooms"`
}

//InventoryUpdateResponse ...
type InventoryUpdateResponse struct {
	XMLName xml.Name `xml:"message"`
	Success string   `xml:"success"`
}

//ProvideBookingResponse ...
type ProvideBookingResponse struct {
	XMLName  xml.Name   `xml:"message"`
	XmlnsXsi string     `xml:"xmlns:xsi,attr"`
	XmlnsXsd string     `xml:"xmlns:xsd,attr"`
	Success  string     `xml:"success"`
	Bookings []Bookings `xml:"bookings"`
}

//FailureMessage ...
type FailureMessage struct {
	XMLName xml.Name `xml:"message"`
	Failure struct {
		Type    int    `xml:"type,attr"`
		Message string `xml:"message,attr"`
		Comment string `xml:"comment"`
	} `xml:"failure"`
}

//WarningMessage ...
type WarningMessage struct {
	XMLName xml.Name `xml:"message"`
	Warning struct {
		Type    int    `xml:"type,attr"`
		Message string `xml:"message,attr"`
		Comment string `xml:"comment"`
	} `xml:"warning"`
}

//BindErrorResponse ...
func BindErrorResponse(errorType int, errorComment string) FailureMessage {
	var errResp FailureMessage
	if errorComment == "" {
		errorComment = ErrorCommentMap[errorType]
	}

	errResp.Failure.Type = errorType
	errResp.Failure.Message = ErrorMessageMap[errorType]
	errResp.Failure.Comment = errorComment

	return errResp
}

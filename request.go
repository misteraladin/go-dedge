package dedge

import "encoding/xml"

//InventoryUpdateRequest ...
type InventoryUpdateRequest struct {
	XMLName         xml.Name        `xml:"message"`
	XmlnsXsi        string          `xml:"xmlns:xsi,attr"`
	XmlnsXsd        string          `xml:"xmlns:xsd,attr"`
	Authentication  Authentication  `xml:"authentication"`
	InventoryUpdate InventoryUpdate `xml:"inventoryUpdate"`
	Raw             string          `xml:",innerxml"`
}

//ProvideBookingRequest ...
type ProvideBookingRequest struct {
	From     string
	To       string
	Duration string
}

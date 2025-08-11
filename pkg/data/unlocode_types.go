package data

type Unlocode struct {
	ChangeIndicator  string
	LoCode           string
	Name             string
	NameWoDiacritics string
	SubDiv           string
	Function         *[]Function
	Status           Status
	Date             string
	IATA             string
	Coordinates      Point
	Remarks          string
}

type Point struct {
	Latitude   float64
	Longtitude float64
}

type Function string

const (
	UNKNOWN                Function = "0" // A value "0" in the first position specifies that the functional use of a location is not known and is to be specified.
	PORT                   Function = "1" // Specifies that the location is a Port, as defined in UN/ECE Recommendation 16.
	RAIL_TERMINAL          Function = "2" // Specifies that the location is a Rail terminal.
	ROAD_TERMINAL          Function = "3" // Specifies that the location is a Road terminal.
	AIRPORT                Function = "4" // Specifies that the location is an Airport.
	POSTAL_EXCHANGE_OFFICE Function = "5" // Specifies that the location is a Postal exchange office.
	MULTIMODAL             Function = "6" // Value reserved for multimodal functions, ICDs etc.
	FIXED_TRANSPORT        Function = "7" // Value reserved for fixed transport functions (e.g. oil platform).
	BORDER_CROSSING        Function = "B" // Specifies that the location is Border crossing.
)

type Status string

const (
	UNSET_STATUS                                     Status = ""
	UNKNOWN_STATUS_AM                                Status = "AM"
	UNKNOWN_STATUS_UR                                Status = "UR"
	UNKNOWN_STATUS_AQ                                 Status = "AQ"
	GOVERNMENT_AGENCY_APPROVED                       Status = "AA" // Approved by competent national government agency
	CUSTOMS_AUTHORITY_APPROVED                       Status = "AC" //	Approved by Customs Authority
	NATIONAL_FACILITATION_BODY_APPROVED              Status = "AF" //	Approved by national facilitation body
	INTERNATIONAL_ORGANISATION_ADOPTED               Status = "AI" //	Code adopted by international organisation (IATA or ECLAC)
	NATIONAL_STANDARDISATION_BODY_APPROVED           Status = "AS" //	Approved by national standardisation body
	RECOGNISED_LOCATION                              Status = "RL" //	Recognised location - Existence and representation of location name confirmed by check against nominated gazetteer or other reference work
	REQUEST_FROM_CREDIBLE_NATION_SOURCE              Status = "RN" // Request from credible national sources for locations in their own country
	REQUEST_UNDER_CONSIDERATION                      Status = "RQ" // Request under consideration
	REQUEST_REJECTED                                 Status = "RR" // Request rejected
	ORIGINAL_ENTRY_NOT_VERIFIED_SINCE_DATE_INDICATED Status = "QQ" // Original entry not verified since date indicated
	ENTRY_WILL_BE_REMOVED                            Status = "XX" // Entry that will be removed from the next issue of UN/LOCODE
)

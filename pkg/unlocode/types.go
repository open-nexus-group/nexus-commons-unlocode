//go:generate ../generator/parse.go -output=../unlocode/data.go SubDiv.csv UNLOC-1.csv UNLOC-2.csv UNLOC-3.csv UNLOC-4.csv 
package unlocode
import ("log")

type Unlocode struct {
	ChangeIndicator  string
	LoCode           string
	Name             string
	NameWoDiacritics string
	SubDiv           string
	Function         []Function
	Status           []Status
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

type SubDivision struct {
	CountryCode string
	Code        string
	Name        string
	Category    Category
}

type SubDivCollection struct {
	subdivisions []SubDivision
}

type Category string

const (
	ADMINISTRATION                                     Category = "Administration"
	ADMINISTRATIVE_ATOLL                               Category = "Administrative Atoll"
	ADMINISTRATIVE_REGION                              Category = "Administrative Region"
	ADMINISTRATIVE_TERRITORY                           Category = "Administrative Territory"
	ARCTIC_REGION                                      Category = "Arctic Region"
	AREA                                               Category = "Area"
	AUTONOMOUS_CITY                                    Category = "Autonomous City"
	AUTONOMOUS_CITY_IN_NORTH_AFRICA                    Category = "Autonomous City In North Africa"
	AUTONOMOUS_DISTRICT                                Category = "Autonomous District"
	AUTONOMOUS_MUNICIPALITY                            Category = "Autonomous Municipality"
	AUTONOMOUS_PROVINCE                                Category = "Autonomous Province"
	AUTONOMOUS_REGION                                  Category = "Autonomous Region"
	AUTONOMOUS_REPUBLIC                                Category = "Autonomous Republic"
	AUTONOMOUS_SECTOR                                  Category = "Autonomous Sector"
	AUTONOMOUS_TERRITORY_UNIT                          Category = "Autonomous Territory Unit"
	BOROUGH                                            Category = "Borough"
	CANTON                                             Category = "Canton"
	CAPITAL                                            Category = "Capital"
	CAPITAL_CITY                                       Category = "Capital City"
	CAPITAL_DISTRICT                                   Category = "Capital District"
	CAPITAL_TERRITORY                                  Category = "Capital Territory"
	CHAINS_OF_ISLANDS                                  Category = "Chains (Of Islands)"
	CITY                                               Category = "City"
	CITY_CORPORATION                                   Category = "City Corporation"
	CITY_MUNICIPALITY                                  Category = "City Municipality"
	CITY_WITH_COUNTY_RIGHTS                            Category = "City With County Rights"
	COMMUNE                                            Category = "Commune"
	COUNCIL_AREA                                       Category = "Council Area"
	COUNTY                                             Category = "County"
	DECENTRALIZED_REGIONAL                             Category = "Decentralized Regional"
	DECENTRALIZED_REGIONAL_ENTITY                      Category = "Decentralized Regional Entity"
	DEPARTMENT                                         Category = "Department"
	DEPARTMENTS                                        Category = "Departments"
	DEPENDENCY                                         Category = "Dependency"
	DISTRICT                                           Category = "District"
	DISTRICT_MUNICIPALITY                              Category = "District Municipality"
	DISTRICT_WITH_SPECIAL_STATUS                       Category = "District With Special Status"
	DISTRICTS_UNDER_REPUBLIC_ADMINISTRATION            Category = "Districts Under Republic Administration"
	DIVISION                                           Category = "Division"
	ECONOMIC_PREFECTURE                                Category = "Economic Prefecture"
	EMIRATE                                            Category = "Emirate"
	ENTITY                                             Category = "Entity"
	EUROPEAN_COLLECTIVITY                              Category = "European Collectivity"
	FEDERAL_CAPITAL_TERRITORY                          Category = "Federal Capital Territory"
	FEDERAL_DEPENDENCIES                               Category = "Federal Dependencies"
	FEDERAL_DISTRICT                                   Category = "Federal District"
	FEDERAL_ENTITY                                     Category = "Federal Entity"
	FEDERAL_TERRITORY                                  Category = "Federal Territory"
	FREE_COMMUNAL_CONSORTIA                            Category = "Free Communal Consortia"
	GEOGRAPHICAL_ENTITY                                Category = "Geographical Entity"
	GOVERNORAT                                         Category = "Governorat"
	GOVERNORATE                                        Category = "Governorate"
	GROUP_OF_ISLANDS                                   Category = "Group Of Islands"
	INDIGENOUS_REGION                                  Category = "Indigenous Region"
	ISLAND                                             Category = "Island"
	ISLAND_COUNCIL                                     Category = "Island Council"
	ISLANDS_GROUPS_OF_ISLANDS                          Category = "Islands/Groups Of Islands"
	LAND                                               Category = "Land"
	LOCAL_COUNCIL                                      Category = "Local Council"
	LONDON_BOROUGH                                     Category = "London Borough"
	METROPOLITAN_ADMINISTRATION                        Category = "Metropolitan Administration"
	METROPOLITAN_CITY                                  Category = "Metropolitan City"
	METROPOLITAN_COLLECTIVITY_WITH_SPECIAL_STATUS      Category = "Metropolitan Collectivity With Special Status"
	METROPOLITAN_DEPARTMENT                            Category = "Metropolitan Department"
	METROPOLITAN_REGION                                Category = "Metropolitan Region"
	MUNICIPALITY                                       Category = "Municipality"
	OBLAST                                             Category = "Oblast"
	OUTLYING_AREA                                      Category = "Outlying Area"
	PAKISTAN_ADMINISTRERED_AREA                        Category = "Pakistan Administrered Area"
	PARISH                                             Category = "Parish"
	PARTISH                                            Category = "Partish" // Note: Likely a typo in the source data for "Parish"
	POPULARATE                                         Category = "Popularate"
	PREFECTURE                                         Category = "Prefecture"
	PROVICNE                                           Category = "Provicne" // Note: Likely a typo in the source data for "Province"
	PROVINCE                                           Category = "Province"
	QUARTER                                            Category = "Quarter"
	RAYON                                              Category = "Rayon"
	REGION                                             Category = "Region"
	REGIONAL_STATE                                     Category = "Regional State"
	REPUBLIC                                           Category = "Republic"
	SPECIAL_ADMINISTRATIVE_REGION                      Category = "Special Administrative Region"
	SPECIAL_ADMINISTRATIVE_SPECIAL_ADMINISTRATIVE_CITY Category = "Special Administrative Special Administrative City"
	SPECIAL_CITY                                       Category = "Special City"
	SPECIAL_ISLAND_AUTHORITHY                          Category = "Special Island Authorithy" // Note: Likely a typo in the source data for "Authority"
	SPECIAL_MUNICIPALITY                               Category = "Special Municipality"
	SPECIAL_REGION                                     Category = "Special Region"
	SPECIAL_SELF_GOVERNING_CITY                        Category = "Special Self-Governing City"
	SPECIAL_SELF_GOVERNING_PROVINCE                    Category = "Special Self-Governing Province"
	STATE                                              Category = "State"
	STATE_CITY                                         Category = "State City"
	TERRITORIAL_UNIT                                   Category = "Territorial Unit"
	TERRITORY                                          Category = "Territory"
	TOWN                                               Category = "Town"
	TOWN_COUNCIL                                       Category = "Town Council"
	UNION_TERRITORY                                    Category = "Union Territory"
	UNITARY_AUTHORITY                                  Category = "Unitary Authority"
	URBAN_COMMUNITY                                    Category = "Urban Community"
	URBAN_MUNICIPALITY                                 Category = "Urban Municipality"
	VOIVODSHIP                                         Category = "Voivodship"
	WARD                                               Category = "Ward"
)


func (s *SubDivCollection) PrintAll() {
	log.Println(s.subdivisions)
}

func NewSubDivisionCollection() *SubDivCollection {
	collection := SubDivCollection{
		subdivisions: Subdivisions,
	}
	return &collection
}

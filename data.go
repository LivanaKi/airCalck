package main

type City struct {
	code       string
	cityName   string
	longtitude int
	latitude   int
}

var cities = []City{
	{ "BOS", "Boston", -710589, 423601 },
	{ "LHR", "London Heathrow", 1276, 515072 },
	{ "HKG", "Hong Kong International", 1141694, 223193 },
	{ "NRT", "Tokyo Narita", 1403187, 357767 },
	{ "DEL", "New Delhi International", 772090, 286139 },
	{ "KTM", "Tribhuvan International", 853240, 277172 },
	{ "DXB", "Dubai International", 553537, 252481 },
	{ "SYD", "Sydney Kingsford Smith", 1512093, -338688 },
	{ "MEX", "Mexico City Benito Juarez", -990719, 194361 },
	{ "MIA", "Miami International", -802795, 257951 },
	{ "DFW", "Dallas Fort Worth", -970336, 328990 },
	{ "SEA", "Seattle Tacoma", -1223086, 474484 },
	{ "SIN", "Singapore Changi", 1039886, 13545 },
	{ "TPE", "Taipei Taoyuan", 1212322, 250772 },
	{ "CDG", "Paris Charles de Gaulle", 25508, 490079 },
	{ "EZE", "Buenos Aires Ezeiza International", -585373, -348165 },
	{ "JFK", "John F. Kennedy International", -737797, 406446 },
	{ "ORD", "Chicago O'Hare International", -737797, 406446},
	{ "YYZ", "Toronto Pearson International", -796334, 436771 },
	{ "MAD", "Madrid Adolfo Su‡rez MadridŠBarajas", 35643, 404895 },
	{ "IST", "Istanbul International", 287424, 412607 },
	{ "DEN", "Denver International", -1046764, 398563 },
}

type CabinClass struct {
	code      string
	className string
	rate      int
}

var cabinClasses = []CabinClass{
	{"F", "first", 60},
	{"B", "business", 35},
	{"E", "economy", 12},
}

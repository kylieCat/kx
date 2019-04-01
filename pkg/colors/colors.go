package colors

type ConsoleColors map[string]string

const (
	ColorOff          = "\033[0m"
	Aquamarine1       = "\033[0;38;5;122m"
	Aquamarine3       = "\033[0;38;5;79m"
	Black             = "\033[0;38;5;0m"
	Blue1             = "\033[0;38;5;21m"
	Blue3             = "\033[0;38;5;19m"
	Blue              = "\033[0;38;5;12m"
	Blueviolet        = "\033[0;38;5;57m"
	Cadetblue         = "\033[0;38;5;73m"
	Chartreuse1       = "\033[0;38;5;118m"
	Chartreuse2       = "\033[0;38;5;112m"
	Chartreuse3       = "\033[0;38;5;76m"
	Chartreuse4       = "\033[0;38;5;64m"
	CornSilk1         = "\033[0;38;5;230m"
	Cornflowerblue    = "\033[0;38;5;69m"
	Cyan1             = "\033[0;38;5;51m"
	Cyan2             = "\033[0;38;5;50m"
	Cyan3             = "\033[0;38;5;43m"
	DarkOliveGreen1   = "\033[0;38;5;191m"
	DarkOrange        = "\033[0;38;5;208m"
	DarkSeaGreen1     = "\033[0;38;5;193m"
	Darkblue          = "\033[0;38;5;18m"
	Darkcyan          = "\033[0;38;5;36m"
	Darkgoldenrod     = "\033[0;38;5;136m"
	Darkgreen         = "\033[0;38;5;22m"
	Darkkhaki         = "\033[0;38;5;143m"
	Darkmagenta       = "\033[0;38;5;90m"
	Darkolivegreen2   = "\033[0;38;5;155m"
	Darkolivegreen3   = "\033[0;38;5;107m"
	Darkorange3       = "\033[0;38;5;130m"
	Darkred           = "\033[0;38;5;88m"
	Darkseagreen1     = "\033[0;38;5;158m"
	Darkseagreen2     = "\033[0;38;5;151m"
	Darkseagreen3     = "\033[0;38;5;115m"
	Darkseagreen4     = "\033[0;38;5;71m"
	Darkseagreen      = "\033[0;38;5;108m"
	Darkslategray1    = "\033[0;38;5;123m"
	Darkslategray2    = "\033[0;38;5;87m"
	Darkslategray3    = "\033[0;38;5;116m"
	Darkturquoise     = "\033[0;38;5;44m"
	Darkviolet        = "\033[0;38;5;128m"
	DeepPink1         = "\033[0;38;5;198m"
	DeepPink2         = "\033[0;38;5;197m"
	Deeppink3         = "\033[0;38;5;161m"
	Deeppink4         = "\033[0;38;5;125m"
	Deepskyblue1      = "\033[0;38;5;39m"
	Deepskyblue2      = "\033[0;38;5;38m"
	Deepskyblue3      = "\033[0;38;5;31m"
	Deepskyblue4      = "\033[0;38;5;23m"
	Dodgerblue1       = "\033[0;38;5;33m"
	Dodgerblue2       = "\033[0;38;5;27m"
	Dodgerblue3       = "\033[0;38;5;26m"
	Fuchsia           = "\033[0;38;5;13m"
	Gold1             = "\033[0;38;5;220m"
	Gold3             = "\033[0;38;5;142m"
	Green1            = "\033[0;38;5;46m"
	Green3            = "\033[0;38;5;34m"
	Green4            = "\033[0;38;5;28m"
	Green             = "\033[0;38;5;2m"
	Greenyellow       = "\033[0;38;5;154m"
	Grey0             = "\033[0;38;5;16m"
	Grey100           = "\033[0;38;5;231m"
	Grey11            = "\033[0;38;5;234m"
	Grey15            = "\033[0;38;5;235m"
	Grey19            = "\033[0;38;5;236m"
	Grey23            = "\033[0;38;5;237m"
	Grey27            = "\033[0;38;5;238m"
	Grey30            = "\033[0;38;5;239m"
	Grey35            = "\033[0;38;5;240m"
	Grey37            = "\033[0;38;5;59m"
	Grey39            = "\033[0;38;5;241m"
	Grey3             = "\033[0;38;5;232m"
	Grey42            = "\033[0;38;5;242m"
	Grey46            = "\033[0;38;5;243m"
	Grey50            = "\033[0;38;5;244m"
	Grey53            = "\033[0;38;5;102m"
	Grey54            = "\033[0;38;5;245m"
	Grey58            = "\033[0;38;5;246m"
	Grey62            = "\033[0;38;5;247m"
	Grey63            = "\033[0;38;5;139m"
	Grey66            = "\033[0;38;5;248m"
	Grey69            = "\033[0;38;5;145m"
	Grey70            = "\033[0;38;5;249m"
	Grey74            = "\033[0;38;5;250m"
	Grey78            = "\033[0;38;5;251m"
	Grey7             = "\033[0;38;5;233m"
	Grey82            = "\033[0;38;5;252m"
	Grey84            = "\033[0;38;5;188m"
	Grey85            = "\033[0;38;5;253m"
	Grey89            = "\033[0;38;5;254m"
	Grey93            = "\033[0;38;5;255m"
	Grey              = "\033[0;38;5;8m"
	Honeydew2         = "\033[0;38;5;194m"
	HotPink           = "\033[0;38;5;206m"
	Hotpink2          = "\033[0;38;5;169m"
	Hotpink3          = "\033[0;38;5;132m"
	IndianRed1        = "\033[0;38;5;204m"
	Indianred         = "\033[0;38;5;131m"
	Khaki1            = "\033[0;38;5;228m"
	Khaki3            = "\033[0;38;5;185m"
	LightCoral        = "\033[0;38;5;210m"
	LightCyan1        = "\033[0;38;5;195m"
	LightGoldenrod1   = "\033[0;38;5;227m"
	LightGoldenrod2   = "\033[0;38;5;221m"
	LightPink1        = "\033[0;38;5;217m"
	LightSalmon1      = "\033[0;38;5;216m"
	LightSteelBlue1   = "\033[0;38;5;189m"
	Lightcyan3        = "\033[0;38;5;152m"
	Lightgoldenrod2   = "\033[0;38;5;186m"
	Lightgoldenrod3   = "\033[0;38;5;179m"
	Lightgreen        = "\033[0;38;5;119m"
	Lightpink3        = "\033[0;38;5;174m"
	Lightpink4        = "\033[0;38;5;95m"
	Lightsalmon3      = "\033[0;38;5;137m"
	Lightseagreen     = "\033[0;38;5;37m"
	Lightskyblue1     = "\033[0;38;5;153m"
	Lightskyblue3     = "\033[0;38;5;109m"
	Lightslateblue    = "\033[0;38;5;105m"
	Lightslategrey    = "\033[0;38;5;103m"
	Lightsteelblue3   = "\033[0;38;5;146m"
	Lightsteelblue    = "\033[0;38;5;147m"
	Lightyellow3      = "\033[0;38;5;187m"
	Lime              = "\033[0;38;5;10m"
	Magenta1          = "\033[0;38;5;201m"
	Magenta2          = "\033[0;38;5;165m"
	Magenta3          = "\033[0;38;5;127m"
	Maroon            = "\033[0;38;5;1m"
	MayaBlue          = "\033[0;38;5;81m"
	MediumOrchid1     = "\033[0;38;5;207m"
	Mediumorchid1     = "\033[0;38;5;171m"
	Mediumorchid3     = "\033[0;38;5;133m"
	Mediumorchid      = "\033[0;38;5;134m"
	Mediumpurple1     = "\033[0;38;5;141m"
	Mediumpurple2     = "\033[0;38;5;135m"
	Mediumpurple3     = "\033[0;38;5;97m"
	Mediumpurple4     = "\033[0;38;5;60m"
	Mediumpurple      = "\033[0;38;5;104m"
	Mediumspringgreen = "\033[0;38;5;49m"
	Mediumturquoise   = "\033[0;38;5;80m"
	Mediumvioletred   = "\033[0;38;5;126m"
	MistyRose1        = "\033[0;38;5;224m"
	Mistyrose3        = "\033[0;38;5;181m"
	NavajoWhite1      = "\033[0;38;5;223m"
	Navajowhite3      = "\033[0;38;5;144m"
	Navy              = "\033[0;38;5;4m"
	Navyblue          = "\033[0;38;5;17m"
	Olive             = "\033[0;38;5;3m"
	Orange1           = "\033[0;38;5;214m"
	Orange3           = "\033[0;38;5;172m"
	Orange4           = "\033[0;38;5;94m"
	OrangeRed1        = "\033[0;38;5;202m"
	Orchid1           = "\033[0;38;5;213m"
	Orchid2           = "\033[0;38;5;212m"
	Orchid            = "\033[0;38;5;170m"
	PaleVioletred1    = "\033[0;38;5;211m"
	Palegreen1        = "\033[0;38;5;121m"
	Palegreen3        = "\033[0;38;5;114m"
	Paleturquoise1    = "\033[0;38;5;159m"
	Paleturquoise4    = "\033[0;38;5;66m"
	Pink1             = "\033[0;38;5;218m"
	Pink3             = "\033[0;38;5;175m"
	Plum1             = "\033[0;38;5;219m"
	Plum2             = "\033[0;38;5;183m"
	Plum3             = "\033[0;38;5;176m"
	Plum4             = "\033[0;38;5;96m"
	Purple3           = "\033[0;38;5;56m"
	Purple4           = "\033[0;38;5;55m"
	Purple            = "\033[0;38;5;129m"
	Red1              = "\033[0;38;5;196m"
	Red3              = "\033[0;38;5;124m"
	Red               = "\033[0;38;5;9m"
	Rosybrown         = "\033[0;38;5;138m"
	Royalblue1        = "\033[0;38;5;63m"
	Salmon1           = "\033[0;38;5;209m"
	SandyBrow         = "\033[0;38;5;215m"
	Seagreen1         = "\033[0;38;5;84m"
	Seagreen2         = "\033[0;38;5;83m"
	Seagreen3         = "\033[0;38;5;78m"
	Silver            = "\033[0;38;5;7m"
	Skyblue           = "\033[0;38;5;117m"
	Skyblue3          = "\033[0;38;5;74m"
	Slateblue1        = "\033[0;38;5;99m"
	Slateblue3        = "\033[0;38;5;62m"
	Springgreen1      = "\033[0;38;5;48m"
	Springgreen2      = "\033[0;38;5;47m"
	Springgreen3      = "\033[0;38;5;35m"
	Springgreen4      = "\033[0;38;5;29m"
	Steelblue1        = "\033[0;38;5;75m"
	Steelblue3        = "\033[0;38;5;68m"
	Steelblue         = "\033[0;38;5;67m"
	Tan               = "\033[0;38;5;180m"
	Teal              = "\033[0;38;5;6m"
	Thistle1          = "\033[0;38;5;225m"
	Thistle3          = "\033[0;38;5;182m"
	Turquoise2        = "\033[0;38;5;45m"
	Turquoise4        = "\033[0;38;5;30m"
	Violet            = "\033[0;38;5;177m"
	Wheat1            = "\033[0;38;5;229m"
	Wheat4            = "\033[0;38;5;101m"
	White             = "\033[0;38;5;15m"
	Yellow1           = "\033[0;38;5;226m"
	Yellow2           = "\033[0;38;5;190m"
	Yellow3           = "\033[0;38;5;148m"
	Yellow4           = "\033[0;38;5;100m"
	Yellow            = "\033[0;38;5;11m"
	K8sBlue           = "\033[0;38;5;38m"

	// Names
	ColorOffName          = "ColorOff"
	Aquamarine1Name       = "Aquamarine1"
	Aquamarine3Name       = "Aquamarine3"
	BlackName             = "Black"
	Blue1Name             = "Blue1"
	Blue3Name             = "Blue3"
	BlueName              = "Blue"
	BluevioletName        = "Blueviolet"
	CadetblueName         = "Cadetblue"
	Chartreuse1Name       = "Chartreuse1"
	Chartreuse2Name       = "Chartreuse2"
	Chartreuse3Name       = "Chartreuse3"
	Chartreuse4Name       = "Chartreuse4"
	CornSilk1Name         = "CornSilk1"
	CornflowerblueName    = "Cornflowerblue"
	Cyan1Name             = "Cyan1"
	Cyan2Name             = "Cyan2"
	Cyan3Name             = "Cyan3"
	DarkOliveGreen1Name   = "DarkOliveGreen1"
	DarkOrangeName        = "DarkOrange"
	DarkSeaGreen1Name     = "DarkSeaGreen1"
	DarkblueName          = "Darkblue"
	DarkcyanName          = "Darkcyan"
	DarkgoldenrodName     = "Darkgoldenrod"
	DarkgreenName         = "Darkgreen"
	DarkkhakiName         = "Darkkhaki"
	DarkmagentaName       = "Darkmagenta"
	Darkolivegreen2Name   = "Darkolivegreen2"
	Darkolivegreen3Name   = "Darkolivegreen3"
	Darkorange3Name       = "Darkorange3"
	DarkredName           = "Darkred"
	Darkseagreen1Name     = "Darkseagreen1"
	Darkseagreen2Name     = "Darkseagreen2"
	Darkseagreen3Name     = "Darkseagreen3"
	Darkseagreen4Name     = "Darkseagreen4"
	DarkseagreenName      = "Darkseagreen"
	Darkslategray1Name    = "Darkslategray1"
	Darkslategray2Name    = "Darkslategray2"
	Darkslategray3Name    = "Darkslategray3"
	DarkturquoiseName     = "Darkturquoise"
	DarkvioletName        = "Darkviolet"
	DeepPink1Name         = "DeepPink1"
	DeepPink2Name         = "DeepPink2"
	Deeppink3Name         = "Deeppink3"
	Deeppink4Name         = "Deeppink4"
	Deepskyblue1Name      = "Deepskyblue1"
	Deepskyblue2Name      = "Deepskyblue2"
	Deepskyblue3Name      = "Deepskyblue3"
	Deepskyblue4Name      = "Deepskyblue4"
	Dodgerblue1Name       = "Dodgerblue1"
	Dodgerblue2Name       = "Dodgerblue2"
	Dodgerblue3Name       = "Dodgerblue3"
	FuchsiaName           = "Fuchsia"
	Gold1Name             = "Gold1"
	Gold3Name             = "Gold3"
	Green1Name            = "Green1"
	Green3Name            = "Green3"
	Green4Name            = "Green4"
	GreenName             = "Green"
	GreenyellowName       = "Greenyellow"
	Grey0Name             = "Grey0"
	Grey100Name           = "Grey100"
	Grey11Name            = "Grey11"
	Grey15Name            = "Grey15"
	Grey19Name            = "Grey19"
	Grey23Name            = "Grey23"
	Grey27Name            = "Grey27"
	Grey30Name            = "Grey30"
	Grey35Name            = "Grey35"
	Grey37Name            = "Grey37"
	Grey39Name            = "Grey39"
	Grey3Name             = "Grey3"
	Grey42Name            = "Grey42"
	Grey46Name            = "Grey46"
	Grey50Name            = "Grey50"
	Grey53Name            = "Grey53"
	Grey54Name            = "Grey54"
	Grey58Name            = "Grey58"
	Grey62Name            = "Grey62"
	Grey63Name            = "Grey63"
	Grey66Name            = "Grey66"
	Grey69Name            = "Grey69"
	Grey70Name            = "Grey70"
	Grey74Name            = "Grey74"
	Grey78Name            = "Grey78"
	Grey7Name             = "Grey7"
	Grey82Name            = "Grey82"
	Grey84Name            = "Grey84"
	Grey85Name            = "Grey85"
	Grey89Name            = "Grey89"
	Grey93Name            = "Grey93"
	GreyName              = "Grey"
	Honeydew2Name         = "Honeydew2"
	HotPinkName           = "HotPink"
	Hotpink2Name          = "Hotpink2"
	Hotpink3Name          = "Hotpink3"
	IndianRed1Name        = "IndianRed1"
	IndianredName         = "Indianred"
	Khaki1Name            = "Khaki1"
	Khaki3Name            = "Khaki3"
	LightCoralName        = "LightCoral"
	LightCyan1Name        = "LightCyan1"
	LightGoldenrod1Name   = "LightGoldenrod1"
	LightGoldenrod2Name   = "LightGoldenrod2"
	LightPink1Name        = "LightPink1"
	LightSalmon1Name      = "LightSalmon1"
	LightSteelBlue1Name   = "LightSteelBlue1"
	Lightcyan3Name        = "Lightcyan3"
	Lightgoldenrod2Name   = "Lightgoldenrod2"
	Lightgoldenrod3Name   = "Lightgoldenrod3"
	LightgreenName        = "Lightgreen"
	Lightpink3Name        = "Lightpink3"
	Lightpink4Name        = "Lightpink4"
	Lightsalmon3Name      = "Lightsalmon3"
	LightseagreenName     = "Lightseagreen"
	Lightskyblue1Name     = "Lightskyblue1"
	Lightskyblue3Name     = "Lightskyblue3"
	LightslateblueName    = "Lightslateblue"
	LightslategreyName    = "Lightslategrey"
	Lightsteelblue3Name   = "Lightsteelblue3"
	LightsteelblueName    = "Lightsteelblue"
	Lightyellow3Name      = "Lightyellow3"
	LimeName              = "Lime"
	Magenta1Name          = "Magenta1"
	Magenta2Name          = "Magenta2"
	Magenta3Name          = "Magenta3"
	MaroonName            = "Maroon"
	MayaBlueName          = "MayaBlue"
	MediumOrchid1Name     = "MediumOrchid1"
	Mediumorchid1Name     = "Mediumorchid1"
	Mediumorchid3Name     = "Mediumorchid3"
	MediumorchidName      = "Mediumorchid"
	Mediumpurple1Name     = "Mediumpurple1"
	Mediumpurple2Name     = "Mediumpurple2"
	Mediumpurple3Name     = "Mediumpurple3"
	Mediumpurple4Name     = "Mediumpurple4"
	MediumpurpleName      = "Mediumpurple"
	MediumspringgreenName = "Mediumspringgre"
	MediumturquoiseName   = "Mediumturquoise"
	MediumvioletredName   = "Mediumvioletred"
	MistyRose1Name        = "MistyRose1"
	Mistyrose3Name        = "Mistyrose3"
	NavajoWhite1Name      = "NavajoWhite1"
	Navajowhite3Name      = "Navajowhite3"
	NavyName              = "Navy"
	NavyblueName          = "Navyblue"
	OliveName             = "Olive"
	Orange1Name           = "Orange1"
	Orange3Name           = "Orange3"
	Orange4Name           = "Orange4"
	OrangeRed1Name        = "OrangeRed1"
	Orchid1Name           = "Orchid1"
	Orchid2Name           = "Orchid2"
	OrchidName            = "Orchid"
	PaleVioletred1Name    = "PaleVioletred1"
	Palegreen1Name        = "Palegreen1"
	Palegreen3Name        = "Palegreen3"
	Paleturquoise1Name    = "Paleturquoise1"
	Paleturquoise4Name    = "Paleturquoise4"
	Pink1Name             = "Pink1"
	Pink3Name             = "Pink3"
	Plum1Name             = "Plum1"
	Plum2Name             = "Plum2"
	Plum3Name             = "Plum3"
	Plum4Name             = "Plum4"
	Purple3Name           = "Purple3"
	Purple4Name           = "Purple4"
	PurpleName            = "Purple"
	Red1Name              = "Red1"
	Red3Name              = "Red3"
	RedName               = "Red"
	RosybrownName         = "Rosybrown"
	Royalblue1Name        = "Royalblue1"
	Salmon1Name           = "Salmon1"
	Seagreen1Name         = "Seagreen1"
	Seagreen2Name         = "Seagreen2"
	Seagreen3Name         = "Seagreen3"
	SilverName            = "Silver"
	Skyblue3Name          = "Skyblue3"
	Slateblue1Name        = "Slateblue1"
	Slateblue3Name        = "Slateblue3"
	Springgreen1Name      = "Springgreen1"
	Springgreen2Name      = "Springgreen2"
	Springgreen3Name      = "Springgreen3"
	Springgreen4Name      = "Springgreen4"
	Steelblue1Name        = "Steelblue1 "
	Steelblue3Name        = "Steelblue3 "
	SteelblueName         = "Steelblue"
	TanName               = "Tan"
	TealName              = "Teal"
	Thistle1Name          = "Thistle1"
	Thistle3Name          = "Thistle3"
	Turquoise2Name        = "Turquoise2 "
	Turquoise4Name        = "Turquoise4 "
	VioletName            = "Violet"
	Wheat1Name            = "Wheat1"
	Wheat4Name            = "Wheat4"
	WhiteName             = "White"
	Yellow1Name           = "Yellow1"
	Yellow2Name           = "Yellow2"
	Yellow3Name           = "Yellow3"
	Yellow4Name           = "Yellow4"
	YellowName            = "Yellow"
	K8sBlueName           = "K8sBlue"
)

var ColorOrder = []string{
	Aquamarine1Name,
	Aquamarine3Name,
	BlackName,
	Blue1Name,
	Blue3Name,
	BlueName,
	BluevioletName,
	CadetblueName,
	Chartreuse1Name,
	Chartreuse2Name,
	Chartreuse3Name,
	Chartreuse4Name,
	CornSilk1Name,
	CornflowerblueName,
	Cyan1Name,
	Cyan2Name,
	Cyan3Name,
	DarkOliveGreen1Name,
	DarkOrangeName,
	DarkSeaGreen1Name,
	DarkblueName,
	DarkcyanName,
	DarkgoldenrodName,
	DarkgreenName,
	DarkkhakiName,
	DarkmagentaName,
	Darkolivegreen2Name,
	Darkolivegreen3Name,
	Darkorange3Name,
	DarkredName,
	Darkseagreen1Name,
	Darkseagreen2Name,
	Darkseagreen3Name,
	Darkseagreen4Name,
	DarkseagreenName,
	Darkslategray1Name,
	Darkslategray2Name,
	Darkslategray3Name,
	DarkturquoiseName,
	DarkvioletName,
	DeepPink1Name,
	DeepPink2Name,
	Deeppink3Name,
	Deeppink4Name,
	Deepskyblue1Name,
	Deepskyblue2Name,
	Deepskyblue3Name,
	Deepskyblue4Name,
	Dodgerblue1Name,
	Dodgerblue2Name,
	Dodgerblue3Name,
	FuchsiaName,
	Gold1Name,
	Gold3Name,
	Green1Name,
	Green3Name,
	Green4Name,
	GreenName,
	GreenyellowName,
	Grey0Name,
	Grey100Name,
	Grey11Name,
	Grey15Name,
	Grey19Name,
	Grey23Name,
	Grey27Name,
	Grey30Name,
	Grey35Name,
	Grey37Name,
	Grey39Name,
	Grey3Name,
	Grey42Name,
	Grey46Name,
	Grey50Name,
	Grey53Name,
	Grey54Name,
	Grey58Name,
	Grey62Name,
	Grey63Name,
	Grey66Name,
	Grey69Name,
	Grey70Name,
	Grey74Name,
	Grey78Name,
	Grey7Name,
	Grey82Name,
	Grey84Name,
	Grey85Name,
	Grey89Name,
	Grey93Name,
	GreyName,
	Honeydew2Name,
	HotPinkName,
	Hotpink2Name,
	Hotpink3Name,
	IndianRed1Name,
	IndianredName,
	Khaki1Name,
	Khaki3Name,
	LightCoralName,
	LightCyan1Name,
	LightGoldenrod1Name,
	LightGoldenrod2Name,
	LightPink1Name,
	LightSalmon1Name,
	LightSteelBlue1Name,
	Lightcyan3Name,
	Lightgoldenrod2Name,
	Lightgoldenrod3Name,
	LightgreenName,
	Lightpink3Name,
	Lightpink4Name,
	Lightsalmon3Name,
	LightseagreenName,
	Lightskyblue1Name,
	Lightskyblue3Name,
	LightslateblueName,
	LightslategreyName,
	Lightsteelblue3Name,
	LightsteelblueName,
	Lightyellow3Name,
	LimeName,
	Magenta1Name,
	Magenta2Name,
	Magenta3Name,
	MaroonName,
	MayaBlueName,
	MediumOrchid1Name,
	Mediumorchid1Name,
	Mediumorchid3Name,
	MediumorchidName,
	Mediumpurple1Name,
	Mediumpurple2Name,
	Mediumpurple3Name,
	Mediumpurple4Name,
	MediumpurpleName,
	MediumspringgreenName,
	MediumturquoiseName,
	MediumvioletredName,
	MistyRose1Name,
	Mistyrose3Name,
	NavajoWhite1Name,
	Navajowhite3Name,
	NavyName,
	NavyblueName,
	OliveName,
	Orange1Name,
	Orange3Name,
	Orange4Name,
	OrangeRed1Name,
	Orchid1Name,
	Orchid2Name,
	OrchidName,
	PaleVioletred1Name,
	Palegreen1Name,
	Palegreen3Name,
	Paleturquoise1Name,
	Paleturquoise4Name,
	Pink1Name,
	Pink3Name,
	Plum1Name,
	Plum2Name,
	Plum3Name,
	Plum4Name,
	Purple3Name,
	Purple4Name,
	PurpleName,
	Red1Name,
	Red3Name,
	RedName,
	RosybrownName,
	Royalblue1Name,
	Salmon1Name,
	Seagreen1Name,
	Seagreen2Name,
	Seagreen3Name,
	SilverName,
	Skyblue3Name,
	Slateblue1Name,
	Slateblue3Name,
	Springgreen1Name,
	Springgreen2Name,
	Springgreen3Name,
	Springgreen4Name,
	Steelblue1Name,
	Steelblue3Name,
	SteelblueName,
	TanName,
	TealName,
	Thistle1Name,
	Thistle3Name,
	Turquoise2Name,
	Turquoise4Name,
	VioletName,
	Wheat1Name,
	Wheat4Name,
	WhiteName,
	Yellow1Name,
	Yellow2Name,
	Yellow3Name,
	Yellow4Name,
	YellowName,
	K8sBlueName,
}

var Colors = ConsoleColors{
	ColorOffName:          ColorOff,
	Aquamarine1Name:       Aquamarine1,
	Aquamarine3Name:       Aquamarine3,
	BlackName:             Black,
	Blue1Name:             Blue1,
	Blue3Name:             Blue3,
	BlueName:              Blue,
	BluevioletName:        Blueviolet,
	CadetblueName:         Cadetblue,
	Chartreuse1Name:       Chartreuse1,
	Chartreuse2Name:       Chartreuse2,
	Chartreuse3Name:       Chartreuse3,
	Chartreuse4Name:       Chartreuse4,
	CornSilk1Name:         CornSilk1,
	CornflowerblueName:    Cornflowerblue,
	Cyan1Name:             Cyan1,
	Cyan2Name:             Cyan2,
	Cyan3Name:             Cyan3,
	DarkOliveGreen1Name:   DarkOliveGreen1,
	DarkOrangeName:        DarkOrange,
	DarkSeaGreen1Name:     DarkSeaGreen1,
	DarkblueName:          Darkblue,
	DarkcyanName:          Darkcyan,
	DarkgoldenrodName:     Darkgoldenrod,
	DarkgreenName:         Darkgreen,
	DarkkhakiName:         Darkkhaki,
	DarkmagentaName:       Darkmagenta,
	Darkolivegreen2Name:   Darkolivegreen2,
	Darkolivegreen3Name:   Darkolivegreen3,
	Darkorange3Name:       Darkorange3,
	DarkredName:           Darkred,
	Darkseagreen1Name:     Darkseagreen1,
	Darkseagreen2Name:     Darkseagreen2,
	Darkseagreen3Name:     Darkseagreen3,
	Darkseagreen4Name:     Darkseagreen4,
	DarkseagreenName:      Darkseagreen,
	Darkslategray1Name:    Darkslategray1,
	Darkslategray2Name:    Darkslategray2,
	Darkslategray3Name:    Darkslategray3,
	DarkturquoiseName:     Darkturquoise,
	DarkvioletName:        Darkviolet,
	DeepPink1Name:         DeepPink1,
	DeepPink2Name:         DeepPink2,
	Deeppink3Name:         Deeppink3,
	Deeppink4Name:         Deeppink4,
	Deepskyblue1Name:      Deepskyblue1,
	Deepskyblue2Name:      Deepskyblue2,
	Deepskyblue3Name:      Deepskyblue3,
	Deepskyblue4Name:      Deepskyblue4,
	Dodgerblue1Name:       Dodgerblue1,
	Dodgerblue2Name:       Dodgerblue2,
	Dodgerblue3Name:       Dodgerblue3,
	FuchsiaName:           Fuchsia,
	Gold1Name:             Gold1,
	Gold3Name:             Gold3,
	Green1Name:            Green1,
	Green3Name:            Green3,
	Green4Name:            Green4,
	GreenName:             Green,
	GreenyellowName:       Greenyellow,
	Grey0Name:             Grey0,
	Grey100Name:           Grey100,
	Grey11Name:            Grey11,
	Grey15Name:            Grey15,
	Grey19Name:            Grey19,
	Grey23Name:            Grey23,
	Grey27Name:            Grey27,
	Grey30Name:            Grey30,
	Grey35Name:            Grey35,
	Grey37Name:            Grey37,
	Grey39Name:            Grey39,
	Grey3Name:             Grey3,
	Grey42Name:            Grey42,
	Grey46Name:            Grey46,
	Grey50Name:            Grey50,
	Grey53Name:            Grey53,
	Grey54Name:            Grey54,
	Grey58Name:            Grey58,
	Grey62Name:            Grey62,
	Grey63Name:            Grey63,
	Grey66Name:            Grey66,
	Grey69Name:            Grey69,
	Grey70Name:            Grey70,
	Grey74Name:            Grey74,
	Grey78Name:            Grey78,
	Grey7Name:             Grey7,
	Grey82Name:            Grey82,
	Grey84Name:            Grey84,
	Grey85Name:            Grey85,
	Grey89Name:            Grey89,
	Grey93Name:            Grey93,
	GreyName:              Grey,
	Honeydew2Name:         Honeydew2,
	HotPinkName:           HotPink,
	Hotpink2Name:          Hotpink2,
	Hotpink3Name:          Hotpink3,
	IndianRed1Name:        IndianRed1,
	IndianredName:         Indianred,
	Khaki1Name:            Khaki1,
	Khaki3Name:            Khaki3,
	LightCoralName:        LightCoral,
	LightCyan1Name:        LightCyan1,
	LightGoldenrod1Name:   LightGoldenrod1,
	LightGoldenrod2Name:   LightGoldenrod2,
	LightPink1Name:        LightPink1,
	LightSalmon1Name:      LightSalmon1,
	LightSteelBlue1Name:   LightSteelBlue1,
	Lightcyan3Name:        Lightcyan3,
	Lightgoldenrod2Name:   Lightgoldenrod2,
	Lightgoldenrod3Name:   Lightgoldenrod3,
	LightgreenName:        Lightgreen,
	Lightpink3Name:        Lightpink3,
	Lightpink4Name:        Lightpink4,
	Lightsalmon3Name:      Lightsalmon3,
	LightseagreenName:     Lightseagreen,
	Lightskyblue1Name:     Lightskyblue1,
	Lightskyblue3Name:     Lightskyblue3,
	LightslateblueName:    Lightslateblue,
	LightslategreyName:    Lightslategrey,
	Lightsteelblue3Name:   Lightsteelblue3,
	LightsteelblueName:    Lightsteelblue,
	Lightyellow3Name:      Lightyellow3,
	LimeName:              Lime,
	Magenta1Name:          Magenta1,
	Magenta2Name:          Magenta2,
	Magenta3Name:          Magenta3,
	MaroonName:            Maroon,
	MayaBlueName:          MayaBlue,
	MediumOrchid1Name:     MediumOrchid1,
	Mediumorchid1Name:     Mediumorchid1,
	Mediumorchid3Name:     Mediumorchid3,
	MediumorchidName:      Mediumorchid,
	Mediumpurple1Name:     Mediumpurple1,
	Mediumpurple2Name:     Mediumpurple2,
	Mediumpurple3Name:     Mediumpurple3,
	Mediumpurple4Name:     Mediumpurple4,
	MediumpurpleName:      Mediumpurple,
	MediumspringgreenName: Mediumspringgreen,
	MediumturquoiseName:   Mediumturquoise,
	MediumvioletredName:   Mediumvioletred,
	MistyRose1Name:        MistyRose1,
	Mistyrose3Name:        Mistyrose3,
	NavajoWhite1Name:      NavajoWhite1,
	Navajowhite3Name:      Navajowhite3,
	NavyName:              Navy,
	NavyblueName:          Navyblue,
	OliveName:             Olive,
	Orange1Name:           Orange1,
	Orange3Name:           Orange3,
	Orange4Name:           Orange4,
	OrangeRed1Name:        OrangeRed1,
	Orchid1Name:           Orchid1,
	Orchid2Name:           Orchid2,
	OrchidName:            Orchid,
	PaleVioletred1Name:    PaleVioletred1,
	Palegreen1Name:        Palegreen1,
	Palegreen3Name:        Palegreen3,
	Paleturquoise1Name:    Paleturquoise1,
	Paleturquoise4Name:    Paleturquoise4,
	Pink1Name:             Pink1,
	Pink3Name:             Pink3,
	Plum1Name:             Plum1,
	Plum2Name:             Plum2,
	Plum3Name:             Plum3,
	Plum4Name:             Plum4,
	Purple3Name:           Purple3,
	Purple4Name:           Purple4,
	PurpleName:            Purple,
	Red1Name:              Red1,
	Red3Name:              Red3,
	RedName:               Red,
	RosybrownName:         Rosybrown,
	Royalblue1Name:        Royalblue1,
	Salmon1Name:           Salmon1,
	Seagreen1Name:         Seagreen1,
	Seagreen2Name:         Seagreen2,
	Seagreen3Name:         Seagreen3,
	SilverName:            Silver,
	Skyblue3Name:          Skyblue3,
	Slateblue1Name:        Slateblue1,
	Slateblue3Name:        Slateblue3,
	Springgreen1Name:      Springgreen1,
	Springgreen2Name:      Springgreen2,
	Springgreen3Name:      Springgreen3,
	Springgreen4Name:      Springgreen4,
	Steelblue1Name:        Steelblue1,
	Steelblue3Name:        Steelblue3,
	SteelblueName:         Steelblue,
	TanName:               Tan,
	TealName:              Teal,
	Thistle1Name:          Thistle1,
	Thistle3Name:          Thistle3,
	Turquoise2Name:        Turquoise2,
	Turquoise4Name:        Turquoise4,
	VioletName:            Violet,
	Wheat1Name:            Wheat1,
	Wheat4Name:            Wheat4,
	WhiteName:             White,
	Yellow1Name:           Yellow1,
	Yellow2Name:           Yellow2,
	Yellow3Name:           Yellow3,
	Yellow4Name:           Yellow4,
	YellowName:            Yellow,
	K8sBlueName:           K8sBlue,
}

func (c ConsoleColors) Get(name string) string {
	var color string
	var ok bool

	if color, ok = c[name]; !ok {
		color = c["K8sBlue"]
	}
	return color
}

func (c ConsoleColors) GetSafe(name string) (string, bool) {
	var color string
	var ok bool

	color, ok = c[name]
	return color, ok
}

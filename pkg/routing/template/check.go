package template

import (
	"errors"
	"regexp"

	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/lib/pq"
)

var (
	// ErrNoRefID is returned, if no ref id is set in a config
	ErrNoRefID = errors.New("no RefID set")

	// ErrNoName is returned, if no name is set in a config
	ErrNoName = errors.New("no Name set")

	// ErrNoListenStatement is returned, if no listen statment is present
	ErrNoListenStatement = errors.New("no ListenStatement set")

	// ErrPortRange is returned, if port is <1024
	ErrPortRange = errors.New("port not in acceptable range")

	// ErrKeyword is returned, if the used keyword is not allowed in the used router
	ErrKeyword = errors.New("keyword not allowed")

	// ErrEmptyServerName is returned, if the length of the server name is 0
	ErrEmptyServerName = errors.New("no server name exists")

	// ErrInvalidName is returned, if a servername is no url
	ErrInvalidName = errors.New("servername no valid url format")

	urlRegex = regexp.MustCompile(`^([^\pM\pC\pZ]+\.)+(aaa|aarp|abarth|abb|abbott|abbvie|abc|able|abogado|abudhabi|ac|academy|accenture|accountant|accountants|aco|active|actor|ad|adac|ads|adult|ae|aeg|aero|aetna|af|afamilycompany|afl|africa|ag|agakhan|agency|ai|aig|aigo|airbus|airforce|airtel|akdn|al|alfaromeo|alibaba|alipay|allfinanz|allstate|ally|alsace|alstom|am|americanexpress|americanfamily|amex|amfam|amica|amsterdam|an|analytics|android|anquan|anz|ao|aol|apartments|app|apple|aq|aquarelle|ar|aramco|archi|army|arpa|art|arte|as|asda|asia|associates|at|athleta|attorney|au|auction|audi|audible|audio|auspost|author|auto|autos|avianca|aw|aws|ax|axa|az|azure|ba|baby|baidu|banamex|bananarepublic|band|bank|bar|barcelona|barclaycard|barclays|barefoot|bargains|baseball|basketball|bauhaus|bayern|bb|bbc|bbt|bbva|bcg|bcn|bd|be|beats|beauty|beer|bentley|berlin|best|bestbuy|bet|bf|bg|bh|bharti|bi|bible|bid|bike|bing|bingo|bio|biz|bj|bl|black|blackfriday|blanco|blockbuster|blog|bloomberg|blue|bm|bms|bmw|bn|bnl|bnpparibas|bo|boats|boehringer|bofa|bom|bond|boo|book|booking|boots|bosch|bostik|boston|bot|boutique|box|bq|br|bradesco|bridgestone|broadway|broker|brother|brussels|bs|bt|budapest|bugatti|build|builders|business|buy|buzz|bv|bw|by|bz|bzh|ca|cab|cafe|cal|call|calvinklein|cam|camera|camp|cancerresearch|canon|capetown|capital|capitalone|car|caravan|cards|care|career|careers|cars|cartier|casa|case|caseih|cash|casino|cat|catering|catholic|cba|cbn|cbre|cbs|cc|cd|ceb|center|ceo|cern|cf|cfa|cfd|cg|ch|chanel|channel|chase|chat|cheap|chintai|chloe|christmas|chrome|chrysler|church|ci|cipriani|circle|cisco|citadel|citi|citic|city|cityeats|ck|cl|claims|cleaning|click|clinic|clinique|clothing|cloud|club|clubmed|cm|cn|co|coach|codes|coffee|college|cologne|com|comcast|commbank|community|company|compare|computer|comsec|condos|construction|consulting|contact|contractors|cooking|cookingchannel|cool|coop|corsica|country|coupon|coupons|courses|cr|credit|creditcard|creditunion|cricket|crown|crs|cruise|cruises|csc|cu|cuisinella|cv|cw|cx|cy|cymru|cyou|cz|dabur|dad|dance|data|date|dating|datsun|day|dclk|dds|de|deal|dealer|deals|degree|delivery|dell|deloitte|delta|democrat|dental|dentist|desi|design|dev|dhl|diamonds|diet|digital|direct|directory|discount|discover|dish|diy|dj|dk|dm|dnp|do|docs|doctor|dodge|dog|doha|domains|doosan|dot|download|drive|dtv|dubai|duck|dunlop|duns|dupont|durban|dvag|dvr|dz|earth|eat|ec|eco|edeka|edu|education|ee|eg|eh|email|emerck|energy|engineer|engineering|enterprises|epost|epson|equipment|er|ericsson|erni|es|esq|estate|esurance|et|eu|eurovision|eus|events|everbank|exchange|expert|exposed|express|extraspace|fage|fail|fairwinds|faith|family|fan|fans|farm|farmers|fashion|fast|fedex|feedback|ferrari|ferrero|fi|fiat|fidelity|fido|film|final|finance|financial|fire|firestone|firmdale|fish|fishing|fit|fitness|fj|fk|flickr|flights|flir|florist|flowers|flsmidth|fly|fm|fo|foo|food|foodnetwork|football|ford|forex|forsale|forum|foundation|fox|fr|free|fresenius|frl|frogans|frontdoor|frontier|ftr|fujitsu|fujixerox|fun|fund|furniture|futbol|fyi|ga|gal|gallery|gallo|gallup|game|games|gap|garden|gb|gbiz|gd|gdn|ge|gea|gent|genting|george|gf|gg|ggee|gh|gi|gift|gifts|gives|giving|gl|glade|glass|gle|global|globo|gm|gmail|gmbh|gmo|gmx|gn|godaddy|gold|goldpoint|golf|goo|goodhands|goodyear|goog|google|gop|got|gov|gp|gq|gr|grainger|graphics|gratis|green|gripe|group|gs|gt|gu|guardian|gucci|guge|guide|guitars|guru|gw|gy|hair|hamburg|hangout|haus|hbo|hdfc|hdfcbank|health|healthcare|help|helsinki|here|hermes|hgtv|hiphop|hisamitsu|hitachi|hiv|hk|hkt|hm|hn|hockey|holdings|holiday|homedepot|homegoods|homes|homesense|honda|honeywell|horse|hospital|host|hosting|hot|hoteles|hotmail|house|how|hr|hsbc|ht|htc|hu|hughes|hyatt|hyundai|ibm|icbc|ice|icu|id|ie|ieee|ifm|iinet|ikano|il|im|imamat|imdb|immo|immobilien|in|industries|infiniti|info|ing|ink|institute|insurance|insure|int|intel|international|intuit|investments|io|ipiranga|iq|ir|irish|is|iselect|ismaili|ist|istanbul|it|itau|itv|iveco|iwc|jaguar|java|jcb|jcp|je|jeep|jetzt|jewelry|jio|jlc|jll|jm|jmp|jnj|jo|jobs|joburg|jot|joy|jp|jpmorgan|jprs|juegos|juniper|kaufen|kddi|ke|kerryhotels|kerrylogistics|kerryproperties|kfh|kg|kh|ki|kia|kim|kinder|kindle|kitchen|kiwi|km|kn|koeln|komatsu|kosher|kp|kpmg|kpn|kr|krd|kred|kuokgroup|kw|ky|kyoto|kz|la|lacaixa|ladbrokes|lamborghini|lamer|lancaster|lancia|lancome|land|landrover|lanxess|lasalle|lat|latino|latrobe|law|lawyer|lb|lc|lds|lease|leclerc|lefrak|legal|lego|lexus|lgbt|li|liaison|lidl|life|lifeinsurance|lifestyle|lighting|like|lilly|limited|limo|lincoln|linde|link|lipsy|live|living|lixil|lk|loan|loans|locker|locus|loft|lol|london|lotte|lotto|love|lpl|lplfinancial|lr|ls|lt|ltd|ltda|lu|lundbeck|lupin|luxe|luxury|lv|ly|ma|macys|madrid|maif|maison|makeup|man|management|mango|market|marketing|markets|marriott|marshalls|maserati|mattel|mba|mc|mcd|mcdonalds|mckinsey|md|me|med|media|meet|melbourne|meme|memorial|men|menu|meo|metlife|mf|mg|mh|miami|microsoft|mil|mini|mint|mit|mitsubishi|mk|ml|mlb|mls|mm|mma|mn|mo|mobi|mobile|mobily|moda|moe|moi|mom|monash|money|monster|montblanc|mopar|mormon|mortgage|moscow|moto|motorcycles|mov|movie|movistar|mp|mq|mr|ms|msd|mt|mtn|mtpc|mtr|mu|museum|mutual|mutuelle|mv|mw|mx|my|mz|na|nab|nadex|nagoya|name|nationwide|natura|navy|nba|nc|ne|nec|net|netbank|netflix|network|neustar|new|newholland|news|next|nextdirect|nexus|nf|nfl|ng|ngo|nhk|ni|nico|nike|nikon|ninja|nissan|nissay|nl|no|nokia|northwesternmutual|norton|now|nowruz|nowtv|np|nr|nra|nrw|ntt|nu|nyc|nz|obi|observer|off|office|okinawa|olayan|olayangroup|oldnavy|ollo|om|omega|one|ong|onl|online|onyourside|ooo|open|oracle|orange|org|organic|orientexpress|origins|osaka|otsuka|ott|ovh|pa|page|pamperedchef|panasonic|panerai|paris|pars|partners|parts|party|passagens|pay|pccw|pe|pet|pf|pfizer|pg|ph|pharmacy|philips|phone|photo|photography|photos|physio|piaget|pics|pictet|pictures|pid|pin|ping|pink|pioneer|pizza|pk|pl|place|play|playstation|plumbing|plus|pm|pn|pnc|pohl|poker|politie|porn|post|pr|pramerica|praxi|press|prime|pro|prod|productions|prof|progressive|promo|properties|property|protection|pru|prudential|ps|pt|pub|pw|pwc|py|qa|qpon|quebec|quest|qvc|racing|radio|raid|re|read|realestate|realtor|realty|recipes|red|redstone|redumbrella|rehab|reise|reisen|reit|reliance|ren|rent|rentals|repair|report|republican|rest|restaurant|review|reviews|rexroth|rich|richardli|ricoh|rightathome|ril|rio|rip|rmit|ro|rocher|rocks|rodeo|rogers|room|rs|rsvp|ru|ruhr|run|rw|rwe|ryukyu|sa|saarland|safe|safety|sakura|sale|salon|samsclub|samsung|sandvik|sandvikcoromant|sanofi|sap|sapo|sarl|sas|save|saxo|sb|sbi|sbs|sc|sca|scb|schaeffler|schmidt|scholarships|school|schule|schwarz|science|scjohnson|scor|scot|sd|se|seat|secure|security|seek|select|sener|services|ses|seven|sew|sex|sexy|sfr|sg|sh|shangrila|sharp|shaw|shell|shia|shiksha|shoes|shop|shopping|shouji|show|showtime|shriram|si|silk|sina|singles|site|sj|sk|ski|skin|sky|skype|sl|sling|sm|smart|smile|sn|sncf|so|soccer|social|softbank|software|sohu|solar|solutions|song|sony|soy|space|spiegel|spot|spreadbetting|sr|srl|srt|ss|st|stada|staples|star|starhub|statebank|statefarm|statoil|stc|stcgroup|stockholm|storage|store|stream|studio|study|style|su|sucks|supplies|supply|support|surf|surgery|suzuki|sv|swatch|swiftcover|swiss|sx|sy|sydney|symantec|systems|sz|tab|taipei|talk|taobao|target|tatamotors|tatar|tattoo|tax|taxi|tc|tci|td|tdk|team|tech|technology|tel|telecity|telefonica|temasek|tennis|teva|tf|tg|th|thd|theater|theatre|tiaa|tickets|tienda|tiffany|tips|tires|tirol|tj|tjmaxx|tjx|tk|tkmaxx|tl|tm|tmall|tn|to|today|tokyo|tools|top|toray|toshiba|total|tours|town|toyota|toys|tp|tr|trade|trading|training|travel|travelchannel|travelers|travelersinsurance|trust|trv|tt|tube|tui|tunes|tushu|tv|tvs|tw|tz|ua|ubank|ubs|uconnect|ug|uk|um|unicom|university|uno|uol|ups|us|uy|uz|va|vacations|vana|vanguard|vc|ve|vegas|ventures|verisign|vermögensberater|vermögensberatung|versicherung|vet|vg|vi|viajes|video|vig|viking|villas|vin|vip|virgin|visa|vision|vista|vistaprint|viva|vivo|vlaanderen|vn|vodka|volkswagen|volvo|vote|voting|voto|voyage|vu|vuelos|wales|walmart|walter|wang|wanggou|warman|watch|watches|weather|weatherchannel|webcam|weber|website|wed|wedding|weibo|weir|wf|whoswho|wien|wiki|williamhill|win|windows|wine|winners|wme|wolterskluwer|woodside|work|works|world|wow|ws|wtc|wtf|xbox|xerox|xfinity|xihuan|xin|xperia|xxx|xyz|yachts|yahoo|yamaxun|yandex|ye|yodobashi|yoga|yokohama|you|youtube|yt|yun|za|zappos|zara|zero|zip|zippo|zm|zone|zuerich|zw|δοκιμή|ελ|бг|бел|дети|ею|испытание|католик|ком|мкд|мон|москва|онлайн|орг|рус|рф|сайт|срб|укр|қаз|հայ|טעסט|קום|آزمایشی|إختبار|ابوظبي|ارامكو|الاردن|الجزائر|السعودية|العليان|المغرب|امارات|ایران|بارت|بازار|بيتك|بھارت|تونس|سودان|سورية|شبكة|عراق|عمان|فلسطين|قطر|كاثوليك|كوم|مصر|مليسيا|موبايلي|موقع|همراه|پاكستان|پاکستان|ڀارت|कॉम|नेट|परीक्षा|भारत|भारतम्|भारोत|संगठन|বাংলা|ভারত|ভাৰত|ਭਾਰਤ|ભારત|ଭାରତ|இந்தியா|இலங்கை|சிங்கப்பூர்|பரிட்சை|భారత్|ಭಾರತ|ഭാരതം|ලංකා|คอม|ไทย|გე|みんな|クラウド|グーグル|コム|ストア|セール|テスト|ファッション|ポイント|世界|中信|中国|中國|中文网|企业|佛山|信息|健康|八卦|公司|公益|台湾|台灣|商城|商店|商标|嘉里|嘉里大酒店|在线|大众汽车|大拿|天主教|娱乐|家電|工行|广东|微博|慈善|我爱你|手机|手表|政务|政府|新加坡|新闻|时尚|書籍|机构|测试|淡马锡|測試|游戏|澳門|点看|珠宝|移动|组织机构|网址|网店|网站|网络|联通|诺基亚|谷歌|购物|通販|集团|電訊盈科|飞利浦|食品|餐厅|香格里拉|香港|닷넷|닷컴|삼성|테스트|한국)$`)
)

// Check defines functions needed to check every part of a routing.RouterConfig
type Check interface {
	ListenStatement(r *routing.ListenStatement) error
	ServerName(s pq.StringArray) error
	Path(p string) error
	Log(l *routing.Log) error
	SSLSettings(s *routing.SSLSettings) error
	LocationRule(l *routing.LocationRule) error
	LocationRules(l *routing.LocationRules) error
	Config(r *routing.RouterConfig, edit bool) error
}

type check struct {
	r Router
}

func (c *check) ListenStatement(r *routing.ListenStatement) error {
	return c.listenStatement(r, false)
}

func (c *check) listenStatement(r *routing.ListenStatement, edit bool) error {
	if r == nil {
		if edit {
			return nil
		}
		return ErrNoListenStatement
	}

	// TODO: get IP pool for check: if !pool.In(inet) return err
	if r.Port <= 1024 {
		return ErrPortRange
	}

	switch c.r {
	case Nginx:
		regex := regexp.MustCompile(`^ssl$|^$`)
		if !regex.MatchString(r.Keyword) {
			return ErrKeyword
		}
	default:
		if r.Keyword != "" {
			return ErrKeyword
		}
	}

	return nil
}

func (c *check) ServerName(s pq.StringArray) error {
	return c.serverName(s, false)
}

func (c *check) serverName(s pq.StringArray, edit bool) error {
	if len(s) == 0 {
		if edit {
			return nil
		}
		return ErrEmptyServerName
	}

	for _, n := range s {
		if !urlRegex.MatchString(n) {
			return ErrInvalidName
		}
	}
	return nil
}

func (c *check) Path(p string) error {
	return nil
}

func (c *check) Log(l *routing.Log) error {
	return nil
}

func (c *check) SSLSettings(s *routing.SSLSettings) error {
	return nil
}

func (c *check) LocationRule(l *routing.LocationRule) error {
	return nil
}

func (c *check) LocationRules(l *routing.LocationRules) error {
	return nil
}

func (c *check) Config(r *routing.RouterConfig, edit bool) error {
	var err error

	if r.RefID == 0 && !edit {
		return ErrNoRefID
	}

	if r.Name == "" && !edit {
		return ErrNoName
	}

	err = c.listenStatement(r.ListenStatement, edit)
	if err != nil {
		return err
	}

	err = c.serverName(r.ServerName, edit)
	if err != nil {
		return err
	}

	err = c.Log(&r.AccessLog)
	if err != nil {
		return err
	}

	err = c.Log(&r.ErrorLog)
	if err != nil {
		return err
	}

	err = c.Path(r.RootPath)
	if err != nil {
		return err
	}

	err = c.SSLSettings(&r.SSLSettings)
	if err != nil {
		return err
	}

	err = c.LocationRules(&r.LocationRules)
	if err != nil {
		return err
	}

	return nil
}

// NewCheck returns a new Check
func NewCheck(r Router) Check {
	return &check{r}
}
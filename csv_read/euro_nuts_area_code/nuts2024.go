package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"strings"
)

func main() {
	result := make(map[Country]map[string]string) // map[Country]map[NUTSCode]AreaDescription
	r := csv.NewReader(strings.NewReader(AreaCodeNutsData))
	r.Comma = ';'
	for i := 0; true; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error csv.Reader.Read: %v", err)
			continue
		}
		//log.Printf("debug row: %+v", row)
		if i == 0 {
			continue // the first row are fields name
		}
		if len(row) < 3 {
			log.Printf("unexpected row format: %v", err)
			continue
		}
		countryAlpha2 := strings.TrimSpace(row[0])
		if len(countryAlpha2) != 2 {
			log.Printf("unexpected countryAlpha2 from: %+v", row)
		}
		country := CountryCodesInverse[CountryAlpha2(countryAlpha2)]
		nutsCode := strings.TrimSpace(row[1])
		areaDescription := strings.TrimSpace(row[2])
		if country != "" && nutsCode != "" && areaDescription != "" {
			if result[country] == nil {
				result[country] = make(map[string]string)
			}
			result[country][nutsCode] = areaDescription
		}
	}
	beauty, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Fatalf("error json.MarshalIndent: %v", err)
	}
	log.Printf("resultCellToCarPlate:\n%s", beauty)
}

type Country string

const (
	Austria           Country = "Austria"
	Belgium           Country = "Belgium"
	BosniaHerzegovina Country = "BosniaHerzegovina" // Bosnia and Herzegovina
	Bulgaria          Country = "Bulgaria"
	Croatia           Country = "Croatia"
	Cyprus            Country = "Cyprus"
	Czechia           Country = "Czechia"
	Denmark           Country = "Denmark"
	Estonia           Country = "Estonia"
	Finland           Country = "Finland"
	France            Country = "France"
	Germany           Country = "Germany"
	Greece            Country = "Greece"
	Hungary           Country = "Hungary"
	Iceland           Country = "Iceland"
	Ireland           Country = "Ireland"
	Israel            Country = "Israel"
	Italy             Country = "Italy"
	Latvia            Country = "Latvia"
	Lithuania         Country = "Lithuania"
	Luxembourg        Country = "Luxembourg"
	Malta             Country = "Malta"
	Moldova           Country = "Moldova" // Moldova, Republic of
	Montenegro        Country = "Montenegro"
	Netherlands       Country = "Netherlands"    // Netherlands, Kingdom of the
	NorthMacedonia    Country = "NorthMacedonia" // North Macedonia
	Norway            Country = "Norway"
	Poland            Country = "Poland"
	Portugal          Country = "Portugal"
	Romania           Country = "Romania"
	Serbia            Country = "Serbia"
	Slovakia          Country = "Slovakia"
	Slovenia          Country = "Slovenia"
	Spain             Country = "Spain"
	Sweden            Country = "Sweden"
	Switzerland       Country = "Switzerland"
	Ukraine           Country = "Ukraine"
	UnitedKingdom     Country = "UnitedKingdom" // United Kingdom of Great Britain and Northern Ireland
)

type CountryAlpha2 string

// CountryCodes maps country name to country code alpha-2
// source: https://github.com/lukes/ISO-3166-Countries-with-Regional-Codes/blob/master/all/all.csv
var CountryCodes = map[Country]CountryAlpha2{
	Austria:           "AT",
	BosniaHerzegovina: "BA",
	Belgium:           "BE",
	Bulgaria:          "BG",
	Croatia:           "HR",
	Cyprus:            "CY",
	Czechia:           "CZ",
	Denmark:           "DK",
	Estonia:           "EE",
	Finland:           "FI",
	France:            "FR",
	Germany:           "DE",
	Greece:            "GR",
	Hungary:           "HU",
	Iceland:           "IS",
	Ireland:           "IE",
	Israel:            "IL",
	Italy:             "IT",
	Latvia:            "LV",
	Lithuania:         "LT",
	Luxembourg:        "LU",
	Malta:             "MT",
	Moldova:           "MD",
	Montenegro:        "ME",
	Netherlands:       "NL",
	NorthMacedonia:    "MK",
	Norway:            "NO",
	Poland:            "PL",
	Portugal:          "PT",
	Romania:           "RO",
	Serbia:            "RS",
	Slovakia:          "SK",
	Slovenia:          "SI",
	Spain:             "ES",
	Sweden:            "SE",
	Switzerland:       "CH",
	Ukraine:           "UA",
	UnitedKingdom:     "GB",
}
var CountryCodesInverse = map[CountryAlpha2]Country{}

func init() {
	for k, v := range CountryCodes {
		CountryCodesInverse[v] = k
	}
}

// areaCodeNutsData is from https://ec.europa.eu/eurostat/web/nuts/overview:
// NUTS 2024 classification: https://ec.europa.eu/eurostat/documents/345175/629341/NUTS2021-NUTS2024.xlsx/2b35915f-9c14-6841-8197-353408c4522d?t=1717505289640
const AreaCodeNutsData string = `Country code;NUTS Code;NUTS label;NUTS level;Country order;#
BE;BE1;Région de Bruxelles-Capitale/Brussels Hoofdstedelijk Gewest;1;1;1
BE;BE10;Région de Bruxelles-Capitale/Brussels Hoofdstedelijk Gewest;2;1;2
BE;BE100;Arr. de Bruxelles-Capitale/Arr. Brussel-Hoofdstad;3;1;3
BE;BE2;Vlaams Gewest;1;1;4
BE;BE21;Prov. Antwerpen;2;1;5
BE;BE211;Arr. Antwerpen;3;1;6
BE;BE212;Arr. Mechelen;3;1;7
BE;BE213;Arr. Turnhout;3;1;8
BE;BE22;Prov. Limburg (BE);2;1;9
BE;BE223;Arr. Tongeren;3;1;10
BE;BE224;Arr. Hasselt;3;1;11
BE;BE225;Arr. Maaseik;3;1;12
BE;BE23;Prov. Oost-Vlaanderen;2;1;13
BE;BE231;Arr. Aalst;3;1;14
BE;BE232;Arr. Dendermonde;3;1;15
BE;BE233;Arr. Eeklo;3;1;16
BE;BE234;Arr. Gent;3;1;17
BE;BE235;Arr. Oudenaarde;3;1;18
BE;BE236;Arr. Sint-Niklaas;3;1;19
BE;BE24;Prov. Vlaams-Brabant;2;1;20
BE;BE241;Arr. Halle-Vilvoorde;3;1;21
BE;BE242;Arr. Leuven;3;1;22
BE;BE25;Prov. West-Vlaanderen;2;1;23
BE;BE251;Arr. Brugge;3;1;24
BE;BE252;Arr. Diksmuide;3;1;25
BE;BE253;Arr. Ieper;3;1;26
BE;BE254;Arr. Kortrijk;3;1;27
BE;BE255;Arr. Oostende;3;1;28
BE;BE256;Arr. Roeselare;3;1;29
BE;BE257;Arr. Tielt;3;1;30
BE;BE258;Arr. Veurne;3;1;31
BE;BE3;Région wallonne;1;1;32
BE;BE31;Prov. Brabant wallon;2;1;33
BE;BE310;Arr. Nivelles;3;1;34
BE;BE32;Prov. Hainaut;2;1;35
BE;BE323;Arr. Mons;3;1;36
BE;BE328;Arr. Tournai-Mouscron;3;1;37
BE;BE329;Arr. La Louvière;3;1;38
BE;BE32A;Arr. Ath;3;1;39
BE;BE32B;Arr. Charleroi;3;1;40
BE;BE32C;Arr. Soignies;3;1;41
BE;BE32D;Arr. Thuin;3;1;42
BE;BE33;Prov. Liège;2;1;43
BE;BE331;Arr. Huy;3;1;44
BE;BE332;Arr. Liège;3;1;45
BE;BE334;Arr. Waremme;3;1;46
BE;BE335;Arr. Verviers — communes francophones;3;1;47
BE;BE336;Bezirk Verviers — Deutschsprachige Gemeinschaft;3;1;48
BE;BE34;Prov. Luxembourg (BE);2;1;49
BE;BE341;Arr. Arlon;3;1;50
BE;BE342;Arr. Bastogne;3;1;51
BE;BE343;Arr. Marche-en-Famenne;3;1;52
BE;BE344;Arr. Neufchâteau;3;1;53
BE;BE345;Arr. Virton;3;1;54
BE;BE35;Prov. Namur;2;1;55
BE;BE351;Arr. Dinant;3;1;56
BE;BE352;Arr. Namur;3;1;57
BE;BE353;Arr. Philippeville;3;1;58
BE;BEZ;Extra-Regio NUTS 1;1;1;59
BE;BEZZ;Extra-Regio NUTS 2;2;1;60
BE;BEZZZ;Extra-Regio NUTS 3;3;1;61
BG;BG3;Северна и Югоизточна България;1;2;62
BG;BG31;Северозападен;2;2;63
BG;BG311;Видин;3;2;64
BG;BG312;Монтана;3;2;65
BG;BG313;Враца;3;2;66
BG;BG314;Плевен;3;2;67
BG;BG315;Ловеч;3;2;68
BG;BG32;Северен централен;2;2;69
BG;BG321;Велико Търново;3;2;70
BG;BG322;Габрово;3;2;71
BG;BG323;Русе;3;2;72
BG;BG324;Разград;3;2;73
BG;BG325;Силистра;3;2;74
BG;BG33;Североизточен;2;2;75
BG;BG331;Варна;3;2;76
BG;BG332;Добрич;3;2;77
BG;BG333;Шумен;3;2;78
BG;BG334;Търговище;3;2;79
BG;BG34;Югоизточен;2;2;80
BG;BG341;Бургас;3;2;81
BG;BG342;Сливен;3;2;82
BG;BG343;Ямбол;3;2;83
BG;BG344;Стара Загора;3;2;84
BG;BG4;Югозападна и Южна централна България;1;2;85
BG;BG41;Югозападен;2;2;86
BG;BG411;София (столица);3;2;87
BG;BG412;София;3;2;88
BG;BG413;Благоевград;3;2;89
BG;BG414;Перник;3;2;90
BG;BG415;Кюстендил;3;2;91
BG;BG42;Южен централен;2;2;92
BG;BG421;Пловдив;3;2;93
BG;BG422;Хасково;3;2;94
BG;BG423;Пазарджик;3;2;95
BG;BG424;Смолян;3;2;96
BG;BG425;Кърджали;3;2;97
BG;BGZ;Extra-Regio NUTS 1;1;2;98
BG;BGZZ;Extra-Regio NUTS 2;2;2;99
BG;BGZZZ;Extra-Regio NUTS 3;3;2;100
CZ;CZ0;Česko;1;3;101
CZ;CZ01;Praha;2;3;102
CZ;CZ010;Hlavní město Praha;3;3;103
CZ;CZ02;Střední Čechy;2;3;104
CZ;CZ020;Středočeský kraj;3;3;105
CZ;CZ03;Jihozápad;2;3;106
CZ;CZ031;Jihočeský kraj;3;3;107
CZ;CZ032;Plzeňský kraj;3;3;108
CZ;CZ04;Severozápad;2;3;109
CZ;CZ041;Karlovarský kraj;3;3;110
CZ;CZ042;Ústecký kraj;3;3;111
CZ;CZ05;Severovýchod;2;3;112
CZ;CZ051;Liberecký kraj;3;3;113
CZ;CZ052;Královéhradecký kraj;3;3;114
CZ;CZ053;Pardubický kraj;3;3;115
CZ;CZ06;Jihovýchod;2;3;116
CZ;CZ063;Kraj Vysočina;3;3;117
CZ;CZ064;Jihomoravský kraj;3;3;118
CZ;CZ07;Střední Morava;2;3;119
CZ;CZ071;Olomoucký kraj;3;3;120
CZ;CZ072;Zlínský kraj;3;3;121
CZ;CZ08;Moravskoslezsko;2;3;122
CZ;CZ080;Moravskoslezský kraj;3;3;123
CZ;CZZ;Extra-Regio NUTS 1;1;3;124
CZ;CZZZ;Extra-Regio NUTS 2;2;3;125
CZ;CZZZZ;Extra-Regio NUTS 3;3;3;126
DK;DK0;Danmark;1;4;127
DK;DK01;Hovedstaden;2;4;128
DK;DK011;Byen København;3;4;129
DK;DK012;Københavns omegn;3;4;130
DK;DK013;Nordsjælland;3;4;131
DK;DK014;Bornholm;3;4;132
DK;DK02;Sjælland;2;4;133
DK;DK021;Østsjælland;3;4;134
DK;DK022;Vest- og Sydsjælland;3;4;135
DK;DK03;Syddanmark;2;4;136
DK;DK031;Fyn;3;4;137
DK;DK032;Sydjylland;3;4;138
DK;DK04;Midtjylland;2;4;139
DK;DK041;Vestjylland;3;4;140
DK;DK042;Østjylland;3;4;141
DK;DK05;Nordjylland;2;4;142
DK;DK050;Nordjylland;3;4;143
DK;DKZ;Extra-Regio NUTS 1;1;4;144
DK;DKZZ;Extra-Regio NUTS 2;2;4;145
DK;DKZZZ;Extra-Regio NUTS 3;3;4;146
DE;DE1;Baden-Württemberg;1;5;147
DE;DE11;Stuttgart;2;5;148
DE;DE111;Stuttgart, Stadtkreis;3;5;149
DE;DE112;Böblingen;3;5;150
DE;DE113;Esslingen;3;5;151
DE;DE114;Göppingen;3;5;152
DE;DE115;Ludwigsburg;3;5;153
DE;DE116;Rems-Murr-Kreis;3;5;154
DE;DE117;Heilbronn, Stadtkreis;3;5;155
DE;DE118;Heilbronn, Landkreis;3;5;156
DE;DE119;Hohenlohekreis;3;5;157
DE;DE11A;Schwäbisch Hall;3;5;158
DE;DE11B;Main-Tauber-Kreis;3;5;159
DE;DE11C;Heidenheim;3;5;160
DE;DE11D;Ostalbkreis;3;5;161
DE;DE12;Karlsruhe;2;5;162
DE;DE121;Baden-Baden, Stadtkreis;3;5;163
DE;DE122;Karlsruhe, Stadtkreis;3;5;164
DE;DE123;Karlsruhe, Landkreis;3;5;165
DE;DE124;Rastatt;3;5;166
DE;DE125;Heidelberg, Stadtkreis;3;5;167
DE;DE126;Mannheim, Stadtkreis;3;5;168
DE;DE127;Neckar-Odenwald-Kreis;3;5;169
DE;DE128;Rhein-Neckar-Kreis;3;5;170
DE;DE129;Pforzheim, Stadtkreis;3;5;171
DE;DE12A;Calw;3;5;172
DE;DE12B;Enzkreis;3;5;173
DE;DE12C;Freudenstadt;3;5;174
DE;DE13;Freiburg;2;5;175
DE;DE131;Freiburg im Breisgau, Stadtkreis;3;5;176
DE;DE132;Breisgau-Hochschwarzwald;3;5;177
DE;DE133;Emmendingen;3;5;178
DE;DE134;Ortenaukreis;3;5;179
DE;DE135;Rottweil;3;5;180
DE;DE136;Schwarzwald-Baar-Kreis;3;5;181
DE;DE137;Tuttlingen;3;5;182
DE;DE138;Konstanz;3;5;183
DE;DE139;Lörrach;3;5;184
DE;DE13A;Waldshut;3;5;185
DE;DE14;Tübingen;2;5;186
DE;DE141;Reutlingen;3;5;187
DE;DE142;Tübingen, Landkreis;3;5;188
DE;DE143;Zollernalbkreis;3;5;189
DE;DE144;Ulm, Stadtkreis;3;5;190
DE;DE145;Alb-Donau-Kreis;3;5;191
DE;DE146;Biberach;3;5;192
DE;DE147;Bodenseekreis;3;5;193
DE;DE148;Ravensburg;3;5;194
DE;DE149;Sigmaringen;3;5;195
DE;DE2;Bayern;1;5;196
DE;DE21;Oberbayern;2;5;197
DE;DE211;Ingolstadt, Kreisfreie Stadt;3;5;198
DE;DE212;München, Kreisfreie Stadt;3;5;199
DE;DE213;Rosenheim, Kreisfreie Stadt;3;5;200
DE;DE214;Altötting;3;5;201
DE;DE215;Berchtesgadener Land;3;5;202
DE;DE216;Bad Tölz-Wolfratshausen;3;5;203
DE;DE217;Dachau;3;5;204
DE;DE218;Ebersberg;3;5;205
DE;DE219;Eichstätt;3;5;206
DE;DE21A;Erding;3;5;207
DE;DE21B;Freising;3;5;208
DE;DE21C;Fürstenfeldbruck;3;5;209
DE;DE21D;Garmisch-Partenkirchen;3;5;210
DE;DE21E;Landsberg am Lech;3;5;211
DE;DE21F;Miesbach;3;5;212
DE;DE21G;Mühldorf a. Inn;3;5;213
DE;DE21H;München, Landkreis;3;5;214
DE;DE21I;Neuburg-Schrobenhausen;3;5;215
DE;DE21J;Pfaffenhofen a. d. Ilm;3;5;216
DE;DE21K;Rosenheim, Landkreis;3;5;217
DE;DE21L;Starnberg;3;5;218
DE;DE21M;Traunstein;3;5;219
DE;DE21N;Weilheim-Schongau;3;5;220
DE;DE22;Niederbayern;2;5;221
DE;DE221;Landshut, Kreisfreie Stadt;3;5;222
DE;DE222;Passau, Kreisfreie Stadt;3;5;223
DE;DE223;Straubing, Kreisfreie Stadt;3;5;224
DE;DE224;Deggendorf;3;5;225
DE;DE225;Freyung-Grafenau;3;5;226
DE;DE226;Kelheim;3;5;227
DE;DE227;Landshut, Landkreis;3;5;228
DE;DE228;Passau, Landkreis;3;5;229
DE;DE229;Regen;3;5;230
DE;DE22A;Rottal-Inn;3;5;231
DE;DE22B;Straubing-Bogen;3;5;232
DE;DE22C;Dingolfing-Landau;3;5;233
DE;DE23;Oberpfalz;2;5;234
DE;DE231;Amberg, Kreisfreie Stadt;3;5;235
DE;DE232;Regensburg, Kreisfreie Stadt;3;5;236
DE;DE233;Weiden i. d. Opf, Kreisfreie Stadt;3;5;237
DE;DE234;Amberg-Sulzbach;3;5;238
DE;DE235;Cham;3;5;239
DE;DE236;Neumarkt i. d. OPf.;3;5;240
DE;DE237;Neustadt a. d. Waldnaab;3;5;241
DE;DE238;Regensburg, Landkreis;3;5;242
DE;DE239;Schwandorf;3;5;243
DE;DE23A;Tirschenreuth;3;5;244
DE;DE24;Oberfranken;2;5;245
DE;DE241;Bamberg, Kreisfreie Stadt;3;5;246
DE;DE242;Bayreuth, Kreisfreie Stadt;3;5;247
DE;DE243;Coburg, Kreisfreie Stadt;3;5;248
DE;DE244;Hof, Kreisfreie Stadt;3;5;249
DE;DE245;Bamberg, Landkreis;3;5;250
DE;DE246;Bayreuth, Landkreis;3;5;251
DE;DE247;Coburg, Landkreis;3;5;252
DE;DE248;Forchheim;3;5;253
DE;DE249;Hof, Landkreis;3;5;254
DE;DE24A;Kronach;3;5;255
DE;DE24B;Kulmbach;3;5;256
DE;DE24C;Lichtenfels;3;5;257
DE;DE24D;Wunsiedel i. Fichtelgebirge;3;5;258
DE;DE25;Mittelfranken;2;5;259
DE;DE251;Ansbach, Kreisfreie Stadt;3;5;260
DE;DE252;Erlangen, Kreisfreie Stadt;3;5;261
DE;DE253;Fürth, Kreisfreie Stadt;3;5;262
DE;DE254;Nürnberg, Kreisfreie Stadt;3;5;263
DE;DE255;Schwabach, Kreisfreie Stadt;3;5;264
DE;DE256;Ansbach, Landkreis;3;5;265
DE;DE257;Erlangen-Höchstadt;3;5;266
DE;DE258;Fürth, Landkreis;3;5;267
DE;DE259;Nürnberger Land;3;5;268
DE;DE25A;Neustadt a. d. Aisch-Bad Windsheim;3;5;269
DE;DE25B;Roth;3;5;270
DE;DE25C;Weißenburg-Gunzenhausen;3;5;271
DE;DE26;Unterfranken;2;5;272
DE;DE261;Aschaffenburg, Kreisfreie Stadt;3;5;273
DE;DE262;Schweinfurt, Kreisfreie Stadt;3;5;274
DE;DE263;Würzburg, Kreisfreie Stadt;3;5;275
DE;DE264;Aschaffenburg, Landkreis;3;5;276
DE;DE265;Bad Kissingen;3;5;277
DE;DE266;Rhön-Grabfeld;3;5;278
DE;DE267;Haßberge;3;5;279
DE;DE268;Kitzingen;3;5;280
DE;DE269;Miltenberg;3;5;281
DE;DE26A;Main-Spessart;3;5;282
DE;DE26B;Schweinfurt, Landkreis;3;5;283
DE;DE26C;Würzburg, Landkreis;3;5;284
DE;DE27;Schwaben;2;5;285
DE;DE271;Augsburg, Kreisfreie Stadt;3;5;286
DE;DE272;Kaufbeuren, Kreisfreie Stadt;3;5;287
DE;DE273;Kempten (Allgäu), Kreisfreie Stadt;3;5;288
DE;DE274;Memmingen, Kreisfreie Stadt;3;5;289
DE;DE275;Aichach-Friedberg;3;5;290
DE;DE276;Augsburg, Landkreis;3;5;291
DE;DE277;Dillingen a.d. Donau;3;5;292
DE;DE278;Günzburg;3;5;293
DE;DE279;Neu-Ulm;3;5;294
DE;DE27A;Lindau (Bodensee);3;5;295
DE;DE27B;Ostallgäu;3;5;296
DE;DE27C;Unterallgäu;3;5;297
DE;DE27D;Donau-Ries;3;5;298
DE;DE27E;Oberallgäu;3;5;299
DE;DE3;Berlin;1;5;300
DE;DE30;Berlin;2;5;301
DE;DE300;Berlin;3;5;302
DE;DE4;Brandenburg;1;5;303
DE;DE40;Brandenburg;2;5;304
DE;DE401;Brandenburg an der Havel, Kreisfreie Stadt;3;5;305
DE;DE402;Cottbus, Kreisfreie Stadt;3;5;306
DE;DE403;Frankfurt (Oder), Kreisfreie Stadt;3;5;307
DE;DE404;Potsdam, Kreisfreie Stadt;3;5;308
DE;DE405;Barnim;3;5;309
DE;DE406;Dahme-Spreewald;3;5;310
DE;DE407;Elbe-Elster;3;5;311
DE;DE408;Havelland;3;5;312
DE;DE409;Märkisch-Oderland;3;5;313
DE;DE40A;Oberhavel;3;5;314
DE;DE40B;Oberspreewald-Lausitz;3;5;315
DE;DE40C;Oder-Spree;3;5;316
DE;DE40D;Ostprignitz-Ruppin;3;5;317
DE;DE40E;Potsdam-Mittelmark;3;5;318
DE;DE40F;Prignitz;3;5;319
DE;DE40G;Spree-Neiße;3;5;320
DE;DE40H;Teltow-Fläming;3;5;321
DE;DE40I;Uckermark;3;5;322
DE;DE5;Bremen;1;5;323
DE;DE50;Bremen;2;5;324
DE;DE501;Bremen, Kreisfreie Stadt;3;5;325
DE;DE502;Bremerhaven, Kreisfreie Stadt;3;5;326
DE;DE6;Hamburg;1;5;327
DE;DE60;Hamburg;2;5;328
DE;DE600;Hamburg;3;5;329
DE;DE7;Hessen;1;5;330
DE;DE71;Darmstadt;2;5;331
DE;DE711;Darmstadt, Kreisfreie Stadt;3;5;332
DE;DE712;Frankfurt am Main, Kreisfreie Stadt;3;5;333
DE;DE713;Offenbach am Main, Kreisfreie Stadt;3;5;334
DE;DE714;Wiesbaden, Kreisfreie Stadt;3;5;335
DE;DE715;Bergstraße;3;5;336
DE;DE716;Darmstadt-Dieburg;3;5;337
DE;DE717;Groß-Gerau;3;5;338
DE;DE718;Hochtaunuskreis;3;5;339
DE;DE719;Main-Kinzig-Kreis;3;5;340
DE;DE71A;Main-Taunus-Kreis;3;5;341
DE;DE71B;Odenwaldkreis;3;5;342
DE;DE71C;Offenbach, Landkreis;3;5;343
DE;DE71D;Rheingau-Taunus-Kreis;3;5;344
DE;DE71E;Wetteraukreis;3;5;345
DE;DE72;Gießen;2;5;346
DE;DE721;Gießen, Landkreis;3;5;347
DE;DE722;Lahn-Dill-Kreis;3;5;348
DE;DE723;Limburg-Weilburg;3;5;349
DE;DE724;Marburg-Biedenkopf;3;5;350
DE;DE725;Vogelsbergkreis;3;5;351
DE;DE73;Kassel;2;5;352
DE;DE731;Kassel, Kreisfreie Stadt;3;5;353
DE;DE732;Fulda;3;5;354
DE;DE733;Hersfeld-Rotenburg;3;5;355
DE;DE734;Kassel, Landkreis;3;5;356
DE;DE735;Schwalm-Eder-Kreis;3;5;357
DE;DE736;Waldeck-Frankenberg;3;5;358
DE;DE737;Werra-Meißner-Kreis;3;5;359
DE;DE8;Mecklenburg-Vorpommern;1;5;360
DE;DE80;Mecklenburg-Vorpommern;2;5;361
DE;DE803;Rostock, Kreisfreie Stadt;3;5;362
DE;DE804;Schwerin, Kreisfreie Stadt;3;5;363
DE;DE80J;Mecklenburgische Seenplatte;3;5;364
DE;DE80K;Landkreis Rostock;3;5;365
DE;DE80L;Vorpommern-Rügen;3;5;366
DE;DE80M;Nordwestmecklenburg;3;5;367
DE;DE80N;Vorpommern-Greifswald;3;5;368
DE;DE80O;Ludwigslust-Parchim;3;5;369
DE;DE9;Niedersachsen;1;5;370
DE;DE91;Braunschweig;2;5;371
DE;DE911;Braunschweig, Kreisfreie Stadt;3;5;372
DE;DE912;Salzgitter, Kreisfreie Stadt;3;5;373
DE;DE913;Wolfsburg, Kreisfreie Stadt;3;5;374
DE;DE914;Gifhorn;3;5;375
DE;DE916;Goslar;3;5;376
DE;DE917;Helmstedt;3;5;377
DE;DE918;Northeim;3;5;378
DE;DE91A;Peine;3;5;379
DE;DE91B;Wolfenbüttel;3;5;380
DE;DE91C;Göttingen;3;5;381
DE;DE92;Hannover;2;5;382
DE;DE922;Diepholz;3;5;383
DE;DE923;Hameln-Pyrmont;3;5;384
DE;DE925;Hildesheim;3;5;385
DE;DE926;Holzminden;3;5;386
DE;DE927;Nienburg (Weser);3;5;387
DE;DE928;Schaumburg;3;5;388
DE;DE929;Region Hannover;3;5;389
DE;DE93;Lüneburg;2;5;390
DE;DE931;Celle;3;5;391
DE;DE932;Cuxhaven;3;5;392
DE;DE933;Harburg;3;5;393
DE;DE934;Lüchow-Dannenberg;3;5;394
DE;DE935;Lüneburg, Landkreis;3;5;395
DE;DE936;Osterholz;3;5;396
DE;DE937;Rotenburg (Wümme);3;5;397
DE;DE938;Heidekreis;3;5;398
DE;DE939;Stade;3;5;399
DE;DE93A;Uelzen;3;5;400
DE;DE93B;Verden;3;5;401
DE;DE94;Weser-Ems;2;5;402
DE;DE941;Delmenhorst, Kreisfreie Stadt;3;5;403
DE;DE942;Emden, Kreisfreie Stadt;3;5;404
DE;DE943;Oldenburg (Oldenburg), Kreisfreie Stadt;3;5;405
DE;DE944;Osnabrück, Kreisfreie Stadt;3;5;406
DE;DE945;Wilhelmshaven, Kreisfreie Stadt;3;5;407
DE;DE946;Ammerland;3;5;408
DE;DE947;Aurich;3;5;409
DE;DE948;Cloppenburg;3;5;410
DE;DE949;Emsland;3;5;411
DE;DE94A;Friesland (DE);3;5;412
DE;DE94B;Grafschaft Bentheim;3;5;413
DE;DE94C;Leer;3;5;414
DE;DE94D;Oldenburg, Landkreis;3;5;415
DE;DE94E;Osnabrück, Landkreis;3;5;416
DE;DE94F;Vechta;3;5;417
DE;DE94G;Wesermarsch;3;5;418
DE;DE94H;Wittmund;3;5;419
DE;DEA;Nordrhein-Westfalen;1;5;420
DE;DEA1;Düsseldorf;2;5;421
DE;DEA11;Düsseldorf, Kreisfreie Stadt;3;5;422
DE;DEA12;Duisburg, Kreisfreie Stadt;3;5;423
DE;DEA13;Essen, Kreisfreie Stadt;3;5;424
DE;DEA14;Krefeld, Kreisfreie Stadt;3;5;425
DE;DEA15;Mönchengladbach, Kreisfreie Stadt;3;5;426
DE;DEA16;Mülheim an der Ruhr, Kreisfreie Stadt;3;5;427
DE;DEA17;Oberhausen, Kreisfreie Stadt;3;5;428
DE;DEA18;Remscheid, Kreisfreie Stadt;3;5;429
DE;DEA19;Solingen, Kreisfreie Stadt;3;5;430
DE;DEA1A;Wuppertal, Kreisfreie Stadt;3;5;431
DE;DEA1B;Kleve;3;5;432
DE;DEA1C;Mettmann;3;5;433
DE;DEA1D;Rhein-Kreis Neuss;3;5;434
DE;DEA1E;Viersen;3;5;435
DE;DEA1F;Wesel;3;5;436
DE;DEA2;Köln;2;5;437
DE;DEA22;Bonn, Kreisfreie Stadt;3;5;438
DE;DEA23;Köln, Kreisfreie Stadt;3;5;439
DE;DEA24;Leverkusen, Kreisfreie Stadt;3;5;440
DE;DEA26;Düren;3;5;441
DE;DEA27;Rhein-Erft-Kreis;3;5;442
DE;DEA28;Euskirchen;3;5;443
DE;DEA29;Heinsberg;3;5;444
DE;DEA2A;Oberbergischer Kreis;3;5;445
DE;DEA2B;Rheinisch-Bergischer Kreis;3;5;446
DE;DEA2C;Rhein-Sieg-Kreis;3;5;447
DE;DEA2D;Städteregion Aachen;3;5;448
DE;DEA3;Münster;2;5;449
DE;DEA31;Bottrop, Kreisfreie Stadt;3;5;450
DE;DEA32;Gelsenkirchen, Kreisfreie Stadt;3;5;451
DE;DEA33;Münster, Kreisfreie Stadt;3;5;452
DE;DEA34;Borken;3;5;453
DE;DEA35;Coesfeld;3;5;454
DE;DEA36;Recklinghausen;3;5;455
DE;DEA37;Steinfurt;3;5;456
DE;DEA38;Warendorf;3;5;457
DE;DEA4;Detmold;2;5;458
DE;DEA41;Bielefeld, Kreisfreie Stadt;3;5;459
DE;DEA42;Gütersloh;3;5;460
DE;DEA43;Herford;3;5;461
DE;DEA44;Höxter;3;5;462
DE;DEA45;Lippe;3;5;463
DE;DEA46;Minden-Lübbecke;3;5;464
DE;DEA47;Paderborn;3;5;465
DE;DEA5;Arnsberg;2;5;466
DE;DEA51;Bochum, Kreisfreie Stadt;3;5;467
DE;DEA52;Dortmund, Kreisfreie Stadt;3;5;468
DE;DEA53;Hagen, Kreisfreie Stadt;3;5;469
DE;DEA54;Hamm, Kreisfreie Stadt;3;5;470
DE;DEA55;Herne, Kreisfreie Stadt;3;5;471
DE;DEA56;Ennepe-Ruhr-Kreis;3;5;472
DE;DEA57;Hochsauerlandkreis;3;5;473
DE;DEA58;Märkischer Kreis;3;5;474
DE;DEA59;Olpe;3;5;475
DE;DEA5A;Siegen-Wittgenstein;3;5;476
DE;DEA5B;Soest;3;5;477
DE;DEA5C;Unna;3;5;478
DE;DEB;Rheinland-Pfalz;1;5;479
DE;DEB1;Koblenz;2;5;480
DE;DEB11;Koblenz, Kreisfreie Stadt;3;5;481
DE;DEB12;Ahrweiler;3;5;482
DE;DEB13;Altenkirchen (Westerwald);3;5;483
DE;DEB14;Bad Kreuznach;3;5;484
DE;DEB15;Birkenfeld;3;5;485
DE;DEB17;Mayen-Koblenz;3;5;486
DE;DEB18;Neuwied;3;5;487
DE;DEB1A;Rhein-Lahn-Kreis;3;5;488
DE;DEB1B;Westerwaldkreis;3;5;489
DE;DEB1C;Cochem-Zell;3;5;490
DE;DEB1D;Rhein-Hunsrück-Kreis;3;5;491
DE;DEB2;Trier;2;5;492
DE;DEB21;Trier, Kreisfreie Stadt;3;5;493
DE;DEB22;Bernkastel-Wittlich;3;5;494
DE;DEB23;Eifelkreis Bitburg-Prüm;3;5;495
DE;DEB24;Vulkaneifel;3;5;496
DE;DEB25;Trier-Saarburg;3;5;497
DE;DEB3;Rheinhessen-Pfalz;2;5;498
DE;DEB31;Frankenthal (Pfalz), Kreisfreie Stadt;3;5;499
DE;DEB32;Kaiserslautern, Kreisfreie Stadt;3;5;500
DE;DEB33;Landau in der Pfalz, Kreisfreie Stadt;3;5;501
DE;DEB34;Ludwigshafen am Rhein, Kreisfreie Stadt;3;5;502
DE;DEB35;Mainz, Kreisfreie Stadt;3;5;503
DE;DEB36;Neustadt an der Weinstraße, Kreisfreie Stadt;3;5;504
DE;DEB37;Pirmasens, Kreisfreie Stadt;3;5;505
DE;DEB38;Speyer, Kreisfreie Stadt;3;5;506
DE;DEB39;Worms, Kreisfreie Stadt;3;5;507
DE;DEB3A;Zweibrücken, Kreisfreie Stadt;3;5;508
DE;DEB3B;Alzey-Worms;3;5;509
DE;DEB3C;Bad Dürkheim;3;5;510
DE;DEB3D;Donnersbergkreis;3;5;511
DE;DEB3E;Germersheim;3;5;512
DE;DEB3F;Kaiserslautern, Landkreis;3;5;513
DE;DEB3G;Kusel;3;5;514
DE;DEB3H;Südliche Weinstraße;3;5;515
DE;DEB3I;Rhein-Pfalz-Kreis;3;5;516
DE;DEB3J;Mainz-Bingen;3;5;517
DE;DEB3K;Südwestpfalz;3;5;518
DE;DEC;Saarland;1;5;519
DE;DEC0;Saarland;2;5;520
DE;DEC01;Regionalverband Saarbrücken;3;5;521
DE;DEC02;Merzig-Wadern;3;5;522
DE;DEC03;Neunkirchen;3;5;523
DE;DEC04;Saarlouis;3;5;524
DE;DEC05;Saarpfalz-Kreis;3;5;525
DE;DEC06;St. Wendel;3;5;526
DE;DED;Sachsen;1;5;527
DE;DED2;Dresden;2;5;528
DE;DED21;Dresden, Kreisfreie Stadt;3;5;529
DE;DED2C;Bautzen;3;5;530
DE;DED2D;Görlitz;3;5;531
DE;DED2E;Meißen;3;5;532
DE;DED2F;Sächsische Schweiz-Osterzgebirge;3;5;533
DE;DED4;Chemnitz;2;5;534
DE;DED41;Chemnitz, Kreisfreie Stadt;3;5;535
DE;DED42;Erzgebirgskreis;3;5;536
DE;DED43;Mittelsachsen;3;5;537
DE;DED44;Vogtlandkreis;3;5;538
DE;DED45;Zwickau;3;5;539
DE;DED5;Leipzig;2;5;540
DE;DED51;Leipzig, Kreisfreie Stadt;3;5;541
DE;DED52;Leipzig;3;5;542
DE;DED53;Nordsachsen;3;5;543
DE;DEE;Sachsen-Anhalt;1;5;544
DE;DEE0;Sachsen-Anhalt;2;5;545
DE;DEE01;Dessau-Roßlau, Kreisfreie Stadt;3;5;546
DE;DEE02;Halle (Saale), Kreisfreie Stadt;3;5;547
DE;DEE03;Magdeburg, Kreisfreie Stadt;3;5;548
DE;DEE04;Altmarkkreis Salzwedel;3;5;549
DE;DEE05;Anhalt-Bitterfeld;3;5;550
DE;DEE06;Jerichower Land;3;5;551
DE;DEE07;Börde;3;5;552
DE;DEE08;Burgenlandkreis;3;5;553
DE;DEE09;Harz;3;5;554
DE;DEE0A;Mansfeld-Südharz;3;5;555
DE;DEE0B;Saalekreis;3;5;556
DE;DEE0C;Salzlandkreis;3;5;557
DE;DEE0D;Stendal;3;5;558
DE;DEE0E;Wittenberg;3;5;559
DE;DEF;Schleswig-Holstein;1;5;560
DE;DEF0;Schleswig-Holstein;2;5;561
DE;DEF01;Flensburg, Kreisfreie Stadt;3;5;562
DE;DEF02;Kiel, Kreisfreie Stadt;3;5;563
DE;DEF03;Lübeck, Kreisfreie Stadt;3;5;564
DE;DEF04;Neumünster, Kreisfreie Stadt;3;5;565
DE;DEF05;Dithmarschen;3;5;566
DE;DEF06;Herzogtum Lauenburg;3;5;567
DE;DEF07;Nordfriesland;3;5;568
DE;DEF08;Ostholstein;3;5;569
DE;DEF09;Pinneberg;3;5;570
DE;DEF0A;Plön;3;5;571
DE;DEF0B;Rendsburg-Eckernförde;3;5;572
DE;DEF0C;Schleswig-Flensburg;3;5;573
DE;DEF0D;Segeberg;3;5;574
DE;DEF0E;Steinburg;3;5;575
DE;DEF0F;Stormarn;3;5;576
DE;DEG;Thüringen;1;5;577
DE;DEG0;Thüringen;2;5;578
DE;DEG01;Erfurt, Kreisfreie Stadt;3;5;579
DE;DEG02;Gera, Kreisfreie Stadt;3;5;580
DE;DEG03;Jena, Kreisfreie Stadt;3;5;581
DE;DEG05;Weimar, Kreisfreie Stadt;3;5;582
DE;DEG06;Eichsfeld;3;5;583
DE;DEG07;Nordhausen;3;5;584
DE;DEG09;Unstrut-Hainich-Kreis;3;5;585
DE;DEG0A;Kyffhäuserkreis;3;5;586
DE;DEG0C;Gotha;3;5;587
DE;DEG0D;Sömmerda;3;5;588
DE;DEG0E;Hildburghausen;3;5;589
DE;DEG0G;Weimarer Land;3;5;590
DE;DEG0J;Saale-Holzland-Kreis;3;5;591
DE;DEG0K;Saale-Orla-Kreis;3;5;592
DE;DEG0L;Greiz;3;5;593
DE;DEG0M;Altenburger Land;3;5;594
DE;DEG0Q;Schmalkalden-Meiningen;3;5;595
DE;DEG0R;Wartburgkreis;3;5;596
DE;DEG0S;Suhl, Kreisfreie Stadt;3;5;597
DE;DEG0T;Ilm-Kreis;3;5;598
DE;DEG0U;Saalfeld-Rudolstadt;3;5;599
DE;DEG0V;Sonneberg;3;5;600
DE;DEZ;Extra-Regio NUTS 1;1;5;601
DE;DEZZ;Extra-Regio NUTS 2;2;5;602
DE;DEZZZ;Extra-Regio NUTS 3;3;5;603
EE;EE0;Eesti;1;6;604
EE;EE00;Eesti;2;6;605
EE;EE001;Põhja-Eesti;3;6;606
EE;EE004;Lääne-Eesti;3;6;607
EE;EE008;Lõuna-Eesti;3;6;608
EE;EE009;Kesk-Eesti;3;6;609
EE;EE00A;Kirde-Eesti;3;6;610
EE;EEZ;Extra-Regio NUTS 1;1;6;611
EE;EEZZ;Extra-Regio NUTS 2;2;6;612
EE;EEZZZ;Extra-Regio NUTS 3;3;6;613
IE;IE0;Ireland;1;7;614
IE;IE04;Northern and Western;2;7;615
IE;IE041;Border;3;7;616
IE;IE042;West;3;7;617
IE;IE05;Southern;2;7;618
IE;IE051;Mid-West ;3;7;619
IE;IE052;South-East ;3;7;620
IE;IE053;South-West ;3;7;621
IE;IE06;Eastern and Midland;2;7;622
IE;IE061;Dublin;3;7;623
IE;IE062;Mid-East;3;7;624
IE;IE063;Midland;3;7;625
IE;IEZ;Extra-Regio NUTS 1;1;7;626
IE;IEZZ;Extra-Regio NUTS 2;2;7;627
IE;IEZZZ;Extra-Regio NUTS 3;3;7;628
EL;EL3;Αττική ;1;8;629
EL;EL30;Aττική;2;8;630
EL;EL301;Βόρειος Τομέας Αθηνών;3;8;631
EL;EL302;Δυτικός Τομέας Αθηνών;3;8;632
EL;EL303;Κεντρικός Τομέας Αθηνών;3;8;633
EL;EL304;Νότιος Τομέας Αθηνών;3;8;634
EL;EL305;Ανατολική Αττική;3;8;635
EL;EL306;Δυτική Αττική;3;8;636
EL;EL307;Πειραιάς, Νήσοι;3;8;637
EL;EL4;Νησιά Αιγαίου, Κρήτη;1;8;638
EL;EL41;Βόρειο Αιγαίο;2;8;639
EL;EL411;Λέσβος, Λήμνος;3;8;640
EL;EL412;Ικαρία, Σάμος;3;8;641
EL;EL413;Χίος;3;8;642
EL;EL42;Νότιο Αιγαίο;2;8;643
EL;EL421;Κάλυμνος, Κάρπαθος – Ηρωική Νήσος Κάσος, Κως, Ρόδος;3;8;644
EL;EL422;Άνδρος, Θήρα, Κέα, Μήλος, Μύκονος, Νάξος, Πάρος, Σύρος, Τήνος;3;8;645
EL;EL43;Κρήτη;2;8;646
EL;EL431;Ηράκλειο;3;8;647
EL;EL432;Λασίθι;3;8;648
EL;EL433;Ρέθυμνο;3;8;649
EL;EL434;Χανιά;3;8;650
EL;EL5;Βόρεια Ελλάδα;1;8;651
EL;EL51;Aνατολική Μακεδονία, Θράκη;2;8;652
EL;EL511;Έβρος;3;8;653
EL;EL512;Ξάνθη;3;8;654
EL;EL513;Ροδόπη;3;8;655
EL;EL514;Δράμα;3;8;656
EL;EL515;Θάσος, Καβάλα;3;8;657
EL;EL52;Κεντρική Μακεδονία;2;8;658
EL;EL521;Ημαθία;3;8;659
EL;EL522;Θεσσαλονίκη;3;8;660
EL;EL523;Κιλκίς;3;8;661
EL;EL524;Πέλλα;3;8;662
EL;EL525;Πιερία;3;8;663
EL;EL526;Σέρρες;3;8;664
EL;EL527;Χαλκιδική;3;8;665
EL;EL53;Δυτική Μακεδονία;2;8;666
EL;EL531;Γρεβενά, Κοζάνη;3;8;667
EL;EL532;Καστοριά;3;8;668
EL;EL533;Φλώρινα;3;8;669
EL;EL54;Ήπειρος;2;8;670
EL;EL541;Άρτα, Πρέβεζα;3;8;671
EL;EL542;Θεσπρωτία;3;8;672
EL;EL543;Ιωάννινα;3;8;673
EL;EL6;Κεντρική Ελλάδα;1;8;674
EL;EL61;Θεσσαλία;2;8;675
EL;EL611;Καρδίτσα, Τρίκαλα;3;8;676
EL;EL612;Λάρισα;3;8;677
EL;EL613;Μαγνησία, Σποράδες;3;8;678
EL;EL62;Ιόνια Νησιά;2;8;679
EL;EL621;Ζάκυνθος;3;8;680
EL;EL622;Κέρκυρα;3;8;681
EL;EL623;Ιθάκη, Κεφαλληνία;3;8;682
EL;EL624;Λευκάδα;3;8;683
EL;EL63;Δυτική Ελλάδα;2;8;684
EL;EL631;Αιτωλοακαρνανία;3;8;685
EL;EL632;Αχαΐα;3;8;686
EL;EL633;Ηλεία;3;8;687
EL;EL64;Στερεά Ελλάδα;2;8;688
EL;EL641;Βοιωτία;3;8;689
EL;EL642;Εύβοια;3;8;690
EL;EL643;Ευρυτανία;3;8;691
EL;EL644;Φθιώτιδα;3;8;692
EL;EL645;Φωκίδα;3;8;693
EL;EL65;Πελοπόννησος;2;8;694
EL;EL651;Αργολίδα, Αρκαδία;3;8;695
EL;EL652;Κορινθία;3;8;696
EL;EL653;Λακωνία, Μεσσηνία;3;8;697
EL;ELZ;Extra-Regio NUTS 1;1;8;698
EL;ELZZ;Extra-Regio NUTS 2;2;8;699
EL;ELZZZ;Extra-Regio NUTS 3;3;8;700
ES;ES1;Noroeste;1;9;701
ES;ES11;Galicia;2;9;702
ES;ES111;A Coruña;3;9;703
ES;ES112;Lugo;3;9;704
ES;ES113;Ourense;3;9;705
ES;ES114;Pontevedra;3;9;706
ES;ES12;Principado de Asturias;2;9;707
ES;ES120;Asturias;3;9;708
ES;ES13;Cantabria;2;9;709
ES;ES130;Cantabria;3;9;710
ES;ES2;Noreste;1;9;711
ES;ES21;País Vasco;2;9;712
ES;ES211;Araba/Álava;3;9;713
ES;ES212;Gipuzkoa;3;9;714
ES;ES213;Bizkaia;3;9;715
ES;ES22;Comunidad Foral de Navarra;2;9;716
ES;ES220;Navarra;3;9;717
ES;ES23;La Rioja;2;9;718
ES;ES230;La Rioja;3;9;719
ES;ES24;Aragón;2;9;720
ES;ES241;Huesca;3;9;721
ES;ES242;Teruel;3;9;722
ES;ES243;Zaragoza;3;9;723
ES;ES3;Comunidad de Madrid;1;9;724
ES;ES30;Comunidad de Madrid;2;9;725
ES;ES300;Madrid;3;9;726
ES;ES4;Centro (ES);1;9;727
ES;ES41;Castilla y León;2;9;728
ES;ES411;Ávila;3;9;729
ES;ES412;Burgos;3;9;730
ES;ES413;León;3;9;731
ES;ES414;Palencia;3;9;732
ES;ES415;Salamanca;3;9;733
ES;ES416;Segovia;3;9;734
ES;ES417;Soria;3;9;735
ES;ES418;Valladolid;3;9;736
ES;ES419;Zamora;3;9;737
ES;ES42;Castilla-La Mancha;2;9;738
ES;ES421;Albacete;3;9;739
ES;ES422;Ciudad Real;3;9;740
ES;ES423;Cuenca;3;9;741
ES;ES424;Guadalajara;3;9;742
ES;ES425;Toledo;3;9;743
ES;ES43;Extremadura;2;9;744
ES;ES431;Badajoz;3;9;745
ES;ES432;Cáceres;3;9;746
ES;ES5;Este;1;9;747
ES;ES51;Cataluña;2;9;748
ES;ES511;Barcelona;3;9;749
ES;ES512;Girona;3;9;750
ES;ES513;Lleida;3;9;751
ES;ES514;Tarragona;3;9;752
ES;ES52;Comunitat Valenciana ;2;9;753
ES;ES521;Alicante/Alacant;3;9;754
ES;ES522;Castellón/Castelló;3;9;755
ES;ES523;Valencia/València;3;9;756
ES;ES53;Illes Balears;2;9;757
ES;ES531;Eivissa y Formentera;3;9;758
ES;ES532;Mallorca;3;9;759
ES;ES533;Menorca;3;9;760
ES;ES6;Sur;1;9;761
ES;ES61;Andalucía;2;9;762
ES;ES611;Almería;3;9;763
ES;ES612;Cádiz;3;9;764
ES;ES613;Córdoba;3;9;765
ES;ES614;Granada;3;9;766
ES;ES615;Huelva;3;9;767
ES;ES616;Jaén;3;9;768
ES;ES617;Málaga;3;9;769
ES;ES618;Sevilla;3;9;770
ES;ES62;Región de Murcia;2;9;771
ES;ES620;Murcia;3;9;772
ES;ES63;Ciudad de Ceuta;2;9;773
ES;ES630;Ceuta;3;9;774
ES;ES64;Ciudad de Melilla;2;9;775
ES;ES640;Melilla;3;9;776
ES;ES7;Canarias;1;9;777
ES;ES70;Canarias;2;9;778
ES;ES703;El Hierro;3;9;779
ES;ES704;Fuerteventura;3;9;780
ES;ES705;Gran Canaria;3;9;781
ES;ES706;La Gomera;3;9;782
ES;ES707;La Palma;3;9;783
ES;ES708;Lanzarote;3;9;784
ES;ES709;Tenerife;3;9;785
ES;ESZ;Extra-Regio NUTS 1;1;9;786
ES;ESZZ;Extra-Regio NUTS 2;2;9;787
ES;ESZZZ;Extra-Regio NUTS 3;3;9;788
FR;FR1;Ile-de-France;1;10;789
FR;FR10;Ile-de-France;2;10;790
FR;FR101;Paris;3;10;791
FR;FR102;Seine-et-Marne ;3;10;792
FR;FR103;Yvelines ;3;10;793
FR;FR104;Essonne;3;10;794
FR;FR105;Hauts-de-Seine ;3;10;795
FR;FR106;Seine-Saint-Denis ;3;10;796
FR;FR107;Val-de-Marne;3;10;797
FR;FR108;Val-d’Oise;3;10;798
FR;FRB;Centre — Val de Loire;1;10;799
FR;FRB0;Centre — Val de Loire;2;10;800
FR;FRB01;Cher;3;10;801
FR;FRB02;Eure-et-Loir;3;10;802
FR;FRB03;Indre;3;10;803
FR;FRB04;Indre-et-Loire;3;10;804
FR;FRB05;Loir-et-Cher;3;10;805
FR;FRB06;Loiret;3;10;806
FR;FRC;Bourgogne-Franche-Comté;1;10;807
FR;FRC1;Bourgogne;2;10;808
FR;FRC11;Côte-d’Or;3;10;809
FR;FRC12;Nièvre;3;10;810
FR;FRC13;Saône-et-Loire;3;10;811
FR;FRC14;Yonne;3;10;812
FR;FRC2;Franche-Comté;2;10;813
FR;FRC21;Doubs;3;10;814
FR;FRC22;Jura;3;10;815
FR;FRC23;Haute-Saône;3;10;816
FR;FRC24;Territoire de Belfort;3;10;817
FR;FRD;Normandie;1;10;818
FR;FRD1;Basse-Normandie ;2;10;819
FR;FRD11;Calvados ;3;10;820
FR;FRD12;Manche ;3;10;821
FR;FRD13;Orne;3;10;822
FR;FRD2;Haute-Normandie ;2;10;823
FR;FRD21;Eure;3;10;824
FR;FRD22;Seine-Maritime;3;10;825
FR;FRE;Hauts-de-France;1;10;826
FR;FRE1;Nord-Pas de Calais;2;10;827
FR;FRE11;Nord;3;10;828
FR;FRE12;Pas-de-Calais;3;10;829
FR;FRE2;Picardie;2;10;830
FR;FRE21;Aisne;3;10;831
FR;FRE22;Oise;3;10;832
FR;FRE23;Somme;3;10;833
FR;FRF;Grand Est;1;10;834
FR;FRF1;Alsace;2;10;835
FR;FRF11;Bas-Rhin;3;10;836
FR;FRF12;Haut-Rhin;3;10;837
FR;FRF2;Champagne-Ardenne;2;10;838
FR;FRF21;Ardennes;3;10;839
FR;FRF22;Aube;3;10;840
FR;FRF23;Marne;3;10;841
FR;FRF24;Haute-Marne;3;10;842
FR;FRF3;Lorraine;2;10;843
FR;FRF31;Meurthe-et-Moselle ;3;10;844
FR;FRF32;Meuse ;3;10;845
FR;FRF33;Moselle;3;10;846
FR;FRF34;Vosges;3;10;847
FR;FRG;Pays de la Loire;1;10;848
FR;FRG0;Pays de la Loire;2;10;849
FR;FRG01;Loire-Atlantique;3;10;850
FR;FRG02;Maine-et-Loire;3;10;851
FR;FRG03;Mayenne;3;10;852
FR;FRG04;Sarthe;3;10;853
FR;FRG05;Vendée;3;10;854
FR;FRH;Bretagne;1;10;855
FR;FRH0;Bretagne;2;10;856
FR;FRH01;Côtes-d’Armor;3;10;857
FR;FRH02;Finistère;3;10;858
FR;FRH03;Ille-et-Vilaine;3;10;859
FR;FRH04;Morbihan;3;10;860
FR;FRI;Nouvelle-Aquitaine;1;10;861
FR;FRI1;Aquitaine;2;10;862
FR;FRI11;Dordogne;3;10;863
FR;FRI12;Gironde;3;10;864
FR;FRI13;Landes;3;10;865
FR;FRI14;Lot-et-Garonne;3;10;866
FR;FRI15;Pyrénées-Atlantiques;3;10;867
FR;FRI2;Limousin;2;10;868
FR;FRI21;Corrèze;3;10;869
FR;FRI22;Creuse;3;10;870
FR;FRI23;Haute-Vienne;3;10;871
FR;FRI3;Poitou-Charentes;2;10;872
FR;FRI31;Charente;3;10;873
FR;FRI32;Charente-Maritime;3;10;874
FR;FRI33;Deux-Sèvres;3;10;875
FR;FRI34;Vienne;3;10;876
FR;FRJ;Occitanie;1;10;877
FR;FRJ1;Languedoc-Roussillon;2;10;878
FR;FRJ11;Aude;3;10;879
FR;FRJ12;Gard;3;10;880
FR;FRJ13;Hérault;3;10;881
FR;FRJ14;Lozère;3;10;882
FR;FRJ15;Pyrénées-Orientales;3;10;883
FR;FRJ2;Midi-Pyrénées;2;10;884
FR;FRJ21;Ariège;3;10;885
FR;FRJ22;Aveyron;3;10;886
FR;FRJ23;Haute-Garonne;3;10;887
FR;FRJ24;Gers;3;10;888
FR;FRJ25;Lot;3;10;889
FR;FRJ26;Hautes-Pyrénées ;3;10;890
FR;FRJ27;Tarn;3;10;891
FR;FRJ28;Tarn-et-Garonne;3;10;892
FR;FRK;Auvergne-Rhône-Alpes;1;10;893
FR;FRK1;Auvergne;2;10;894
FR;FRK11;Allier;3;10;895
FR;FRK12;Cantal;3;10;896
FR;FRK13;Haute-Loire;3;10;897
FR;FRK14;Puy-de-Dôme;3;10;898
FR;FRK2;Rhône-Alpes;2;10;899
FR;FRK21;Ain;3;10;900
FR;FRK22;Ardèche;3;10;901
FR;FRK23;Drôme;3;10;902
FR;FRK24;Isère;3;10;903
FR;FRK25;Loire;3;10;904
FR;FRK26;Rhône;3;10;905
FR;FRK27;Savoie;3;10;906
FR;FRK28;Haute-Savoie;3;10;907
FR;FRL;Provence-Alpes-Côte d’Azur;1;10;908
FR;FRL0;Provence-Alpes-Côte d’Azur;2;10;909
FR;FRL01;Alpes-de-Haute-Provence;3;10;910
FR;FRL02;Hautes-Alpes ;3;10;911
FR;FRL03;Alpes-Maritimes;3;10;912
FR;FRL04;Bouches-du-Rhône;3;10;913
FR;FRL05;Var;3;10;914
FR;FRL06;Vaucluse;3;10;915
FR;FRM;Corse;1;10;916
FR;FRM0;Corse;2;10;917
FR;FRM01;Corse-du-Sud;3;10;918
FR;FRM02;Haute-Corse;3;10;919
FR;FRY;RUP FR — Régions Ultrapériphériques Françaises;1;10;920
FR;FRY1;Guadeloupe;2;10;921
FR;FRY10;Guadeloupe;3;10;922
FR;FRY2;Martinique ;2;10;923
FR;FRY20;Martinique ;3;10;924
FR;FRY3;Guyane;2;10;925
FR;FRY30;Guyane;3;10;926
FR;FRY4;La Réunion ;2;10;927
FR;FRY40;La Réunion;3;10;928
FR;FRY5;Mayotte;2;10;929
FR;FRY50;Mayotte ;3;10;930
FR;FRZ;Extra-Regio NUTS 1;1;10;931
FR;FRZZ;Extra-Regio NUTS 2;2;10;932
FR;FRZZZ;Extra-Regio NUTS 3;3;10;933
HR;HR0;Hrvatska;1;11;934
HR;HR02;Panonska Hrvatska;2;11;935
HR;HR021;Bjelovarsko-bilogorska županija;3;11;936
HR;HR022;Virovitičko-podravska županija;3;11;937
HR;HR023;Požeško-slavonska županija;3;11;938
HR;HR024;Brodsko-posavska županija;3;11;939
HR;HR025;Osječko-baranjska županija;3;11;940
HR;HR026;Vukovarsko-srijemska županija;3;11;941
HR;HR027;Karlovačka županija;3;11;942
HR;HR028;Sisačko-moslavačka županija;3;11;943
HR;HR03;Jadranska Hrvatska;2;11;944
HR;HR031;Primorsko-goranska županija;3;11;945
HR;HR032;Ličko-senjska županija;3;11;946
HR;HR033;Zadarska županija;3;11;947
HR;HR034;Šibensko-kninska županija;3;11;948
HR;HR035;Splitsko-dalmatinska županija;3;11;949
HR;HR036;Istarska županija;3;11;950
HR;HR037;Dubrovačko-neretvanska županija;3;11;951
HR;HR05;Grad Zagreb;2;11;952
HR;HR050;Grad Zagreb;3;11;953
HR;HR06;Sjeverna Hrvatska;2;11;954
HR;HR061;Međimurska županija;3;11;955
HR;HR062;Varaždinska županija;3;11;956
HR;HR063;Koprivničko-križevačka županija;3;11;957
HR;HR064;Krapinsko-zagorska županija;3;11;958
HR;HR065;Zagrebačka županija;3;11;959
HR;HRZ;Extra-Regio NUTS 1;1;11;960
HR;HRZZ;Extra-Regio NUTS 2;2;11;961
HR;HRZZZ;Extra-Regio NUTS 3;3;11;962
IT;ITC;Nord-Ovest;1;12;963
IT;ITC1;Piemonte;2;12;964
IT;ITC11;Torino;3;12;965
IT;ITC12;Vercelli;3;12;966
IT;ITC13;Biella;3;12;967
IT;ITC14;Verbano-Cusio-Ossola;3;12;968
IT;ITC15;Novara;3;12;969
IT;ITC16;Cuneo;3;12;970
IT;ITC17;Asti;3;12;971
IT;ITC18;Alessandria;3;12;972
IT;ITC2;Valle d’Aosta/Vallée d’Aoste;2;12;973
IT;ITC20;Valle d’Aosta/Vallée d’Aoste;3;12;974
IT;ITC3;Liguria;2;12;975
IT;ITC31;Imperia;3;12;976
IT;ITC32;Savona;3;12;977
IT;ITC33;Genova;3;12;978
IT;ITC34;La Spezia;3;12;979
IT;ITC4;Lombardia;2;12;980
IT;ITC41;Varese;3;12;981
IT;ITC42;Como;3;12;982
IT;ITC43;Lecco;3;12;983
IT;ITC44;Sondrio;3;12;984
IT;ITC46;Bergamo;3;12;985
IT;ITC47;Brescia;3;12;986
IT;ITC48;Pavia;3;12;987
IT;ITC49;Lodi;3;12;988
IT;ITC4A;Cremona;3;12;989
IT;ITC4B;Mantova;3;12;990
IT;ITC4C;Milano;3;12;991
IT;ITC4D;Monza e della Brianza;3;12;992
IT;ITF;Sud;1;12;993
IT;ITF1;Abruzzo;2;12;994
IT;ITF11;L’Aquila;3;12;995
IT;ITF12;Teramo;3;12;996
IT;ITF13;Pescara;3;12;997
IT;ITF14;Chieti;3;12;998
IT;ITF2;Molise;2;12;999
IT;ITF21;Isernia;3;12;1000
IT;ITF22;Campobasso;3;12;1001
IT;ITF3;Campania;2;12;1002
IT;ITF31;Caserta;3;12;1003
IT;ITF32;Benevento;3;12;1004
IT;ITF33;Napoli;3;12;1005
IT;ITF34;Avellino;3;12;1006
IT;ITF35;Salerno;3;12;1007
IT;ITF4;Puglia;2;12;1008
IT;ITF43;Taranto;3;12;1009
IT;ITF44;Brindisi;3;12;1010
IT;ITF45;Lecce;3;12;1011
IT;ITF46;Foggia;3;12;1012
IT;ITF47;Bari;3;12;1013
IT;ITF48;Barletta-Andria-Trani;3;12;1014
IT;ITF5;Basilicata;2;12;1015
IT;ITF51;Potenza;3;12;1016
IT;ITF52;Matera;3;12;1017
IT;ITF6;Calabria;2;12;1018
IT;ITF61;Cosenza;3;12;1019
IT;ITF62;Crotone;3;12;1020
IT;ITF63;Catanzaro;3;12;1021
IT;ITF64;Vibo Valentia;3;12;1022
IT;ITF65;Reggio Calabria;3;12;1023
IT;ITG;Isole;1;12;1024
IT;ITG1;Sicilia;2;12;1025
IT;ITG11;Trapani;3;12;1026
IT;ITG12;Palermo;3;12;1027
IT;ITG13;Messina;3;12;1028
IT;ITG14;Agrigento;3;12;1029
IT;ITG15;Caltanissetta;3;12;1030
IT;ITG16;Enna;3;12;1031
IT;ITG17;Catania;3;12;1032
IT;ITG18;Ragusa;3;12;1033
IT;ITG19;Siracusa;3;12;1034
IT;ITG2;Sardegna;2;12;1035
IT;ITG2D;Sassari;3;12;1036
IT;ITG2E;Nuoro;3;12;1037
IT;ITG2F;Cagliari;3;12;1038
IT;ITG2G;Oristano;3;12;1039
IT;ITG2H;Sud Sardegna;3;12;1040
IT;ITH;Nord-Est;1;12;1041
IT;ITH1;Provincia Autonoma di Bolzano/Bozen;2;12;1042
IT;ITH10;Bolzano-Bozen;3;12;1043
IT;ITH2;Provincia Autonoma di Trento;2;12;1044
IT;ITH20;Trento;3;12;1045
IT;ITH3;Veneto;2;12;1046
IT;ITH31;Verona;3;12;1047
IT;ITH32;Vicenza;3;12;1048
IT;ITH33;Belluno;3;12;1049
IT;ITH34;Treviso;3;12;1050
IT;ITH35;Venezia;3;12;1051
IT;ITH36;Padova;3;12;1052
IT;ITH37;Rovigo;3;12;1053
IT;ITH4;Friuli-Venezia Giulia;2;12;1054
IT;ITH41;Pordenone;3;12;1055
IT;ITH42;Udine;3;12;1056
IT;ITH43;Gorizia;3;12;1057
IT;ITH44;Trieste;3;12;1058
IT;ITH5;Emilia-Romagna;2;12;1059
IT;ITH51;Piacenza;3;12;1060
IT;ITH52;Parma;3;12;1061
IT;ITH53;Reggio nell’Emilia;3;12;1062
IT;ITH54;Modena;3;12;1063
IT;ITH55;Bologna;3;12;1064
IT;ITH56;Ferrara;3;12;1065
IT;ITH57;Ravenna;3;12;1066
IT;ITH58;Forlì-Cesena;3;12;1067
IT;ITH59;Rimini;3;12;1068
IT;ITI;Centro (IT);1;12;1069
IT;ITI1;Toscana;2;12;1070
IT;ITI11;Massa-Carrara;3;12;1071
IT;ITI12;Lucca;3;12;1072
IT;ITI13;Pistoia;3;12;1073
IT;ITI14;Firenze;3;12;1074
IT;ITI15;Prato;3;12;1075
IT;ITI16;Livorno;3;12;1076
IT;ITI17;Pisa;3;12;1077
IT;ITI18;Arezzo;3;12;1078
IT;ITI19;Siena;3;12;1079
IT;ITI1A;Grosseto;3;12;1080
IT;ITI2;Umbria;2;12;1081
IT;ITI21;Perugia;3;12;1082
IT;ITI22;Terni;3;12;1083
IT;ITI3;Marche;2;12;1084
IT;ITI31;Pesaro e Urbino;3;12;1085
IT;ITI32;Ancona;3;12;1086
IT;ITI33;Macerata;3;12;1087
IT;ITI34;Ascoli Piceno;3;12;1088
IT;ITI35;Fermo;3;12;1089
IT;ITI4;Lazio;2;12;1090
IT;ITI41;Viterbo;3;12;1091
IT;ITI42;Rieti;3;12;1092
IT;ITI43;Roma;3;12;1093
IT;ITI44;Latina;3;12;1094
IT;ITI45;Frosinone;3;12;1095
IT;ITZ;Extra-Regio NUTS 1;1;12;1096
IT;ITZZ;Extra-Regio NUTS 2;2;12;1097
IT;ITZZZ;Extra-Regio NUTS 3;3;12;1098
CY;CY0;Κύπρος;1;13;1099
CY;CY00;Κύπρος;2;13;1100
CY;CY000;Κύπρος;3;13;1101
CY;CYZ;Extra-Regio NUTS 1;1;13;1102
CY;CYZZ;Extra-Regio NUTS 2;2;13;1103
CY;CYZZZ;Extra-Regio NUTS 3;3;13;1104
LV;LV0;Latvija;1;14;1105
LV;LV00;Latvija;2;14;1106
LV;LV005;Latgale;3;14;1107
LV;LV009;Zemgale;3;14;1108
LV;LV00A;Rīga;3;14;1109
LV;LV00B;Kurzeme;3;14;1110
LV;LV00C;Vidzeme;3;14;1111
LV;LVZ;Extra-Regio NUTS 1;1;14;1112
LV;LVZZ;Extra-Regio NUTS 2;2;14;1113
LV;LVZZZ;Extra-Regio NUTS 3;3;14;1114
LT;LT0;Lietuva;1;15;1115
LT;LT01;Sostinės regionas;2;15;1116
LT;LT011;Vilniaus apskritis;3;15;1117
LT;LT02;Vidurio ir vakarų Lietuvos regionas ;2;15;1118
LT;LT021;Alytaus apskritis;3;15;1119
LT;LT022;Kauno apskritis;3;15;1120
LT;LT023;Klaipėdos apskritis;3;15;1121
LT;LT024;Marijampolės apskritis;3;15;1122
LT;LT025;Panevėžio apskritis;3;15;1123
LT;LT026;Šiaulių apskritis;3;15;1124
LT;LT027;Tauragės apskritis;3;15;1125
LT;LT028;Telšių apskritis;3;15;1126
LT;LT029;Utenos apskritis;3;15;1127
LT;LTZ;Extra-Regio NUTS 1;1;15;1128
LT;LTZZ;Extra-Regio NUTS 2;2;15;1129
LT;LTZZZ;Extra-Regio NUTS 3;3;15;1130
LU;LU0;Luxembourg;1;16;1131
LU;LU00;Luxembourg;2;16;1132
LU;LU000;Luxembourg;3;16;1133
LU;LUZ;Extra-Regio NUTS 1;1;16;1134
LU;LUZZ;Extra-Regio NUTS 2;2;16;1135
LU;LUZZZ;Extra-Regio NUTS 3;3;16;1136
HU;HU1;Közép-Magyarország;1;17;1137
HU;HU11;Budapest;2;17;1138
HU;HU110;Budapest;3;17;1139
HU;HU12;Pest;2;17;1140
HU;HU120;Pest;3;17;1141
HU;HU2;Dunántúl;1;17;1142
HU;HU21;Közép-Dunántúl;2;17;1143
HU;HU211;Fejér;3;17;1144
HU;HU212;Komárom-Esztergom;3;17;1145
HU;HU213;Veszprém;3;17;1146
HU;HU22;Nyugat-Dunántúl;2;17;1147
HU;HU221;Győr-Moson-Sopron;3;17;1148
HU;HU222;Vas;3;17;1149
HU;HU223;Zala;3;17;1150
HU;HU23;Dél-Dunántúl;2;17;1151
HU;HU231;Baranya;3;17;1152
HU;HU232;Somogy;3;17;1153
HU;HU233;Tolna;3;17;1154
HU;HU3;Alföld és Észak;1;17;1155
HU;HU31;Észak-Magyarország;2;17;1156
HU;HU311;Borsod-Abaúj-Zemplén;3;17;1157
HU;HU312;Heves;3;17;1158
HU;HU313;Nógrád;3;17;1159
HU;HU32;Észak-Alföld;2;17;1160
HU;HU321;Hajdú-Bihar;3;17;1161
HU;HU322;Jász-Nagykun-Szolnok;3;17;1162
HU;HU323;Szabolcs-Szatmár-Bereg;3;17;1163
HU;HU33;Dél-Alföld;2;17;1164
HU;HU331;Bács-Kiskun;3;17;1165
HU;HU332;Békés;3;17;1166
HU;HU333;Csongrád-Csanád;3;17;1167
HU;HUZ;Extra-Regio NUTS 1;1;17;1168
HU;HUZZ;Extra-Regio NUTS 2;2;17;1169
HU;HUZZZ;Extra-Regio NUTS 3;3;17;1170
MT;MT0;Malta;1;18;1171
MT;MT00;Malta;2;18;1172
MT;MT001;Malta;3;18;1173
MT;MT002;Gozo and Comino/Għawdex u Kemmuna;3;18;1174
MT;MTZ;Extra-Regio NUTS 1;1;18;1175
MT;MTZZ;Extra-Regio NUTS 2;2;18;1176
MT;MTZZZ;Extra-Regio NUTS 3;3;18;1177
NL;NL1;Noord-Nederland;1;19;1178
NL;NL11;Groningen;2;19;1179
NL;NL112;Delfzijl en omgeving;3;19;1180
NL;NL114;Oost-Groningen;3;19;1181
NL;NL115;Overig Groningen;3;19;1182
NL;NL12;Friesland (NL);2;19;1183
NL;NL126;Zuidoost-Friesland;3;19;1184
NL;NL127;Noord-Friesland;3;19;1185
NL;NL128;Zuidwest-Friesland;3;19;1186
NL;NL13;Drenthe;2;19;1187
NL;NL131;Noord-Drenthe;3;19;1188
NL;NL132;Zuidoost-Drenthe;3;19;1189
NL;NL133;Zuidwest-Drenthe;3;19;1190
NL;NL2;Oost-Nederland;1;19;1191
NL;NL21;Overijssel;2;19;1192
NL;NL211;Noord-Overijssel;3;19;1193
NL;NL212;Zuidwest-Overijssel;3;19;1194
NL;NL213;Twente;3;19;1195
NL;NL22;Gelderland;2;19;1196
NL;NL221;Veluwe;3;19;1197
NL;NL224;Zuidwest-Gelderland;3;19;1198
NL;NL225;Achterhoek;3;19;1199
NL;NL226;Arnhem/Nijmegen;3;19;1200
NL;NL23;Flevoland;2;19;1201
NL;NL230;Flevoland;3;19;1202
NL;NL3;West-Nederland;1;19;1203
NL;NL32;Noord-Holland;2;19;1204
NL;NL321;Kop van Noord-Holland;3;19;1205
NL;NL323;IJmond;3;19;1206
NL;NL325;Zaanstreek;3;19;1207
NL;NL327;Het Gooi en Vechtstreek;3;19;1208
NL;NL328;Alkmaar en omgeving;3;19;1209
NL;NL32A;Agglomeratie Haarlem;3;19;1210
NL;NL32B;Groot-Amsterdam;3;19;1211
NL;NL34;Zeeland;2;19;1212
NL;NL341;Zeeuwsch-Vlaanderen;3;19;1213
NL;NL342;Overig Zeeland;3;19;1214
NL;NL35;Utrecht;2;19;1215
NL;NL350;Utrecht;3;19;1216
NL;NL36;Zuid-Holland;2;19;1217
NL;NL361;Agglomeratie ’s-Gravenhage;3;19;1218
NL;NL362;Delft en Westland;3;19;1219
NL;NL363;Agglomeratie Leiden en Bollenstreek;3;19;1220
NL;NL364;Zuidoost-Zuid-Holland;3;19;1221
NL;NL365;Oost-Zuid-Holland;3;19;1222
NL;NL366;Groot-Rijnmond;3;19;1223
NL;NL4;Zuid-Nederland;1;19;1224
NL;NL41;Noord-Brabant;2;19;1225
NL;NL411;West-Noord-Brabant;3;19;1226
NL;NL414;Zuidoost-Noord-Brabant;3;19;1227
NL;NL415;Midden-Noord-Brabant;3;19;1228
NL;NL416;Noordoost-Noord-Brabant;3;19;1229
NL;NL42;Limburg (NL);2;19;1230
NL;NL421;Noord-Limburg;3;19;1231
NL;NL422;Midden-Limburg;3;19;1232
NL;NL423;Zuid-Limburg;3;19;1233
NL;NLZ;Extra-Regio NUTS 1;1;19;1234
NL;NLZZ;Extra-Regio NUTS 2;2;19;1235
NL;NLZZZ;Extra-Regio NUTS 3;3;19;1236
AT;AT1;Ostösterreich;1;20;1237
AT;AT11;Burgenland;2;20;1238
AT;AT111;Mittelburgenland;3;20;1239
AT;AT112;Nordburgenland;3;20;1240
AT;AT113;Südburgenland;3;20;1241
AT;AT12;Niederösterreich;2;20;1242
AT;AT121;Mostviertel-Eisenwurzen;3;20;1243
AT;AT122;Niederösterreich-Süd;3;20;1244
AT;AT123;Sankt Pölten;3;20;1245
AT;AT124;Waldviertel;3;20;1246
AT;AT125;Weinviertel;3;20;1247
AT;AT126;Wiener Umland/Nordteil;3;20;1248
AT;AT127;Wiener Umland/Südteil;3;20;1249
AT;AT13;Wien;2;20;1250
AT;AT130;Wien;3;20;1251
AT;AT2;Südösterreich;1;20;1252
AT;AT21;Kärnten;2;20;1253
AT;AT211;Klagenfurt-Villach;3;20;1254
AT;AT212;Oberkärnten;3;20;1255
AT;AT213;Unterkärnten;3;20;1256
AT;AT22;Steiermark;2;20;1257
AT;AT221;Graz;3;20;1258
AT;AT222;Liezen;3;20;1259
AT;AT223;Östliche Obersteiermark;3;20;1260
AT;AT224;Oststeiermark;3;20;1261
AT;AT225;West- und Südsteiermark;3;20;1262
AT;AT226;Westliche Obersteiermark;3;20;1263
AT;AT3;Westösterreich;1;20;1264
AT;AT31;Oberösterreich;2;20;1265
AT;AT311;Innviertel;3;20;1266
AT;AT312;Linz-Wels;3;20;1267
AT;AT313;Mühlviertel;3;20;1268
AT;AT314;Steyr-Kirchdorf;3;20;1269
AT;AT315;Traunviertel;3;20;1270
AT;AT32;Salzburg;2;20;1271
AT;AT321;Lungau;3;20;1272
AT;AT322;Pinzgau-Pongau;3;20;1273
AT;AT323;Salzburg und Umgebung;3;20;1274
AT;AT33;Tirol;2;20;1275
AT;AT331;Außerfern;3;20;1276
AT;AT332;Innsbruck;3;20;1277
AT;AT333;Osttirol;3;20;1278
AT;AT334;Tiroler Oberland;3;20;1279
AT;AT335;Tiroler Unterland;3;20;1280
AT;AT34;Vorarlberg;2;20;1281
AT;AT341;Bludenz-Bregenzer Wald;3;20;1282
AT;AT342;Rheintal-Bodenseegebiet;3;20;1283
AT;ATZ;Extra-Regio NUTS 1;1;20;1284
AT;ATZZ;Extra-Regio NUTS 2;2;20;1285
AT;ATZZZ;Extra-Regio NUTS 3;3;20;1286
PL;PL2;Makroregion południowy;1;21;1287
PL;PL21;Małopolskie;2;21;1288
PL;PL213;Miasto Kraków;3;21;1289
PL;PL214;Krakowski;3;21;1290
PL;PL217;Tarnowski;3;21;1291
PL;PL218;Nowosądecki;3;21;1292
PL;PL219;Nowotarski;3;21;1293
PL;PL21A;Oświęcimski;3;21;1294
PL;PL22;Śląskie;2;21;1295
PL;PL224;Częstochowski;3;21;1296
PL;PL225;Bielski;3;21;1297
PL;PL227;Rybnicki;3;21;1298
PL;PL228;Bytomski;3;21;1299
PL;PL229;Gliwicki;3;21;1300
PL;PL22A;Katowicki;3;21;1301
PL;PL22B;Sosnowiecki;3;21;1302
PL;PL22C;Tyski;3;21;1303
PL;PL4;Makroregion północno-zachodni;1;21;1304
PL;PL41;Wielkopolskie;2;21;1305
PL;PL411;Pilski;3;21;1306
PL;PL414;Koniński;3;21;1307
PL;PL415;Miasto Poznań;3;21;1308
PL;PL416;Kaliski;3;21;1309
PL;PL417;Leszczyński;3;21;1310
PL;PL418;Poznański;3;21;1311
PL;PL42;Zachodniopomorskie;2;21;1312
PL;PL424;Miasto Szczecin;3;21;1313
PL;PL426;Koszaliński;3;21;1314
PL;PL427;Szczecinecko-pyrzycki;3;21;1315
PL;PL428;Szczeciński;3;21;1316
PL;PL43;Lubuskie;2;21;1317
PL;PL431;Gorzowski;3;21;1318
PL;PL432;Zielonogórski;3;21;1319
PL;PL5;Makroregion południowo-zachodni;1;21;1320
PL;PL51;Dolnośląskie;2;21;1321
PL;PL514;Miasto Wrocław;3;21;1322
PL;PL515;Jeleniogórski;3;21;1323
PL;PL516;Legnicko-głogowski;3;21;1324
PL;PL517;Wałbrzyski;3;21;1325
PL;PL518;Wrocławski;3;21;1326
PL;PL52;Opolskie;2;21;1327
PL;PL523;Nyski;3;21;1328
PL;PL524;Opolski;3;21;1329
PL;PL6;Makroregion północny;1;21;1330
PL;PL61;Kujawsko-pomorskie;2;21;1331
PL;PL613;Bydgosko-toruński;3;21;1332
PL;PL616;Grudziądzki;3;21;1333
PL;PL617;Inowrocławski;3;21;1334
PL;PL618;Świecki;3;21;1335
PL;PL619;Włocławski;3;21;1336
PL;PL62;Warmińsko-mazurskie;2;21;1337
PL;PL621;Elbląski;3;21;1338
PL;PL622;Olsztyński;3;21;1339
PL;PL623;Ełcki;3;21;1340
PL;PL63;Pomorskie;2;21;1341
PL;PL633;Trójmiejski;3;21;1342
PL;PL634;Gdański;3;21;1343
PL;PL636;Słupski;3;21;1344
PL;PL637;Chojnicki;3;21;1345
PL;PL638;Starogardzki;3;21;1346
PL;PL7;Makroregion centralny;1;21;1347
PL;PL71;Łódzkie;2;21;1348
PL;PL711;Miasto Łódź;3;21;1349
PL;PL712;Łódzki;3;21;1350
PL;PL713;Piotrkowski;3;21;1351
PL;PL714;Sieradzki;3;21;1352
PL;PL715;Skierniewicki;3;21;1353
PL;PL72;Świętokrzyskie;2;21;1354
PL;PL721;Kielecki;3;21;1355
PL;PL722;Sandomiersko-jędrzejowski;3;21;1356
PL;PL8;Makroregion wschodni;1;21;1357
PL;PL81;Lubelskie;2;21;1358
PL;PL811;Bialski;3;21;1359
PL;PL812;Chełmsko-zamojski;3;21;1360
PL;PL814;Lubelski;3;21;1361
PL;PL815;Puławski;3;21;1362
PL;PL82;Podkarpackie;2;21;1363
PL;PL821;Krośnieński;3;21;1364
PL;PL822;Przemyski;3;21;1365
PL;PL823;Rzeszowski;3;21;1366
PL;PL824;Tarnobrzeski;3;21;1367
PL;PL84;Podlaskie;2;21;1368
PL;PL841;Białostocki;3;21;1369
PL;PL842;Łomżyński;3;21;1370
PL;PL843;Suwalski;3;21;1371
PL;PL9;Makroregion województwo mazowieckie;1;21;1372
PL;PL91;Warszawski stołeczny;2;21;1373
PL;PL911;Miasto Warszawa;3;21;1374
PL;PL912;Warszawski wschodni;3;21;1375
PL;PL913;Warszawski zachodni;3;21;1376
PL;PL92;Mazowiecki regionalny;2;21;1377
PL;PL921;Radomski;3;21;1378
PL;PL922;Ciechanowski;3;21;1379
PL;PL923;Płocki;3;21;1380
PL;PL924;Ostrołęcki;3;21;1381
PL;PL925;Siedlecki;3;21;1382
PL;PL926;Żyrardowski;3;21;1383
PL;PLZ;Extra-Regio NUTS 1;1;21;1384
PL;PLZZ;Extra-Regio NUTS 2;2;21;1385
PL;PLZZZ;Extra-Regio NUTS 3;3;21;1386
PT;PT1;Continente;1;22;1387
PT;PT11;Norte;2;22;1388
PT;PT111;Alto Minho;3;22;1389
PT;PT112;Cávado;3;22;1390
PT;PT119;Ave;3;22;1391
PT;PT11A;Área Metropolitana do Porto;3;22;1392
PT;PT11B;Alto Tâmega e Barroso;3;22;1393
PT;PT11C;Tâmega e Sousa;3;22;1394
PT;PT11D;Douro;3;22;1395
PT;PT11E;Terras de Trás-os-Montes;3;22;1396
PT;PT15;Algarve;2;22;1397
PT;PT150;Algarve;3;22;1398
PT;PT19;Centro (PT);2;22;1399
PT;PT191;Região de Aveiro;3;22;1400
PT;PT192;Região de Coimbra;3;22;1401
PT;PT193;Região de Leiria;3;22;1402
PT;PT194;Viseu Dão Lafões;3;22;1403
PT;PT195;Beira Baixa;3;22;1404
PT;PT196;Beiras e Serra da Estrela;3;22;1405
PT;PT1A;Grande Lisboa;2;22;1406
PT;PT1A0;Grande Lisboa;3;22;1407
PT;PT1B;Península de Setúbal;2;22;1408
PT;PT1B0;Península de Setúbal;3;22;1409
PT;PT1C;Alentejo;2;22;1410
PT;PT1C1;Alentejo Litoral;3;22;1411
PT;PT1C2;Baixo Alentejo;3;22;1412
PT;PT1C3;Alto Alentejo;3;22;1413
PT;PT1C4;Alentejo Central;3;22;1414
PT;PT1D;Oeste e Vale do Tejo;2;22;1415
PT;PT1D1;Oeste;3;22;1416
PT;PT1D2;Médio Tejo;3;22;1417
PT;PT1D3;Lezíria do Tejo;3;22;1418
PT;PT2;Região Autónoma dos Açores;1;22;1419
PT;PT20;Região Autónoma dos Açores;2;22;1420
PT;PT200;Região Autónoma dos Açores;3;22;1421
PT;PT3;Região Autónoma da Madeira;1;22;1422
PT;PT30;Região Autónoma da Madeira;2;22;1423
PT;PT300;Região Autónoma da Madeira;3;22;1424
PT;PTZ;Extra-Regio NUTS 1;1;22;1425
PT;PTZZ;Extra-Regio NUTS 2;2;22;1426
PT;PTZZZ;Extra-Regio NUTS 3;3;22;1427
RO;RO1;Macroregiunea Unu;1;23;1428
RO;RO11;Nord-Vest;2;23;1429
RO;RO111;Bihor;3;23;1430
RO;RO112;Bistriţa-Năsăud;3;23;1431
RO;RO113;Cluj;3;23;1432
RO;RO114;Maramureş;3;23;1433
RO;RO115;Satu Mare;3;23;1434
RO;RO116;Sălaj;3;23;1435
RO;RO12;Centru;2;23;1436
RO;RO121;Alba;3;23;1437
RO;RO122;Braşov;3;23;1438
RO;RO123;Covasna;3;23;1439
RO;RO124;Harghita;3;23;1440
RO;RO125;Mureş;3;23;1441
RO;RO126;Sibiu;3;23;1442
RO;RO2;Macroregiunea Doi;1;23;1443
RO;RO21;Nord-Est;2;23;1444
RO;RO211;Bacău;3;23;1445
RO;RO212;Botoşani;3;23;1446
RO;RO213;Iaşi;3;23;1447
RO;RO214;Neamţ;3;23;1448
RO;RO215;Suceava;3;23;1449
RO;RO216;Vaslui;3;23;1450
RO;RO22;Sud-Est;2;23;1451
RO;RO221;Brăila;3;23;1452
RO;RO222;Buzău;3;23;1453
RO;RO223;Constanţa;3;23;1454
RO;RO224;Galaţi;3;23;1455
RO;RO225;Tulcea;3;23;1456
RO;RO226;Vrancea;3;23;1457
RO;RO3;Macroregiunea Trei;1;23;1458
RO;RO31;Sud-Muntenia;2;23;1459
RO;RO311;Argeş;3;23;1460
RO;RO312;Călăraşi;3;23;1461
RO;RO313;Dâmboviţa;3;23;1462
RO;RO314;Giurgiu;3;23;1463
RO;RO315;Ialomiţa;3;23;1464
RO;RO316;Prahova;3;23;1465
RO;RO317;Teleorman;3;23;1466
RO;RO32;Bucureşti-Ilfov;2;23;1467
RO;RO321;Bucureşti;3;23;1468
RO;RO322;Ilfov;3;23;1469
RO;RO4;Macroregiunea Patru;1;23;1470
RO;RO41;Sud-Vest Oltenia;2;23;1471
RO;RO411;Dolj;3;23;1472
RO;RO412;Gorj;3;23;1473
RO;RO413;Mehedinţi;3;23;1474
RO;RO414;Olt;3;23;1475
RO;RO415;Vâlcea;3;23;1476
RO;RO42;Vest;2;23;1477
RO;RO421;Arad;3;23;1478
RO;RO422;Caraş-Severin;3;23;1479
RO;RO423;Hunedoara;3;23;1480
RO;RO424;Timiş;3;23;1481
RO;ROZ;Extra-Regio NUTS 1;1;23;1482
RO;ROZZ;Extra-Regio NUTS 2;2;23;1483
RO;ROZZZ;Extra-Regio NUTS 3;3;23;1484
SI;SI0;Slovenija;1;24;1485
SI;SI03;Vzhodna Slovenija;2;24;1486
SI;SI031;Pomurska;3;24;1487
SI;SI032;Podravska;3;24;1488
SI;SI033;Koroška;3;24;1489
SI;SI034;Savinjska;3;24;1490
SI;SI035;Zasavska;3;24;1491
SI;SI036;Posavska;3;24;1492
SI;SI037;Jugovzhodna Slovenija;3;24;1493
SI;SI038;Primorsko-notranjska;3;24;1494
SI;SI04;Zahodna Slovenija;2;24;1495
SI;SI041;Osrednjeslovenska;3;24;1496
SI;SI042;Gorenjska;3;24;1497
SI;SI043;Goriška;3;24;1498
SI;SI044;Obalno-kraška;3;24;1499
SI;SIZ;Extra-Regio NUTS 1;1;24;1500
SI;SIZZ;Extra-Regio NUTS 2;2;24;1501
SI;SIZZZ;Extra-Regio NUTS 3;3;24;1502
SK;SK0;Slovensko;1;25;1503
SK;SK01;Bratislavský kraj;2;25;1504
SK;SK010;Bratislavský kraj;3;25;1505
SK;SK02;Západné Slovensko;2;25;1506
SK;SK021;Trnavský kraj;3;25;1507
SK;SK022;Trenčiansky kraj;3;25;1508
SK;SK023;Nitriansky kraj;3;25;1509
SK;SK03;Stredné Slovensko;2;25;1510
SK;SK031;Žilinský kraj;3;25;1511
SK;SK032;Banskobystrický kraj;3;25;1512
SK;SK04;Východné Slovensko;2;25;1513
SK;SK041;Prešovský kraj;3;25;1514
SK;SK042;Košický kraj;3;25;1515
SK;SKZ;Extra-Regio NUTS 1;1;25;1516
SK;SKZZ;Extra-Regio NUTS 2;2;25;1517
SK;SKZZZ;Extra-Regio NUTS 3;3;25;1518
FI;FI1;Manner-Suomi;1;26;1519
FI;FI19;Länsi-Suomi;2;26;1520
FI;FI196;Satakunta;3;26;1521
FI;FI198;Keski-Suomi;3;26;1522
FI;FI199;Etelä-Pohjanmaa;3;26;1523
FI;FI19A;Pohjanmaa;3;26;1524
FI;FI19B;Pirkanmaa;3;26;1525
FI;FI1B;Helsinki-Uusimaa;2;26;1526
FI;FI1B1;Helsinki-Uusimaa;3;26;1527
FI;FI1C;Etelä-Suomi;2;26;1528
FI;FI1C1;Varsinais-Suomi;3;26;1529
FI;FI1C2;Kanta-Häme;3;26;1530
FI;FI1C5;Etelä-Karjala;3;26;1531
FI;FI1C6;Päijät-Häme;3;26;1532
FI;FI1C7;Kymenlaakso;3;26;1533
FI;FI1D;Pohjois- ja Itä-Suomi;2;26;1534
FI;FI1D5;Keski-Pohjanmaa;3;26;1535
FI;FI1D7;Lappi;3;26;1536
FI;FI1D8;Kainuu;3;26;1537
FI;FI1D9;Pohjois-Pohjanmaa;3;26;1538
FI;FI1DA;Etelä-Savo;3;26;1539
FI;FI1DB;Pohjois-Savo;3;26;1540
FI;FI1DC;Pohjois-Karjala;3;26;1541
FI;FI2;Åland;1;26;1542
FI;FI20;Åland;2;26;1543
FI;FI200;Åland;3;26;1544
FI;FIZ;Extra-Regio NUTS 1;1;26;1545
FI;FIZZ;Extra-Regio NUTS 2;2;26;1546
FI;FIZZZ;Extra-Regio NUTS 3;3;26;1547
SE;SE1;Östra Sverige;1;27;1548
SE;SE11;Stockholm;2;27;1549
SE;SE110;Stockholms län;3;27;1550
SE;SE12;Östra Mellansverige;2;27;1551
SE;SE121;Uppsala län;3;27;1552
SE;SE122;Södermanlands län;3;27;1553
SE;SE123;Östergötlands län;3;27;1554
SE;SE124;Örebro län;3;27;1555
SE;SE125;Västmanlands län;3;27;1556
SE;SE2;Södra Sverige;1;27;1557
SE;SE21;Småland med öarna;2;27;1558
SE;SE211;Jönköpings län;3;27;1559
SE;SE212;Kronobergs län;3;27;1560
SE;SE213;Kalmar län;3;27;1561
SE;SE214;Gotlands län;3;27;1562
SE;SE22;Sydsverige;2;27;1563
SE;SE221;Blekinge län;3;27;1564
SE;SE224;Skåne län;3;27;1565
SE;SE23;Västsverige;2;27;1566
SE;SE231;Hallands län;3;27;1567
SE;SE232;Västra Götalands län;3;27;1568
SE;SE3;Norra Sverige;1;27;1569
SE;SE31;Norra Mellansverige;2;27;1570
SE;SE311;Värmlands län;3;27;1571
SE;SE312;Dalarnas län;3;27;1572
SE;SE313;Gävleborgs län;3;27;1573
SE;SE32;Mellersta Norrland;2;27;1574
SE;SE321;Västernorrlands län;3;27;1575
SE;SE322;Jämtlands län;3;27;1576
SE;SE33;Övre Norrland;2;27;1577
SE;SE331;Västerbottens län;3;27;1578
SE;SE332;Norrbottens län;3;27;1579
SE;SEZ;Extra-Regio NUTS 1;1;27;1580
SE;SEZZ;Extra-Regio NUTS 2;2;27;1581
SE;SEZZZ;Extra-Regio NUTS 3;3;27;1582
`

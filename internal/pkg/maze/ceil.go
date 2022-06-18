package maze

type Ceil rune

var (
	Blank Ceil = ' ' //U+0020

	Itype_Hrz_Sng Ceil = '─' //U+2500
	Itype_Hrz_Dbl Ceil = '═' //U+2550
	Itype_Vrt_Sng Ceil = '│' //U+2502
	Itype_Vrt_Dbl Ceil = '║' //U+2551
	Ltype_DwnRgh_Sng Ceil = '┌' //U+250C
	Ltype_DwnRgh_Dbl Ceil = '╔' //U+2554
	Ltype_DwnLft_Sng Ceil = '┐' //U+2510
	Ltype_DwnLft_Dbl Ceil = '╗' //U+2557
	Ltype_HghRgh_Sng Ceil = '└' //U+2514
	Ltype_HghRgh_Dbl Ceil = '╚' //U+255A
	Ltype_HghLft_Sng Ceil = '┘' //U+2518
	Ltype_HghLft_Dbl Ceil = '╝' //U+255D
	Ttype_Hrz_Dbl_Dwn_Sng Ceil = '╤' //U+2564
	Ttype_Vrt_Dbl_Rgh_Sng Ceil = '╟' //U+255F
	Ttype_Vrt_Dbl_Lft_Sng Ceil = '╢' //U+2562

	Pill Ceil = '●' //U+25CF
	PowerPill Ceil = '○' //U+25CB
	//Pill and PowerPill Alternatives: • ● ○ ☼
)
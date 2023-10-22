package card

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestDescrambleBytes(t *testing.T) {
	// Some raw bytes from different cards
	examples := []struct {
		bytes, want string
	}{
		{
			"200435043f04430431043b04380447043a043804200044043e043d04340420003704300420003704340440043004320441044204320435043d043e0420003e044104380433044304400430045a043504",
			"Републички фонд за здравствено осигурање",
		},
		{
			"210440043104380458043004",
			"Србија",
		},
		{
			"1d0418041a041e041b041004",
			"НИКОЛА",
		},
		{
			"140420041004130410041d04",
			"ДРАГАН",
		},
		{
			"1c0418041b041e04090423041104",
			"МИЛОЉУБ",
		},
		{
			"080410041d041a041e04",
			"ЈАНКО",
		},
		{
			"21042204150412041e04",
			"СТЕВО",
		},
		{
			"21041204150422041b0410041d041004",
			"СВЕТЛАНА",
		},
		{
			"110415041e041304200410041404",
			"БЕОГРАД",
		},
		{
			"1704200415040a0410041d0418041d04",
			"ЗРЕЊАНИН",
		},
		{
			"170430043f043e0441043b0435043d0438042000430420003f044004380432044004350434043d043e043c04200034044004430448044204320443042c00200034044004430433043e043c0420003f044004300432043d043e043c0420003b043804460443042c0020003a043e04340420003f0440043504340443043704350442043d0438043a0430042c00200046043804320438043b043d04300420003b0438044604300420003d043004200041043b04430436043104380420004304200032043e045804410446043804",
			"Запослени у привредном друштву, другом правном лицу, код предузетника, цивилна лица на служби у војсци",
		},
		{
			"210430043c043e044104420430043b043d0430042000370430043d043004420441043a0435042000340435043b04300442043d043e0441044204",
			"Самостална занатске делатност",
		},
		{
			"110443045f04350442042000200435043f04430431043b0438043a0435042000210440043104380458043504",
			"Буџет Републике Србије",
		},
	}

	for _, example := range examples {
		t.Run(
			example.want,
			func(t *testing.T) {
				decoded, _ := hex.DecodeString(example.bytes)
				d := descrambleBytes(decoded)
				if strings.Compare(string(d), example.want) != 0 {
					t.Error("Got", string(d), "expected", example.want)
				}
			},
		)
	}
}

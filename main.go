package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"strings"
)

type ChemicalUnit int

const (
	Grams ChemicalUnit = iota
	Milliliters
)

type Consumption int

const (
	SafeForConsumption = iota
	NotSafeForConsumption
	MoreInformationRequired
)

type Formula struct {
	Nutrients []Nutrient
}

type Nutrient struct {
	Name            string
	ChemicalFormula string
	Value           float64
	Description     string
	Safe            Consumption
	SafeNotes       string
	Units           ChemicalUnit
}

type Stringer interface {
	String() string
}

func (cu ChemicalUnit) String() string {
	switch cu {
	case Grams:
		return "Grams"
	case Milliliters:
		return "Milliliters"
	default:
		return fmt.Sprintf("%v", int(cu))
	}
}

func (consumption Consumption) String() string {
	switch consumption {
	case SafeForConsumption:
		return "Yes"
	case NotSafeForConsumption:
		return "No"
	case MoreInformationRequired:
		return "More Information Required"
	default:
		return fmt.Sprintf("%v", int(consumption))
	}
}

func (formula Formula) String() string {
	outputString := "Formula:\n"

	for _, nutrient := range formula.Nutrients {
		outputString += "\tNutrient:\n"
		outputString += fmt.Sprintf("\t\tName: %v", nutrient.Name)
		outputString += fmt.Sprintf("\tChemicalFormula: %v\n", nutrient.ChemicalFormula)
		outputString += fmt.Sprintf("\t\tAmount: %v %v/Liter\n", nutrient.Value, nutrient.Units.String())
		outputString += fmt.Sprintf("\t\tDescription: %v\n", insertNewlines(nutrient.Description, 90, 2))
		if nutrient.Safe != SafeForConsumption {
			outputString += fmt.Sprintf("\t\tSafe: %v %v\n", nutrient.Safe.String(), nutrient.SafeNotes)
		}
	}
	return outputString
}

func (formula Formula) CalculateVolumeString(volume float64) string {
	outputString := "Media Amounts:\n"
	for _, nutrient := range formula.Nutrients {
		outputString += "\tNutrient:\n"
		outputString += fmt.Sprintf("\t\tName: %v", nutrient.Name)
		outputString += fmt.Sprintf("\t\tAmount: %v %v\n", nutrient.Value * volume, nutrient.Units.String())
	}
	return outputString
}

func NewFormula(nutrients []Nutrient) Formula {
	return Formula{Nutrients: nutrients}
}

func NewZarouckFormula() Formula {
	return NewFormula(
		[]Nutrient{
			{
				Name:            "Sodium Chloride",
				ChemicalFormula: "NaCl",
				Value:           1.0,
				Description:     "Table Salt",
				Safe:            SafeForConsumption,
				Units:           Grams,
			},
			{
				Name:            "Calcium chloride dihydrate",
				ChemicalFormula: "CaCl₂.2H₂O",
				Value:           0.04,
				Description:     "Calcium chloride dihydrate is a hydrate that is the dihydrate form of calcium chloride. It is a hydrate, a calcium salt and an inorganic chloride. It contains a calcium dichloride.",
				Safe:            SafeForConsumption,
				Units:           Grams,
			},
			{
				Name:            "Potassium nitrate",
				ChemicalFormula: "KNO₃",
				Value:           0.0,
				Description:     "Potassium nitrate (KNO3) is a soluble source of two major essential plant nutrients. It is commonly used as a fertilizer for high-value crops that benefit from nitrate (NO3-) nutrition and a source of potassium (K+) free of chloride (Cl-). Not used in Zarrouk's media",
				Safe:            SafeForConsumption,
				Units:           Grams,
			},
			{
				Name:            "Sodium Nitrate",
				ChemicalFormula: "NaNO₃",
				Value:           2.5,
				Description:     "Sodium nitrate is a white deliquescent solid very soluble in water. It is a readily available source of the nitrate anion (NO3−), which is useful in several reactions carried out on industrial scales for the production of fertilizers, pyrotechnics, smoke bombs and other explosives, glass and pottery enamels, food preservatives (esp. meats), and solid rocket propellant.",
				Safe:            SafeForConsumption,
				Units:           Grams,
			},
			{
				Name:            "Ferrous sulfate heptahydrate",
				ChemicalFormula: "FeSO₄.7H₂O",
				Value:           0.01,
				Description:     "Iron(2+) sulfate heptahydrate is a hydrate that is the heptahydrate form of iron(2+) sulfate. It is used as a source of iron in the treatment of iron-deficiency anaemia (generally in liquid-dosage treatments; for solid-dosage treatments, the monohydrate is normally used). It has a role as a nutraceutical, an anti-anaemic agent and a reducing agent. It is a hydrate and an iron molecular entity. It contains an iron(2+) sulfate (anhydrous).",
				Safe:            SafeForConsumption,
				Units:           Grams,
			},
			{
				Name:            "Ethylenediaminetetraacetic acid",
				ChemicalFormula: "EDTA(Na)",
				Value:           0.08,
				Description:     "EDTA is widely used in industry. It also has applications in food preservation, medicine, cosmetics, water softening, in laboratories, and other fields.",
				Safe:            MoreInformationRequired,
				SafeNotes:       "Safe dose is 700 - 3500 mg intravenous every 12 hours  https://web.archive.org/web/20070504081119/http://www.umm.edu/altmed/articles/ethylenediaminetetraacetic-acid-000302.htm",
				Units:           Grams,
			},
			{
				Name:            "Potassium sulfate",
				ChemicalFormula: "K₂SO₄",
				Value:           1.0,
				Description:     "Potassium sulfate (US) or potassium sulphate (UK), also called sulphate of potash (SOP), arcanite, or archaically potash of sulfur, is the inorganic compound with formula K2SO4, a white water-soluble solid. It is commonly used in fertilizers, providing both potassium and sulfur.",
				Safe:            SafeForConsumption,
				SafeNotes:       "Although ingestion is not thought to produce harmful effects, the material may still be damaging to the health of the individual following ingestion, especially where pre-existing organ (e.g. liver, kidney) damage is evident. Present definitions of harmful or toxic substances are generally based on doses producing mortality (death) rather than those producing morbidity (disease, ill-health). Gastrointestinal tract discomfort may produce nausea and vomiting. In an occupational setting however, ingestion of insignificant quantities is not thought to be cause for concern.",
				Units:           Grams,
			},
			{
				Name:            "Magnesium sulfate heptahydrate (Epsom Salt)",
				ChemicalFormula: "MgSO₄.7H₂O",
				Value:           0.2,
				Description:     "Magnesium sulfate heptahydrate is a hydrate that is the heptahydrate form of magnesium sulfate. It has a role as a laxative and a cathartic. It is a magnesium salt and a hydrate. It contains a magnesium sulfate.",
				Safe:            SafeForConsumption,
				SafeNotes:       "Mild laxitive and pain reliever",
				Units:           Grams,
			},
			{
				Name:            "Sodium Bicarbonate",
				ChemicalFormula: "NaHCO₃",
				Value:           16.8,
				Description:     "Sodium bicarbonate appears as odorless white crystalline powder or lumps. Slightly alkaline (bitter) taste. pH (of freshly prepared 0.1 molar aqueous solution): 8.3 at 77 °F. pH (of saturated solution): 8-9. Non-toxic.",
				Safe:            SafeForConsumption,
				Units:           Grams,
			},
			{
				Name:            "Potassium Phosphate, Dibasic",
				ChemicalFormula: "K₂HPO₄",
				Value:           0.5,
				Description:     "Dipotassium hydrogen phosphate is a potassium salt that is the dipotassium salt of phosphoric acid. It has a role as a buffer. It is a potassium salt and an inorganic phosphate. Dipotassium phosphate (K2HPO4) is a highly water-soluble salt often used as a fertilizer and food additive as a source of phosphorus and potassium as well as a buffering agent. Potassium Phosphate, Dibasic is the dipotassium form of phosphoric acid, that can be used as an electrolyte replenisher and with radio-protective activity. Upon oral administration, potassium phosphate is able to block the uptake of the radioactive isotope phosphorus P 32 (P-32).",
				Safe:            SafeForConsumption,
				SafeNotes:       "Used as an electorlyte replenisher",
				Units:           Grams,
			},
			{
				Name:            "Vitamin A₅",
				ChemicalFormula: "H₃BO₃,MnCl₂.4H₂O,ZnSO₄.4H2ONa₂MoO₄,CuSO₄.5H₂O",
				Value:           1.0,
				Description:     "A new vitamin concept, termed vitamin A5, an umbrella term for vitamin A derivatives being direct nutritional precursors for 9-cis-13,14-dihydroretinoic acid and further induction of RXR-signaling, was recently identified with global importance for mental health and healthy brain and nerve functions. Dietary recommendations in the range of 1.1 (0.5–1.8) mg vitamin A5 / day were suggested by an international expert consortium.",
				Safe:            SafeForConsumption,
				Units:           Milliliters,
			},
		},
	)
}

func insertNewlines(text string, interval, tabDepth int) string {
	var result strings.Builder
	for i, r := range text {
		result.WriteRune(r)
		if (i+1)%interval == 0 {
			result.WriteRune('\n')
			result.WriteString(strings.Repeat("\t", tabDepth))
		}
	}
	return result.String()
}

func main() {
	var opts struct {
		Volume []float64 `short:"o" long:"volume" required:"true" description:"The volume of liquid in liters, that will be used to created the media.  Note: The volume will increase with the nutrients so choose a value smaller than the size of your bioreactor"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}
	zarrouk := NewZarouckFormula()
	fmt.Printf("%v\n", zarrouk.CalculateVolumeString(opts.Volume[0]))
}

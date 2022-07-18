package template

// Alert is a page alert.
type Alert struct {
	Header string
	Level  string
	Text   string
}

// AlertBars is a set of alerts to display together.
type AlertBars *[]Alert

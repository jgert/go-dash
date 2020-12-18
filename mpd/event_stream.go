package mpd

type EventStream struct {
	SchemeIdUri            *string  `xml:"schemeIdUri,attr"`
	Value                  *string  `xml:"value,attr,omitempty"`
	Timescale              *int64   `xml:"timescale,attr"`
	PresentationTimeOffset *uint64  `xml:"presentationTimeOffset,attr,omitempty"`
	Events                 []*Event `xml:"Event,omitempty"`
}

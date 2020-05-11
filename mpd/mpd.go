package mpd

// Type definition for DASH profiles
type DashProfile string

// Constants for supported DASH profiles
const (
	// Live Profile
	DASH_PROFILE_LIVE DashProfile = "urn:mpeg:dash:profile:isoff-live:2011"
	// On Demand Profile
	DASH_PROFILE_ONDEMAND DashProfile = "urn:mpeg:dash:profile:isoff-on-demand:2011"
	// HbbTV Profile
	DASH_PROFILE_HBBTV_1_5_LIVE DashProfile = "urn:hbbtv:dash:profile:isoff-live:2012,urn:mpeg:dash:profile:isoff-live:2011"
)

type AudioChannelConfigurationScheme string

const (
	// Scheme for non-Dolby Audio
	AUDIO_CHANNEL_CONFIGURATION_MPEG_DASH AudioChannelConfigurationScheme = "urn:mpeg:dash:23003:3:audio_channel_configuration:2011"
	// Scheme for Dolby Audio
	AUDIO_CHANNEL_CONFIGURATION_MPEG_DOLBY AudioChannelConfigurationScheme = "tag:dolby.com,2014:dash:audio_channel_configuration:2011"
)

// Constants for some known MIME types, this is a limited list and others can be used.
const (
	DASH_MIME_TYPE_VIDEO_MP4     string = "video/mp4"
	DASH_MIME_TYPE_AUDIO_MP4     string = "audio/mp4"
	DASH_MIME_TYPE_SUBTITLE_VTT  string = "text/vtt"
	DASH_MIME_TYPE_SUBTITLE_TTML string = "application/ttaf+xml"
	DASH_MIME_TYPE_SUBTITLE_SRT  string = "application/x-subrip"
	DASH_MIME_TYPE_SUBTITLE_DFXP string = "application/ttaf+xml"
)

type MPD struct {
	XMLNs                      *string     `xml:"xmlns,attr,omitempty"`
	NSCENC                     *string     `xml:"xmlns:cenc,attr,omitempty"`
	Profiles                   *string     `xml:"profiles,attr,omitempty"`
	Type                       *string     `xml:"type,attr,omitempty"`
	MediaPresentationDuration  *string     `xml:"mediaPresentationDuration,attr,omitempty"`
	MinBufferTime              *string     `xml:"minBufferTime,attr,omitempty"`
	AvailabilityStartTime      *string     `xml:"availabilityStartTime,attr,omitempty"`
	MinimumUpdatePeriod        *string     `xml:"minimumUpdatePeriod,attr,omitempty"`
	PublishTime                *string     `xml:"publishTime,attr,omitempty"`
	TimeShiftBufferDepth       *string     `xml:"timeShiftBufferDepth,attr,omitempty"`
	MaxSegmentDuration         *string     `xml:"maxSegmentDuration,attr,omitempty"`
	BaseURL                    string      `xml:"BaseURL,omitempty"`
	ID                         *string     `xml:"id,attr,omitempty"`
	SuggestedPresentationDelay *string     `xml:"suggestedPresentationDelay,attr,omitempty"`
	Periods                    []*Period   `xml:"Period,omitempty"`
	UTCTiming                  *Descriptor `xml:"UTCTiming,omitempty"`
}

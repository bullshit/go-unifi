// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ fmt.Formatter
	_ context.Context
)

type Hotspot2Conf struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	ApIsolate             bool                                 `json:"ap_isolate"`
	Capab                 []Hotspot2Conf_Capab                 `json:"capab,omitempty"`
	CellularNetworkList   []Hotspot2Conf_CellularNetworkList   `json:"cellular_network_list,omitempty"`
	DeauthReqTimeout      int                                  `json:"deauth_req_timeout,omitempty"` // [1-9][0-9]|[1-9][0-9][0-9]|[1-2][0-9][0-9][0-9]|3[0-5][0-9][0-9]|3600
	DisableDgaf           bool                                 `json:"disable_dgaf"`
	DomainNameList        []string                             `json:"domain_name_list,omitempty"` // .{1,128}
	FriendlyName          []Hotspot2Conf_FriendlyName          `json:"friendly_name,omitempty"`
	GasComebackDelay      int                                  `json:"gas_comeback_delay,omitempty"` // [1-9][0-9]|[1-9][0-9][0-9]|[1-2][0-9][0-9][0-9]|3[0-5][0-9][0-9]|300
	Hessid                string                               `json:"hessid"`                       // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$|^$
	Icons                 []Hotspot2Conf_Icons                 `json:"icons,omitempty"`
	MetricsDownlinkLoad   int                                  `json:"metrics_downlink_load,omitempty"`
	MetricsDownlinkSpeed  int                                  `json:"metrics_downlink_speed,omitempty"`
	MetricsInfo           string                               `json:"metrics_info,omitempty"` // [0-9A-Fa-f]{1,2}
	MetricsMeasurement    int                                  `json:"metrics_measurement,omitempty"`
	MetricsStatus         bool                                 `json:"metrics_status"`
	MetricsUplinkLoad     int                                  `json:"metrics_uplink_load,omitempty"`
	MetricsUplinkSpeed    int                                  `json:"metrics_uplink_speed,omitempty"`
	NaiRealmList          []Hotspot2Conf_NaiRealmList          `json:"nai_realm_list,omitempty"`
	Name                  string                               `json:"name,omitempty"` // .{1,128}
	NetworkAccessAsra     bool                                 `json:"network_access_asra"`
	NetworkAccessEsr      bool                                 `json:"network_access_esr"`
	NetworkAccessInternet bool                                 `json:"network_access_internet"`
	NetworkAccessUesa     bool                                 `json:"network_access_uesa"`
	NetworkType           int                                  `json:"network_type,omitempty"` // 0|1|2|3|4|5|14|15
	Osu                   []Hotspot2Conf_Osu                   `json:"osu,omitempty"`
	OsuSSID               string                               `json:"osu_ssid"`
	P2P                   bool                                 `json:"p2p"`
	ProxyArp              bool                                 `json:"proxy_arp"`
	QOSMapDcsp            []Hotspot2Conf_QOSMapDcsp            `json:"qos_map_dcsp,omitempty"`
	QOSMapExceptions      []Hotspot2Conf_QOSMapExceptions      `json:"qos_map_exceptions,omitempty"`
	QOSMapStatus          bool                                 `json:"qos_map_status"`
	RoamingConsortiumList []Hotspot2Conf_RoamingConsortiumList `json:"roaming_consortium_list,omitempty"`
	TCFilename            string                               `json:"t_c_filename,omitempty"` // .{1,256}
	TCTimestamp           int                                  `json:"t_c_timestamp,omitempty"`
	VenueGroup            int                                  `json:"venue_group,omitempty"` // 0|1|2|3|4|5|6|7|8|9|10|11
	VenueName             []Hotspot2Conf_VenueName             `json:"venue_name,omitempty"`
	VenueType             int                                  `json:"venue_type,omitempty"` // 0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15
}

type Hotspot2Conf_Capab struct {
	Port     string `json:"port,omitempty"`     // (([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])|([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])-([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]))+(,([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])|,([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])-([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])){0,14}
	Protocol string `json:"protocol,omitempty"` // icmp|tcp_udp|tcp|udp
	Status   string `json:"status,omitempty"`   // closed|open|unknown
}

type Hotspot2Conf_CellularNetworkList struct {
	Mcc  int    `json:"mcc,omitempty"`
	Mnc  int    `json:"mnc,omitempty"`
	Name string `json:"name,omitempty"` // .{1,128}
}

type Hotspot2Conf_Description struct {
	Language string `json:"language,omitempty"` // [a-z]{3}
	Text     string `json:"text,omitempty"`     // .{1,128}
}

type Hotspot2Conf_FriendlyName struct {
	Language string `json:"language,omitempty"` // [a-z]{3}
	Text     string `json:"text,omitempty"`     // .{1,128}
}

type Hotspot2Conf_Icon struct {
	Name string `json:"name,omitempty"` // .{1,128}
}

type Hotspot2Conf_Icons struct {
	Data     string `json:"data,omitempty"`
	Filename string `json:"filename,omitempty"` // .{1,256}
	Height   int    `json:"height,omitempty"`
	Language string `json:"language,omitempty"` // [a-z]{3}
	Media    string `json:"media,omitempty"`    // .{1,256}
	Name     string `json:"name,omitempty"`     // .{1,256}
	Size     int    `json:"size,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type Hotspot2Conf_NaiRealmList struct {
	AuthIDs   string `json:"auth_ids,omitempty"`
	AuthVals  string `json:"auth_vals,omitempty"`
	EapMethod int    `json:"eap_method,omitempty"` // 13|21|18|23|50
	Encoding  int    `json:"encoding,omitempty"`   // 0|1
	Name      string `json:"name,omitempty"`       // .{1,128}
	Status    bool   `json:"status"`
}

type Hotspot2Conf_Osu struct {
	Description      []Hotspot2Conf_Description  `json:"description,omitempty"`
	FriendlyName     []Hotspot2Conf_FriendlyName `json:"friendly_name,omitempty"`
	Icon             []Hotspot2Conf_Icon         `json:"icon,omitempty"`
	MethodOmaDm      bool                        `json:"method_oma_dm"`
	MethodSoapXmlSpp bool                        `json:"method_soap_xml_spp"`
	Nai              string                      `json:"nai,omitempty"`
	Nai2             string                      `json:"nai2,omitempty"`
	OperatingClass   string                      `json:"operating_class,omitempty"` // [0-9A-Fa-f]{12}
	ServerUri        string                      `json:"server_uri,omitempty"`
}

type Hotspot2Conf_QOSMapDcsp struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}

type Hotspot2Conf_QOSMapExceptions struct {
	Dcsp int `json:"dcsp,omitempty"`
	Up   int `json:"up,omitempty"` // [0-7]
}

type Hotspot2Conf_RoamingConsortiumList struct {
	Name string `json:"name,omitempty"` // .{1,128}
	Oid  string `json:"oid,omitempty"`  // .{1,128}
}

type Hotspot2Conf_VenueName struct {
	Language string `json:"language,omitempty"` // [a-z]{3}
	Name     string `json:"name,omitempty"`     // .{1,128}
}

func (c *Client) listHotspot2Conf(ctx context.Context, site string) ([]Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/hotspot2conf", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getHotspot2Conf(ctx context.Context, site, id string) (*Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/hotspot2conf/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteHotspot2Conf(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/hotspot2conf/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createHotspot2Conf(ctx context.Context, site string, d *Hotspot2Conf) (*Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/hotspot2conf", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateHotspot2Conf(ctx context.Context, site string, d *Hotspot2Conf) (*Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/hotspot2conf/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
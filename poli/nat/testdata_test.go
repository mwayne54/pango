package nat

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	version version.Number
	desc    string
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{version.Number{5, 0, 0, ""}, "dst only", Entry{
			Name:                 "nat policy",
			Description:          "my nat policy",
			Type:                 "ipv4",
			SourceZones:          []string{"zone1", "zone2"},
			DestinationZone:      "zone3",
			ToInterface:          "ethernet1/7",
			Service:              "myService",
			SourceAddresses:      []string{"any"},
			DestinationAddresses: []string{"any"},
			SatType:              None,
			DatType:              DatTypeStatic,
			DatAddress:           "10.5.1.1",
			DatPort:              1234,
			Tags:                 []string{"tag1", "tag2"},
		}},
		{version.Number{5, 0, 0, ""}, "dynamic ip and port with translated address", Entry{
			Name:                   "nat policy",
			Description:            "my nat policy",
			Type:                   "ipv4",
			SourceZones:            []string{"zone3"},
			DestinationZone:        "zone4",
			ToInterface:            "any",
			Service:                "any",
			SourceAddresses:        []string{"any"},
			DestinationAddresses:   []string{"any"},
			SatType:                DynamicIpAndPort,
			SatAddressType:         TranslatedAddress,
			SatTranslatedAddresses: []string{"10.3.1.1", "10.3.2.1"},
		}},
		{version.Number{5, 0, 0, ""}, "dynamic ip with interface address fallback", Entry{
			Name:                   "nat policy",
			Description:            "my nat policy",
			Type:                   "ipv4",
			SourceZones:            []string{"zone3"},
			DestinationZone:        "zone4",
			ToInterface:            "any",
			Service:                "any",
			SourceAddresses:        []string{"any"},
			DestinationAddresses:   []string{"any"},
			SatType:                DynamicIp,
			SatTranslatedAddresses: []string{"10.5.5.5", "10.6.6.6"},
			SatFallbackType:        InterfaceAddress,
			SatFallbackInterface:   "ethernet1/7",
			SatFallbackIpType:      Ip,
			SatFallbackIpAddress:   "10.10.10.10",
		}},
		{version.Number{5, 0, 0, ""}, "static ip with target", Entry{
			Name:                       "nat policy",
			Description:                "my nat policy",
			Type:                       "ipv4",
			SourceZones:                []string{"zone3"},
			DestinationZone:            "zone4",
			ToInterface:                "any",
			Service:                    "any",
			SourceAddresses:            []string{"any"},
			DestinationAddresses:       []string{"any"},
			SatType:                    StaticIp,
			SatStaticTranslatedAddress: "10.5.5.5",
			SatStaticBiDirectional:     true,
			Targets: map[string][]string{
				"fw1": nil,
				"fw2": {"vsys2", "vsys3"},
			},
			NegateTarget: true,
		}},
		{version.Number{8, 1, 0, ""}, "v2 normal dat", Entry{
			Name:                 "nat policy",
			Description:          "my nat policy",
			Type:                 "ipv4",
			SourceZones:          []string{"zone1", "zone2"},
			DestinationZone:      "zone3",
			ToInterface:          "ethernet1/7",
			Service:              "myService",
			SourceAddresses:      []string{"any"},
			DestinationAddresses: []string{"any"},
			SatType:              None,
			DatType:              DatTypeStatic,
			DatAddress:           "10.5.1.1",
			DatPort:              1234,
			Tags:                 []string{"tag1", "tag2"},
		}},
		{version.Number{8, 1, 0, ""}, "v2 dynamic dat", Entry{
			Name:                   "nat policy",
			Description:            "my nat policy",
			Type:                   "ipv4",
			SourceZones:            []string{"zone1", "zone2"},
			DestinationZone:        "zone3",
			ToInterface:            "ethernet1/7",
			Service:                "myService",
			SourceAddresses:        []string{"any"},
			DestinationAddresses:   []string{"any"},
			SatType:                None,
			DatType:                DatTypeDynamic,
			DatAddress:             "my fqdn object",
			DatPort:                1234,
			DatDynamicDistribution: "round-robin",
			Tags:                   []string{"tag1", "tag2"},
		}},
	}
}

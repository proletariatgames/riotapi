package riotapi

import (
	"strings"
)

// Platform struct describes a Riot Platform
type Platform struct {
	Id       string // e.g. "NA1"
	Name     string // e.g. "North America"
	RegionId string // e.g. "NA"
}

var Platforms = []Platform{
	{"NA1", "North America", "NA"},
	{"BR1", "Brazil", "BR"},
	{"EUN1", "EU Nordic & East ", "EUNE"},
	{"EUW1", "EU West", "EUW"},
	{"JP1", "Japan", "JP"},
	{"KR", "Korea", "KR"},
	{"LA1", "Latin America North", "LAN"},
	{"LA2", "Latin America South", "LAS"},
	{"OC1", "Oceania", "OCE"},
	{"PBE1", "PBE", "PBE"},
	{"RU", "Russia", "RU"},
	{"TR1", "Turkey", "TR"},
}

// GetPlatform returns the platform description for a given platform id (case-insensitive)
func GetPlatform(id string) *Platform {
	upperId := strings.ToUpper(id)
	for _, p := range Platforms {
		if upperId == p.Id {
			return &p
		}
	}
	return nil
}

// GetPlatform returns the platform description for a given platform id (case-insensitive)
func GetPlatformByRegion(region string) *Platform {
	upperRegion := strings.ToUpper(region)
	for _, p := range Platforms {
		if upperRegion == p.RegionId {
			return &p
		}
	}
	return nil
}

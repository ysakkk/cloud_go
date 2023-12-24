package cmd

type clb_describe_loadbalancers struct {
	TotalCount      int `json:"TotalCount"`
	LoadBalancerSet []struct {
		LoadBalancerId   string   `json:"LoadBalancerId"`
		LoadBalancerName string   `json:"LoadBalancerName"`
		LoadBalancerType string   `json:"LoadBalancerType"`
		//Forward          int      `json:"Forward"`
		Domain           string   `json:"Domain"`
		//LoadBalancerVips []any    `json:"LoadBalancerVips"`
		//Status           int      `json:"Status"`
		//CreateTime       string   `json:"CreateTime"`
		//StatusTime       string   `json:"StatusTime"`
		//ProjectId        int      `json:"ProjectId"`
		VpcId            string   `json:"VpcId"`
		//OpenBgp          int      `json:"OpenBgp"`
		//Snat             bool     `json:"Snat"`
		//Isolation        int      `json:"Isolation"`
		//Log              string   `json:"Log"`
		//SubnetId         string   `json:"SubnetId"`
		//Tags             []any    `json:"Tags"`
		SecureGroups     []string `json:"SecureGroups"`
		TargetRegionInfo struct {
			Region string `json:"Region"`
			VpcId  string `json:"VpcId"`
		} `json:"TargetRegionInfo"`
		//AnycastZone      string `json:"AnycastZone"`
		//AddressIPVersion string `json:"AddressIPVersion"`
		//NumericalVpcId   int    `json:"NumericalVpcId"`
		//VipIsp           string `json:"VipIsp"`
		MasterZone       struct {
			ZoneId     int    `json:"ZoneId"`
			Zone       string `json:"Zone"`
			ZoneName   string `json:"ZoneName"`
			ZoneRegion string `json:"ZoneRegion"`
			LocalZone  bool   `json:"LocalZone"`
			EdgeZone   bool   `json:"EdgeZone"`
		} `json:"MasterZone"`
		//BackupZoneSet     any    `json:"BackupZoneSet"`
		//IsolatedTime      any    `json:"IsolatedTime"`
		//ExpireTime        string `json:"ExpireTime"`
		//ChargeType        string `json:"ChargeType"`
		//NetworkAttributes struct {
		//	InternetChargeType      string `json:"InternetChargeType"`
		//	InternetMaxBandwidthOut int    `json:"InternetMaxBandwidthOut"`
		//	BandwidthpkgSubType     any    `json:"BandwidthpkgSubType"`
		//} `json:"NetworkAttributes"`
		//PrepaidAttributes        any    `json:"PrepaidAttributes"`
		//LogSetId                 string `json:"LogSetId"`
		//LogTopicId               string `json:"LogTopicId"`
		//AddressIPv6              string `json:"AddressIPv6"`
		//ExtraInfo                any    `json:"ExtraInfo"`
		//IsDDos                   bool   `json:"IsDDos"`
		//ConfigId                 string `json:"ConfigId"`
		LoadBalancerPassToTarget bool   `json:"LoadBalancerPassToTarget"`
		//ExclusiveCluster         struct {
		//	L4Clusters       any `json:"L4Clusters"`
		//	L7Clusters       any `json:"L7Clusters"`
		//	ClassicalCluster any `json:"ClassicalCluster"`
		//} `json:"ExclusiveCluster"`
		//IPv6Mode           any      `json:"IPv6Mode"`
		//SnatPro            bool     `json:"SnatPro"`
		//SnatIps            []any    `json:"SnatIps"`
		//SLAType            string   `json:"SlaType"`
		//IsBlock            bool     `json:"IsBlock"`
		//IsBlockTime        string   `json:"IsBlockTime"`
		//LocalBgp           bool     `json:"LocalBgp"`
		//ClusterTag         any      `json:"ClusterTag"`
		//MixIPTarget        bool     `json:"MixIpTarget"`
		//Zones              any      `json:"Zones"`
		//NfvInfo            string   `json:"NfvInfo"`
		//HealthLogSetId     string   `json:"HealthLogSetId"`
		//HealthLogTopicId   string   `json:"HealthLogTopicId"`
		//ClusterIds         any      `json:"ClusterIds"`
		//AttributeFlags     []string `json:"AttributeFlags"`
		//LoadBalancerDomain string   `json:"LoadBalancerDomain"`
		//Egress             string   `json:"Egress"`
	} `json:"LoadBalancerSet"`
	//RequestId string `json:"RequestId"`
}

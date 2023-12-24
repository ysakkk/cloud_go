package cmd

type ec2_describe_instances struct {
	Reservations []struct {
		Groups    []any `json:"Groups"`
		Instances []struct {
			//AmiLaunchIndex int       `json:"AmiLaunchIndex"`
			//ImageId        string    `json:"ImageId"`
			InstanceId     string    `json:"InstanceId"`
			InstanceType   string    `json:"InstanceType"`
			//KeyName        string    `json:"KeyName"`
			//LaunchTime     string `json:"LaunchTime"`
			//Monitoring     struct {
				//State string `json:"State"`
			//} `json:"Monitoring"`
			Placement struct {
				AvailabilityZone string `json:"AvailabilityZone"`
				//GroupName        string `json:"GroupName"`
				//Tenancy          string `json:"Tenancy"`
			} `json:"Placement"`
			//PrivateDNSName   string `json:"PrivateDnsName"`
			PrivateIPAddress string `json:"PrivateIpAddress"`
			//ProductCodes     []struct {
			//	ProductCodeId   string `json:"ProductCodeId"`
			//	ProductCodeType string `json:"ProductCodeType"`
			//} `json:"ProductCodes"`
			// PublicDNSName   string `json:"PublicDnsName"`
			PublicIPAddress string `json:"PublicIpAddress"`
			State           struct {
				// Code int    `json:"Code"`
				Name string `json:"Name"`
			} `json:"State"`
			// StateTransitionReason string `json:"StateTransitionReason"`
			SubnetId              string `json:"SubnetId"`
			VpcId                 string `json:"VpcId"`
			Architecture          string `json:"Architecture"`
			BlockDeviceMappings   []struct {
				DeviceName string `json:"DeviceName"`
				Ebs        struct {
					AttachTime          string `json:"AttachTime"`
					DeleteOnTermination bool      `json:"DeleteOnTermination"`
					Status              string    `json:"Status"`
					VolumeId            string    `json:"VolumeId"`
				} `json:"Ebs"`
			} `json:"BlockDeviceMappings"`
			InstanceLifecycle string `json:"InstanceLifecycle"`
			// ClientToken        string `json:"ClientToken"`
			// EbsOptimized       bool   `json:"EbsOptimized"`
			// Hypervisor         string `json:"Hypervisor"`
			// IamInstanceProfile struct {
			// 	Arn string `json:"Arn"`
			// 	Id  string `json:"Id"`
			// } `json:"IamInstanceProfile"`
			// NetworkInterfaces []struct {
			// 	Association struct {
			// 		IPOwnerId     string `json:"IpOwnerId"`
			// 		PublicDNSName string `json:"PublicDnsName"`
			// 		PublicIP      string `json:"PublicIp"`
			// 	} `json:"Association"`
			// 	Attachment struct {
			// 		AttachTime          string `json:"AttachTime"`
			// 		AttachmentId        string    `json:"AttachmentId"`
			// 		DeleteOnTermination bool      `json:"DeleteOnTermination"`
			// 		DeviceIndex         int       `json:"DeviceIndex"`
			// 		Status              string    `json:"Status"`
			// 		NetworkCardIndex    int       `json:"NetworkCardIndex"`
			// 	} `json:"Attachment"`
			// 	Description string `json:"Description"`
			// 	Groups      []struct {
			// 		GroupName string `json:"GroupName"`
			// 		GroupId   string `json:"GroupId"`
			// 	} `json:"Groups"`
			// 	Ipv6Addresses      []any  `json:"Ipv6Addresses"`
			// 	MacAddress         string `json:"MacAddress"`
			// 	NetworkInterfaceId string `json:"NetworkInterfaceId"`
			// 	OwnerId            string `json:"OwnerId"`
			// 	PrivateDNSName     string `json:"PrivateDnsName"`
			// 	PrivateIPAddress   string `json:"PrivateIpAddress"`
			// 	PrivateIPAddresses []struct {
			// 		Association struct {
			// 			IPOwnerId     string `json:"IpOwnerId"`
			// 			PublicDNSName string `json:"PublicDnsName"`
			// 			PublicIP      string `json:"PublicIp"`
			// 		} `json:"Association"`
			// 		Primary          bool   `json:"Primary"`
			// 		PrivateDNSName   string `json:"PrivateDnsName"`
			// 		PrivateIPAddress string `json:"PrivateIpAddress"`
			// 	} `json:"PrivateIpAddresses"`
			// 	SourceDestCheck bool   `json:"SourceDestCheck"`
			// 	Status          string `json:"Status"`
			// 	SubnetId        string `json:"SubnetId"`
			// 	VpcId           string `json:"VpcId"`
			// 	InterfaceType   string `json:"InterfaceType"`
			// } `json:"NetworkInterfaces"`
			RootDeviceName string `json:"RootDeviceName"`
			RootDeviceType string `json:"RootDeviceType"`
			SecurityGroups []struct {
				GroupName string `json:"GroupName"`
				GroupId   string `json:"GroupId"`
			} `json:"SecurityGroups"`
			SourceDestCheck bool `json:"SourceDestCheck"`
			Tags            []struct {
				Key   string `json:"Key"`
				Value string `json:"Value"`
			} `json:"Tags"`
			VirtualizationType string `json:"VirtualizationType"`
			CPUOptions         struct {
				CoreCount      int `json:"CoreCount"`
				ThreadsPerCore int `json:"ThreadsPerCore"`
			} `json:"CpuOptions"`
			CapacityReservationSpecification struct {
				CapacityReservationPreference string `json:"CapacityReservationPreference"`
			} `json:"CapacityReservationSpecification"`
			HibernationOptions struct {
				Configured bool `json:"Configured"`
			} `json:"HibernationOptions"`
			// MetadataOptions struct {
			// 	State                   string `json:"State"`
			// 	HTTPTokens              string `json:"HttpTokens"`
			// 	HTTPPutResponseHopLimit int    `json:"HttpPutResponseHopLimit"`
			// 	HTTPEndpoint            string `json:"HttpEndpoint"`
			// 	HTTPProtocolIpv6        string `json:"HttpProtocolIpv6"`
			// 	InstanceMetadataTags    string `json:"InstanceMetadataTags"`
			// } `json:"MetadataOptions"`
			EnclaveOptions struct {
				Enabled bool `json:"Enabled"`
			} `json:"EnclaveOptions"`
			PlatformDetails          string    `json:"PlatformDetails"`
			UsageOperation           string    `json:"UsageOperation"`
			UsageOperationUpdateTime string `json:"UsageOperationUpdateTime"`
			PrivateDNSNameOptions    struct {
			} `json:"PrivateDnsNameOptions"`
			MaintenanceOptions struct {
				AutoRecovery string `json:"AutoRecovery"`
			} `json:"MaintenanceOptions"`
		} `json:"Instances"`
		OwnerId       string `json:"OwnerId"`
		ReservationId string `json:"ReservationId"`
	} `json:"Reservations"`
}



type ec2_describe_volumes struct {
	Volumes []struct {
		Attachments []struct {
			AttachTime          string `json:"AttachTime"`
			Device              string    `json:"Device"`
			InstanceId          string    `json:"InstanceId"`
			State               string    `json:"State"`
			VolumeId            string    `json:"VolumeId"`
			DeleteOnTermination bool      `json:"DeleteOnTermination"`
		} `json:"Attachments"`
		AvailabilityZone   string    `json:"AvailabilityZone"`
		CreateTime         string `json:"CreateTime"`
		Encrypted          bool      `json:"Encrypted"`
		Size               int       `json:"Size"`
		SnapshotId         string    `json:"SnapshotId"`
		State              string    `json:"State"`
		VolumeId           string    `json:"VolumeId"`
		Iops               int       `json:"Iops"`
		VolumeType         string    `json:"VolumeType"`
		MultiAttachEnabled bool      `json:"MultiAttachEnabled"`
		Tags               []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"Tags,omitempty"`
		Throughput int `json:"Throughput,omitempty"`
	} `json:"Volumes"`
}

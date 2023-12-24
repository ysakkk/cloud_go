package cmd

import (
	"context"
	"os"
	//dataframe "github.com/rocketlaunchr/dataframe-go"
	// "github.com/rocketlaunchr/dataframe-go/exports"
	"github.com/rocketlaunchr/dataframe-go/imports"

	"encoding/json"
	//"time"
	//"sort"
	"strings"
	//"reflect"
	//"io/ioutil"
	//"bytes"
	//"os/exec"
	"fmt"
	//"os"
	"github.com/spf13/cobra"
	//toml "github.com/BurntSushi/toml"
	//toml "github.com/pelletier/go-toml/v2"
	//"gopkg.in/ini.v1"
)

var ec2 = &cobra.Command{
	Use: "ec2",
	Run: func(cmd *cobra.Command, args []string) {

		if profile == "" {
			fmt.Println("Availabe Profiles:")
			listAWSProfiles()
			return
		}

		j := ec2_describe_instances{}
		result := execCommand("aws", "--profile", profile, "ec2", "describe-instances")
		err := json.Unmarshal([]byte(result), &j)
		//fmt.Println(result)
		b,_ := json.Marshal(j)
		fmt.Println(string(b))
		os.Exit(1)
		if err != nil {
			fmt.Println(err)
			return
		}
		var csv_data string
		csv_data = fmt.Sprintf(`"%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s"`,
			"Name", "State", "InstanceId", "InstanceType", "Spot", "vCPU", "Volumes", "AZ", "VPC", "Subnet", "PrivateIP", "PublicIP", "SecurityGroup", "Tags")
		csv_data += "\n"

		for _, v := range j.Reservations {
			var name string
			var tags string
			var sg string
			var volumes string
			i := v.Instances[0]
			for _, v2 := range i.Tags {
				if v2.Key == "Name" {
					name = v2.Value
				} else {
					//tags += fmt.Sprintf("%s:%s\n", v2.Key, v2.Value)
					tags += fmt.Sprintf("%s:%s,", v2.Key, v2.Value)
				}
			}
			for _, v2 := range i.SecurityGroups {
				//sg += fmt.Sprintf("%s:%s\n", v2.GroupName, v2.GroupId)
				sg += fmt.Sprintf("%s:%s,", v2.GroupName, v2.GroupId)
			}
			for _, v2 := range i.BlockDeviceMappings {
				//volumes += fmt.Sprintf("%s:%s\n", v2.DeviceName, v2.Ebs.VolumeId)
				volumes += fmt.Sprintf("%s:%s,", v2.DeviceName, v2.Ebs.VolumeId)
			}
			tags = strings.TrimRight(tags, "\n")
			sg = strings.TrimRight(sg, "\n")
			//fmt.Printf("%40v %20v %15v %15v %15v %15v %v\n", name, i.InstanceId, i.InstanceType, i.Placement.AvailabilityZone, i.PrivateIPAddress, i.PublicIPAddress, tags)
			csv_data = csv_data + fmt.Sprintf(`"%s","%s","%s","%s","%s","%v","%s","%s","%s","%s","%s","%s","%s","%s"`,
				name, i.State.Name, i.InstanceId, i.InstanceType, i.InstanceLifecycle, i.CPUOptions.CoreCount, volumes, i.Placement.AvailabilityZone,
				i.VpcId, i.SubnetId, i.PrivateIPAddress, i.PublicIPAddress, sg, tags)
			csv_data += "\n"
		}
		if csv {
			fmt.Println(csv_data)
			return
		}

		fmt.Println("next")
		df, _ := imports.LoadFromCSV(context.Background(), strings.NewReader(csv_data))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(df.Table())

	},
}

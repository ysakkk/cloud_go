package cmd

import (

	//"os"
	//dataframe "github.com/rocketlaunchr/dataframe-go"
	// "github.com/rocketlaunchr/dataframe-go/exports"

	"context"
	"encoding/json"
	"strings"

	//"time"
	//"sort"

	//"reflect"
	//"io/ioutil"
	"bytes"
	"fmt"
	"os/exec"

	//"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
	"github.com/spf13/cobra"
	//toml "github.com/BurntSushi/toml"
	//toml "github.com/pelletier/go-toml/v2"
	//"gopkg.in/ini.v1"
)

var ebs = &cobra.Command{
	Use: "ebs",
	Run: func(cmd *cobra.Command, args []string) {

		if profile == "" {
			fmt.Println("Availabe Profiles:")
			listAWSProfiles()
			return
		}

		command := exec.Command("aws", "--profile", profile, "ec2", "describe-volumes")
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		err := command.Run()
		if err != nil {
			fmt.Printf("Stdout: %s\n", stdout.String())
			fmt.Printf("Stderr: %s\n", stderr.String())
			return
		}

		j := ec2_describe_volumes{}
		err = json.Unmarshal([]byte(stdout.String()), &j)
		if err != nil {
			fmt.Println(err)
			return
		}
		var csv_data string

		csv_data = fmt.Sprintf(`"%s","%s","%s","%s","%s","%s","%s","%s","%s"`,
			"VolumeId", "InstanceId", "State", "Device", "Attach", "AvilabilityZone", "Size", "Iops", "VolumeType")
		csv_data += "\n"

		for _, v := range j.Volumes {

			//fmt.Println(v.Attachments)
			if len(v.Attachments) == 0 {
				continue
			}
			csv_data = csv_data + fmt.Sprintf(`"%s","%s","%s","%s","%s","%s","%v","%v","%s"`,
				v.VolumeId, v.Attachments[0].InstanceId, v.State, v.Attachments[0].Device, v.Attachments[0].State, v.AvailabilityZone, v.Size, v.Iops, v.VolumeType)
			//csv_data = csv_data + fmt.Sprintf(`"%s","%s"`, v.VolumeId,v.Attachments[0].InstanceId )
			csv_data += "\n"
		}
		if csv {
			fmt.Println(csv_data)
			return
		}

		fmt.Println(csv_data)
		df, _ := imports.LoadFromCSV(context.Background(), strings.NewReader(csv_data))
		if err != nil {
			fmt.Println(err)
			return
		}
		//df.SetOption("display.max_colwidth", len(df))
		//df.SetOption("display.max_columns", len(df.Names()))
		//fmt.Print(reflect.TypeOf(df))
		fmt.Println(df.Table())

		//var buf bytes.Buffer
		//exports.ExportToCSV(context.Background(), &buf, df)
		//fmt.Println(buf.String())

	},
}


package cmd

import (
	"context"
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

var clb = &cobra.Command{
    Use:   "clb",
    Run: func(cmd *cobra.Command, args []string) {

		if profile == "" {
			fmt.Println("Availabe Profiles:")
			listTCProfiles()
			return
		}

	   j := clb_describe_loadbalancers{}
       result := execCommand("tccli", "--profile", profile, "clb", "DescribeLoadBalancers")
       err := json.Unmarshal([]byte(result), &j)
       if err != nil {
		   fmt.Println(err)
		   return
	   }
	   var csv_data string
       

	   csv_data = fmt.Sprintf(`"%s","%s","%s","%s"`,
	       "LoadBalancerId", "LoadBalancerName", "Zone", "Domain")
	   csv_data += "\n"
	   fmt.Println(csv_data)

	   for _, v := range j.LoadBalancerSet {
	       csv_data = csv_data + fmt.Sprintf(`"%s","%s","%s","%s"`,
               v.LoadBalancerId, v.LoadBalancerName, v.MasterZone.Zone, v.Domain)
	       csv_data += "\n"
	   }
	   if csv {
	       fmt.Println(csv_data)
	       return
       }

	   //fmt.Println("next")
	   df, _ := imports.LoadFromCSV(context.Background(), strings.NewReader(csv_data))
	   if err != nil {
	      fmt.Println(err)
	      return
	   }
	   fmt.Println(df)

    },
}


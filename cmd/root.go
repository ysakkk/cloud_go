package cmd


import (
	//"context"
	//dataframe "github.com/rocketlaunchr/dataframe-go"
    //"github.com/rocketlaunchr/dataframe-go/exports"
    //"github.com/rocketlaunchr/dataframe-go/imports"
	//"encoding/json"
	//"time"
	//"sort"
	"strings"
	//"reflect"
	//"io/ioutil"
    "bytes"
	"os/exec"
    "fmt"
    "os"
    "github.com/spf13/cobra"
    //toml "github.com/BurntSushi/toml"
     //toml "github.com/pelletier/go-toml/v2"
    "path/filepath"
	//"os/exec"
	"gopkg.in/ini.v1"
)

func getHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("cannot get user home dir path: %s", err.Error())
	}
    return home
}
var rootCmd = &cobra.Command{Use: "cloud_go"}

// go test -v cmd/*go
func testRootCmd() *cobra.Command {
    return &cobra.Command{}
}

func listAWSProfiles() {
	credFile := getHome() + "/.aws/credentials"
	cfg, _ := ini.Load(credFile)
	for _, p := range  cfg.SectionStrings() {
		if p == "DEFAULT" {
			continue
		}
		fmt.Println(p)
	}
}

func listTCProfiles() {
    fmt.Println("Tencent Profiles")
    pat := getHome() + `/.tccli/*.credential`
    match, err := filepath.Glob(pat)
	if err != nil {
		fmt.Errorf("cannot get tccli credentals: %s", err.Error())
	}
    for _, v := range match {
        vv := strings.Split(v, "/")
        vvv := strings.Split(vv[4], ".")
    	fmt.Println(vvv[0])
    }
}

func execCommand(c ...string) string {
       command := exec.Command(c[0], c[1:]...)
       var stdout bytes.Buffer
       var stderr bytes.Buffer
       command.Stdout = &stdout
       command.Stderr = &stderr
       err := command.Run()
       if err != nil {
           fmt.Printf("Stdout: %s\n", stdout.String())
           fmt.Printf("Stderr: %s\n", stderr.String())
	       panic("command error")
       } 
       return stdout.String()
}

var aws = &cobra.Command{
    Use:   "aws",
    Short: "aws",
    Run: func(cmd *cobra.Command, args []string)  {
    	if list {
			listAWSProfiles()
			return
    	}
		cmd.Help()
    },
}
var tc = &cobra.Command{
	Use: "tc", 
	Run: func(cmd *cobra.Command, args []string) {
    	if list {
			listTCProfiles()
			return
    	}
		cmd.Help()
    },
}

// 初期化
var profile string
var list bool
var csv bool
func init() {

    rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "", "AWS Profile")
    rootCmd.PersistentFlags().BoolVarP(&csv, "csv", "", false, "CSV Output")
    //rootCmd.MarkFlagRequired("profile")

    rootCmd.AddCommand(aws)
    aws.Flags().BoolVarP(&list, "list", "", false, "AWS Profile List")
    aws.AddCommand(ec2)
    aws.AddCommand(ebs)

    rootCmd.AddCommand(tc)
    tc.Flags().BoolVarP(&list, "list", "", false, "Tencent Profile List")
    tc.AddCommand(clb)
    //tc.PersistentFlags().StringVarP(&profile, "profile", "p", "", "Tencent Profile")
    //tc.MarkFlagRequired("profile")

    // 認証設定定数読み込み
    //loadConfig()
}

func Execute() {
    rootCmd.Execute()
}

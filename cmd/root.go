package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// 直接go run main.go 代表直接執行root
var rootCmd = &cobra.Command{
	Use:   "root",       //代表root命令該如何使用 git add
	Short: "short desc", //簡短描述 git
	Long:  "long desc",
	//命令處理函數
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		//打印flag
		fmt.Println(
			cmd.PersistentFlags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.PersistentFlags().Lookup("config").Value,
			cmd.PersistentFlags().Lookup("license").Value,
			cmd.Flags().Lookup("source").Value,
		)
		fmt.Println("-------------------------------")
		fmt.Println(
			//指定flag值以後viper的值產生變化但檔案不會被修改
			viper.GetString("author"),
			viper.GetString("license"),
		)
		fmt.Println("-------------------------------")
		fmt.Println("root cmd run end")
	},
	//允許子命令
	TraverseChildren: true,
}

func Execute() {
	rootCmd.Execute()
}

var (
	cfgFile     string
	userLicense string
)

// 使用flag --[flag名稱] -[縮寫]
func init() {
	cobra.OnInitialize(initConfig)
	//PersistentFlags 持久話標示->可以傳給子命令 value->默認值 usage->提示
	//因為此處為root命令，所以皆為全局命令
	//按名稱接受命令行參數
	rootCmd.PersistentFlags().Bool("viper", true, "")
	//指定flag縮寫
	rootCmd.PersistentFlags().StringP("author", "a", "Your name", "")
	//通過指針，儲存字段變量
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	//通過指針，儲存字段變量，並指定flag縮寫
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")

	//添加本地標誌
	rootCmd.Flags().StringP("source", "s", "", "")

	//綁定author
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	viper.SetDefault("author", "default author")
	viper.SetDefault("license", "default license")
}

// 配置綁定是為了改變viper對象的內容而非配置文件
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir() //默認為 /home
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	//檢查環境變量，將配致的鍵值加載到viper
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("using config file :", viper.ConfigFileUsed())
}

// go run main.go --viper=true -a nick -l apache --config myconfig.yaml -s local

package configuracion

import "github.com/spf13/viper"

func cargarConfiguraciones() {
	if !viper.IsSet("server") {
		viper.SetConfigFile(`config.json`)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
}

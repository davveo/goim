package helper

import (
	"fmt"

	"github.com/davveo/goim/cornet/lib/localIp"
	"github.com/spf13/viper"
)

var (
	rpcPort = viper.GetString("app.rpcPort")
)

func GenClientId() string {
	return fmt.Sprintf("%s:%s", localIp.OfferServerLocalIp(), rpcPort)
}

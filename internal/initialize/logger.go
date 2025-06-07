package initialize

import (
    "ecommerce/global"
    "github.com/sirupsen/logrus"
)

func InitLogger() {
    logger := logrus.New()
    logger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    global.Logger = logger
}

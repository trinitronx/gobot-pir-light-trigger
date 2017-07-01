package main

import (
   "fmt"

   "gobot.io/x/gobot"
   "gobot.io/x/gobot/drivers/gpio"
   "gobot.io/x/gobot/platforms/intel-iot/edison"
)

func main() {
   e := edison.NewAdaptor()

   // Was 5
   sensor := gpio.NewPIRMotionDriver(e, "7")
   led := gpio.NewLedDriver(e, "13")

   work := func() {
      sensor.On(gpio.MotionDetected, func(data interface{}) {
         fmt.Println(gpio.MotionDetected)
         led.On()
      })
      sensor.On(gpio.MotionStopped, func(data interface{}) {
         fmt.Println(gpio.MotionStopped)
         led.Off()
      })
   }

   robot := gobot.NewRobot("motionBot",
      []gobot.Connection{e},
      []gobot.Device{sensor, led},
      work,
   )

   robot.Start()
}

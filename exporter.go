package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"

    owm "github.com/3rein/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")
var zipCode, _ = strconv.Atoi(os.Getenv("OWM_ZIP_CODE"))

func main() {
    http.HandleFunc("/", func(respW http.ResponseWriter, r *http.Request) {
        futW, err := owm.NewForecast("5", "F", "en", apiKey)
        if err != nil {
            log.Fatalln(err)
        }

        w, err := owm.NewCurrent("F", "en", apiKey)
        if err != nil {
            log.Fatalln(err)
        }

        err = w.CurrentByZip(zipCode, "US")

        if err != nil {
            log.Fatalln(err)
        }

        err = futW.DailyByZip(zipCode, "US", 5)

        temp := w.Main.Temp
        pressure := w.Main.GrndLevel
        humidity := w.Main.Humidity

        rainLastHour := w.Rain.OneH
        cloudiness := w.Clouds.All
        windSpeed := w.Wind.Speed
        windDir := w.Wind.Deg

        fmt.Fprintf(respW, `
owm_temperature %f
owm_pressure_grnd %f
owm_humidity %d
owm_rain %f
owm_cloudiness %d
owm_wind_speed %f
owm_wind_dir %f
`, temp, pressure, humidity, rainLastHour, cloudiness, windSpeed, windDir)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr               = flag.String("listen-address", ":1971", "The address to listen on for HTTP requests.")
	location           = flag.String("location", "Gelsted", "The location to ask weather data for.")
	m_feels_like_c     = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_feels_like_c"})
	m_feels_like_f     = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_feels_like_f"})
	m_cloudcover       = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_cloudcover"})
	m_humidity         = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_humidity"})
	m_precipitation_mm = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_precipitation_mm"})
	m_pressure_mbar    = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_pressure_mbar"})
	m_temp_c           = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_temp_c"})
	m_temp_f           = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_temp_f"})
	m_uv_index         = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_uv_index"})
	m_visibility       = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_visibility"})
	m_winddir_degree   = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_winddir_degree"})
	m_windspeed_kmph   = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_windspeed_kmph"})
	m_windspeed_miles  = promauto.NewGauge(prometheus.GaugeOpts{Name: "wttr_windspeed_miles"})
)

func recordMetrics() {
	go func() {
		for {
			url := fmt.Sprintf("https://wttr.in/%v?format=j1", *location)
			resp, err := http.Get(url)
			if err != nil {
				// handle error
				fmt.Println("Error: %s", err.Error())
			}

			var jsonResp JSONResp
			err = json.NewDecoder(resp.Body).Decode(&jsonResp)
			if err != nil {
				fmt.Println("Decode Error: %s", err.Error())
			}

			currentCondition := jsonResp.CurrentCondition[0]

			m_feels_like_c.Set(currentCondition.FeelsLikeC)
			m_feels_like_f.Set(currentCondition.FeelsLikeF)
			m_cloudcover.Set(currentCondition.CloudCover)
			m_humidity.Set(currentCondition.Humidity)
			m_precipitation_mm.Set(currentCondition.PrecipitationMM)
			m_pressure_mbar.Set(currentCondition.PressureMBar)
			m_temp_c.Set(currentCondition.TempC)
			m_temp_f.Set(currentCondition.TempF)
			m_uv_index.Set(currentCondition.UVIndex)
			m_visibility.Set(currentCondition.Visibility)
			m_winddir_degree.Set(currentCondition.WinddirDegree)
			m_windspeed_kmph.Set(currentCondition.WindspeedKmpH)
			m_windspeed_miles.Set(currentCondition.WindspeedMiles)

			time.Sleep(60 * time.Second)
		}
	}()
}

func main() {
	flag.Parse()

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Exporter running at %s…", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

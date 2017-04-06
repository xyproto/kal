package calendar

// Equinox and solstice algorithms from Astronomical Algorithms by Jean Meeus
// Adapted from jgiesen.de/astro/astroJS/seasons/seasons.js

import (
	"math"
	"time"
)

const (
	// One degree expressed in radians
	degrees = math.Pi / 180.0
)

//// Vårjevndogn / Vernal equinox / March equinox
//func Equinox(year int) time.Time {
//	// Source: http://www.phpro.org/examples/Get-Vernal-Equinox.html
//	days_from_base := 79.3125 + float64(year - 1970) * 365.2425
//	seconds_from_base := days_from_base * 86400.0
//	return time.Unix(round(seconds_from_base), 0)
//}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func roundf(x float64) float64 {
	return math.Floor(0.5 + x)
}

func round(x float64) int64 {
	return int64(roundf(x))
}

func tableFormula(x float64) (result float64) {
	// TODO: Replace with a table and a loop
	result += 485 * math.Cos(degrees*(324.96+x*1934.136))
	result += 203 * math.Cos(degrees*(337.23+x*32964.467))
	result += 199 * math.Cos(degrees*(342.08+x*20.186))
	result += 182 * math.Cos(degrees*(27.85+x*445267.112))
	result += 156 * math.Cos(degrees*(73.14+x*45036.886))
	result += 136 * math.Cos(degrees*(171.52+x*22518.443))
	result += 77 * math.Cos(degrees*(222.54+x*65928.934))
	result += 74 * math.Cos(degrees*(296.72+x*3034.906))
	result += 70 * math.Cos(degrees*(243.58+x*9037.513))
	result += 58 * math.Cos(degrees*(119.81+x*33718.147))
	result += 52 * math.Cos(degrees*(297.17+x*150.678))
	result += 50 * math.Cos(degrees*(21.02+x*2281.226))
	result += 45 * math.Cos(degrees*(247.54+x*29929.562))
	result += 44 * math.Cos(degrees*(325.15+x*31555.956))
	result += 29 * math.Cos(degrees*(60.93+x*4443.417))
	result += 18 * math.Cos(degrees*(155.12+x*67555.328))
	result += 17 * math.Cos(degrees*(288.79+x*4562.452))
	result += 16 * math.Cos(degrees*(198.04+x*62894.029))
	result += 14 * math.Cos(degrees*(199.76+x*31436.921))
	result += 12 * math.Cos(degrees*(95.39+x*14577.848))
	result += 12 * math.Cos(degrees*(287.11+x*31931.756))
	result += 12 * math.Cos(degrees*(320.81+x*34777.259))
	result += 9 * math.Cos(degrees*(227.73+x*1222.114))
	result += 8 * math.Cos(degrees*(15.45+x*16859.074))
	return
}

// Calculate vårjevndøgn, sommersolverv, høstjevndøgn or vintersolverv
func calculateEquinoxOrSolstice(year int, fn func(float64) float64) time.Time {
	// TODO: Simplify with a symbolic calculator
	a := fn((float64(year) - 2000.0) / 1000.0)
	b := (a - 2451545.0) / 36525.0
	c := (35999.373*b - 2.47) * degrees
	d := a + (0.00001*tableFormula(b))/(1.0+0.0334*math.Cos(c)+0.0007*math.Cos(2*c)) - (66.0+float64(year-2000)*1.0)/86400.0
	e := roundf(d)
	f := math.Floor((e - 1867216.25) / 36524.25)
	g := e + f - math.Floor(f/4) + 1525.0
	h := math.Floor((g - 122.1) / 365.25)
	i := 365.0*h + math.Floor(h/4)
	k := math.Floor((g - i) / 30.6001)
	l := 24.0 * (d + 0.5 - e)
	day := int(roundf(g-i) - math.Floor(30.6001*k))
	month := k - 1 - 12*math.Floor(k/14)
	hour := int(math.Floor(l))
	minute := int(round((abs(l) - math.Floor(abs(l))) * 60.0))
	if minute == 60 {
		minute = 0
		hour++
	}
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
}

// Spring equinox for the northern hemisphere
func northwardEquinox(year int) time.Time {
	march := func(y float64) float64 {
		return 2451623.80984 + 365242.37404*y + 0.05169*y*y - 0.00411*y*y*y - 0.00057*y*y*y*y
	}
	return calculateEquinoxOrSolstice(year, march)
}

// Summer solstice for the northern hemisphere
func northernSolstice(year int) time.Time {
	june := func(y float64) float64 {
		return 2451716.56767 + 365241.62603*y + 0.00325*y*y + 0.00888*y*y*y - 0.00030*y*y*y*y
	}
	return calculateEquinoxOrSolstice(year, june)
}

// Autumn equinox for the northern hemisphere
func southwardEquinox(year int) time.Time {
	september := func(y float64) float64 {
		return 2451810.21715 + 365242.01767*y - 0.11575*y*y + 0.00337*y*y*y + 0.00078*y*y*y*y
	}
	return calculateEquinoxOrSolstice(year, september)
}

// Winter solstice for the northern hemisphere
func southernSolstice(year int) time.Time {
	december := func(y float64) float64 {
		return 2451900.05952 + 365242.74049*y - 0.06223*y*y - 0.00823*y*y*y + 0.00032*y*y*y*y
	}
	return calculateEquinoxOrSolstice(year, december)
}

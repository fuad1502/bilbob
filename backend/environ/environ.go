package environ

import "os"

var (
	Hostname       = os.Getenv("HOSTNAME")
	WebappUrl      = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME")
	LandingPageUrl = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME")
)

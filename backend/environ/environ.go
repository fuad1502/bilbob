package environ

import "os"

var (
	Hostname       = os.Getenv("HOSTNAME")
	WebappUrl      = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME") + ":" + os.Getenv("WEBAPP_PORT")
	LandingPageUrl = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME") + ":" + os.Getenv("LP_PORT")
)

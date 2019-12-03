package app

import (
	"ffxiv-profit-api/app/controllers"
	"ffxiv-profit-api/app/models"
	"net/http"
	"time"

	"github.com/revel/revel"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Builds the react scripts before anything runs
	// This only hot-reloads if we change this function.
	// Comment out if we're building
	/*BuildFrontEnd := exec.Command("npm", "run", "build")
	if err := BuildFrontEnd.Run(); err != nil {
		log.Fatal(err)
	}*/

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		ValidateOrigin,
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(controllers.InitDB)
	// revel.OnAppStart(FillCache)
}

// ValidateOrigin enables CORS policy, and handles pre-flight requests
var ValidateOrigin = func(c *revel.Controller, fc []revel.Filter) {
	originString := "https://ffxivprofit.com"
	// Allow only specific headers to be accessed.
	switch c.Request.GetHttpHeader("Origin") {
	case "http://localhost:3000":
		// Local Web App
		originString = "http://localhost:3000"
	case "https://ffxivprofit.com":
		// Public Web App
		originString = "https://ffxivprofit.com"
	case "http://localhost:3001":
		// Local Analytics
		originString = "http://localhost:3001"
	case "https://analytics.ffxivprofit.com":
		// Public Analytics
		originString = "https://analytics.ffxivprofit.com"
	case "https://example.com":
		// Kubernetes Web App
		originString = "https://example.com"
	case "https://analytics.example.com":
		// Kubernetes Analytics
		originString = "https://analytics.example.com"
	}
	// Log API Requests
	APILog := models.EndpointRequest{
		ClientIP:      c.ClientIP,
		Endpoint:      c.Request.URL.String(),
		RequestedTime: time.Now(),
	}
	controllers.LogEndpointRequest(APILog)

	// Pre-flight and Flight Request
	if c.Request.Method == "OPTIONS" {
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", originString)
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Response.Out.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Response.SetStatus(http.StatusNoContent)
	} else {
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", originString)
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Response.Out.Header().Add("Content-Type", "application/json; charset=UTF-8")
		c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
		c.Response.Out.Header().Add("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")

		fc[0](c, fc[1:]) // Execute the next filter stage.
	}
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

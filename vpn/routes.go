/*
 * App vpn
 *
 * API version: 1.0.0
 * Contact: support@peraMIC.io
 */

package vpn

import (
	"github.com/peramic/utils"
)

// Routes returns all of the api route for the Controller
var Routes = utils.Routes{

	utils.Route{
		Name:        "UploadOpenVPNConfig",
		Method:      "PUT",
		Pattern:     "/rest/default/config",
		HandlerFunc: UploadOpenVPNConfig,
	},

	utils.Route{
		Name:        "IsEnabled",
		Method:      "GET",
		Pattern:     "/rest/default/enable",
		HandlerFunc: IsEnabled,
	},

	utils.Route{
		Name:        "SetEnable",
		Method:      "PUT",
		Pattern:     "/rest/default/enable",
		HandlerFunc: SetEnable,
	},

	utils.Route{
		Name:        "DownloadOpenVPNlog",
		Method:      "GET",
		Pattern:     "/rest/default/log",
		HandlerFunc: DownloadOpenVPNlog,
	},
}

package daemon

import (
	"github.com/gin-gonic/gin"

	"github.com/replicatedhq/ship/pkg/api"
)

func (d *NavcycleRoutes) getMetadata(release *api.Release) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch release.Metadata.Type {
		default:
			fallthrough
		case "k8s":
			fallthrough
		case "helm":
			info := release.Metadata.ShipAppMetadata
			c.JSON(200, info)
			return
		case "runbook.replicated.app":
			fallthrough
		case "replicated.app":
			c.JSON(200, map[string]interface{}{
				"name": release.Metadata.ChannelName,
				"icon": release.Metadata.ChannelIcon,
			})
			return
		case "inline.replicated.app":
			info := map[string]interface{}{
				"name": release.Metadata.ChannelName,
				"icon": release.Metadata.ChannelIcon,
			}

			// the raw url tends to be ugly, but better than nothing
			if info["name"] == "" {
				info["name"] = release.Metadata.ShipAppMetadata.URL
			}

			c.JSON(200, info)
			return
		}
	}
}

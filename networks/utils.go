package networks

import "github.com/gin-gonic/gin"

// Get method routes
func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

// Post method routes
func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

// Put method routes
func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

// Delete method routes
func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

// Success Response
func (n *Network) okResponse (c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

// Failed Response
func (n *Network) failedResponse (c *gin.Context, result interface{}) {
	c.JSON(400, result)
}

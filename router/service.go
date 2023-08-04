package router

import (
	"hotelPMS/auth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {

	router := gin.Default()
	router.Use(CORS())

	cookieStore := cookie.NewStore([]byte(auth.SECRECT_KEY))
	router.Use(sessions.Sessions("go-pMs", cookieStore))

	//router.HTMLRender = loadTemplates("./web/templates")

	// router.LoadHTMLGlob("./web/templates/pages/**/*")

	// router.Static("/img", "./web/assets/dist/img")
	// router.Static("/css", "./web/assets/dist/css")
	// router.Static("/js", "./web/assets/dist/js")
	// router.Static("/plugins", "./web/assets/plugins")
	// router.Static("/fonts", "./web/assets/dist/fonts")

	apiPath := "api/v1"
	authService := auth.NewService()

	Handler(db, router, apiPath, authService)

	return router

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// func loadTemplates(templatesDir string) multitemplate.Renderer {
// 	r := multitemplate.NewRenderer()

// 	mainLayouts, err := filepath.Glob(templatesDir + "/layouts/main/*")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	mainIncludes, err := filepath.Glob(templatesDir + "/webapp/**/*")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// Generate our templates map from our layouts/ and includes/ directories

// 	for _, include := range mainIncludes {
// 		layoutCopy := make([]string, len(mainLayouts))
// 		copy(layoutCopy, mainLayouts)
// 		files := append(layoutCopy, include)
// 		r.AddFromFiles(filepath.Base(include), files...)
// 	}

// 	loginLayouts, err := filepath.Glob(templatesDir + "/layouts/login/*")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	sessionIncludes, err := filepath.Glob(templatesDir + "/session/*")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for _, sessionIncludes := range sessionIncludes {
// 		layoutCopy := make([]string, len(loginLayouts))
// 		copy(layoutCopy, loginLayouts)
// 		files := append(layoutCopy, sessionIncludes)
// 		r.AddFromFiles(filepath.Base(sessionIncludes), files...)
// 	}

// 	// menuLayouts, err := filepath.Glob(templatesDir + "/layouts/partials/sidebar.html")

// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// menuIncludes, err := filepath.Glob(templatesDir + "/webapp/**/*")

// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// // Generate our templates map from our layouts/ and includes/ directories

// 	// for _, include := range menuIncludes {
// 	// 	layoutCopy := make([]string, len(menuLayouts))
// 	// 	copy(layoutCopy, menuLayouts)
// 	// 	files := append(layoutCopy, include)
// 	// 	r.AddFromFiles(filepath.Base(include), files...)

// 	// 	fmt.Println(files)
// 	// }

// 	return r
// }

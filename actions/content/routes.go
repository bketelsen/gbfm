package content

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo"
)

func AddRoutes(app *buffalo.App) {
	app.Use(Auth)

	app.GET("/admin", adaptHandler(adminHandler))

	app.ANY("/admin/init", adaptHandler(initHandler))

	app.ANY("/admin/login", adaptHandler(loginHandler))
	app.GET("/admin/logout", adaptHandler(logoutHandler))

	app.ANY("/admin/recover", adaptHandler(forgotPasswordHandler))
	app.ANY("/admin/recover/key", adaptHandler(recoveryKeyHandler))

	app.ANY("/admin/addons", adaptHandler(addonsHandler))
	app.ANY("/admin/addon", adaptHandler(addonHandler))

	app.ANY("/admin/configure", adaptHandler(configHandler))
	app.ANY("/admin/configure/users", adaptHandler(configUsersHandler))
	app.ANY("/admin/configure/users/edit", adaptHandler(configUsersEditHandler))
	app.ANY("/admin/configure/users/delete", adaptHandler(configUsersDeleteHandler))

	app.GET("/admin/uploads", adaptHandler(uploadContentsHandler))
	app.GET("/admin/uploads/search", adaptHandler(uploadSearchHandler))

	app.GET("/admin/contents", adaptHandler(contentsHandler))
	app.GET("/admin/contents/search", adaptHandler(searchHandler))
	app.GET("/admin/contents/export", adaptHandler(exportHandler))

	app.ANY("/admin/edit", adaptHandler(editHandler))
	app.POST("/admin/edit/delete", adaptHandler(deleteHandler))
	app.POST("/admin/edit/approve", adaptHandler(approveContentHandler))
	app.ANY("/admin/edit/upload", adaptHandler(editUploadHandler))
	app.POST("/admin/edit/upload/delete", adaptHandler(deleteUploadHandler))

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Couldn't find current directory for file server.")
	}

	staticDir := filepath.Join(pwd, "pkg", "system")
	// TODO: I don't think these two are the same. Fix
	// http.Handle("/admin/static/", db.CacheControl(http.FileServer(restrict(http.Dir(staticDir)))))
	app.ServeFiles("/admin/static/", http.Dir(staticDir))

	// API path needs to be registered within server package so that it is handled
	// even if the API server is not running. Otherwise, images/files uploaded
	// through the editor will not load within the admin system.
	uploadsDir := filepath.Join(pwd, "uploads")
	// TODO: I don't think these two are the same. Fix
	// http.Handle("/api/uploads/", api.Record(api.CORS(db.CacheControl(http.StripPrefix("/api/uploads/", http.FileServer(restrict(http.Dir(uploadsDir))))))))
	app.ServeFiles("/api/uploads/", http.Dir(uploadsDir))

	// Database & uploads backup via HTTP route registered with Basic Auth middleware.
	app.GET("/admin/backup", BasicAuth(adaptHandler(backupHandler)))
}

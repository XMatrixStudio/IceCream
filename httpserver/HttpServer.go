package httpserver

import (
	"time"

	"github.com/XMatrixStudio/IceCream/httpserver/controllers"
	"github.com/XMatrixStudio/IceCream/httpserver/models"
	"github.com/XMatrixStudio/IceCream/httpserver/services"
	"github.com/XMatrixStudio/Violet.SDK.Go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// HTTPConfig 配置文件
type HTTPConfig struct {
	URL  string `yaml:"Url"`  // Url
	Host string `yaml:"Host"` // 服务器监听地址
	Port string `yaml:"Port"` // 服务器监听端口
	Dev  bool   `yaml:"Dev"`  // 是否开发环境
}

// Config 配置文件
type Config struct {
	Mongo      models.Mongo     `yaml:"Mongo"`  // mongoDB配置
	HTTPServer HTTPConfig       `yaml:"Server"` // iris配置
	Violet     violetSdk.Config `yaml:"Violet"` // Violet配置
}

// RunServer 监听APIs
func RunServer(c Config) {
	// 初始化数据库
	Model, err := models.NewModel(c.Mongo)
	if err != nil {
		panic(err)
	}
	// 初始化服务
	Service := services.NewService(Model)
	userService := Service.NewUserService()
	userService.InitViolet(c.Violet)
	// 启动服务器
	app := iris.New()
	if c.HTTPServer.Dev {
		app.Logger().SetLevel("debug")
	}

	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessionBug",
		Expires: 24 * time.Hour,
	})
	// "/users" based mvc application.
	users := mvc.New(app.Party("/users"))
	// Bind the "userService" to the UserController's Service (interface) field.
	users.Register(userService, sessManager.Start)
	users.Handle(new(controllers.UsersController))

	articleService := Service.NewArticleService()
	articles := mvc.New(app.Party("/articles"))
	articles.Register(articleService)
	articles.Handle(new(controllers.ArticlesController))

	app.Run(
		// Starts the web server
		iris.Addr(c.HTTPServer.Host+":"+c.HTTPServer.Port),
		// Disables the updater.
		iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}

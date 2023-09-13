// routes.go recebe e redireciona as requests da API
package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//security "github.com/vidacalura/BonziTech-TCC/internal/security"
	services "github.com/vidacalura/BonziTech-TCC/internal/services"
	utils "github.com/vidacalura/BonziTech-TCC/internal/utils"
)

func CriarRouter() *gin.Engine {
	services.DB = utils.ConectarBD()
	r := gin.Default()

	/* Middleware */
	//r.Use(security.ValidacaoRequest)
	r.Use(services.CriarLogDBMiddleware)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
        AllowHeaders: 	  []string{"*"},
        AllowCredentials: true,
	}))

	/* Rotas */	
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", services.ValidarLogin)
		//auth.POST("/usuario", services.ValidarPermissoesUsuario)
	}

	cat := r.Group("/api/categorias")
	{
		cat.GET("/", services.MostrarTodasCategorias)
		cat.GET("/:codCat", services.MostrarComponentesCategoria) 
		cat.POST("/", services.CriarCategoria)
		cat.PUT("/", services.AtualizarCategoria)
		cat.DELETE("/:codCat", services.DeletarCategoria)
	}

	comp := r.Group("/api/componentes")
	{
		comp.GET("/", services.MostrarTodosComponentes)
		comp.GET("/:codComp", services.MostrarComponente)
		comp.POST("/", services.AdicionarComponente)
		comp.PUT("/", services.AtualizarComponente)
		comp.DELETE("/:codComp", services.DeletarComponente)
	}

	// compEnt := r.Group("/api/componentes/entrada")

	// compSaida := r.Group("/api/componentes/saida")

	estq := r.Group("/api/estoque")
	{
		estq.GET("/", services.MostrarEstoque)
		estq.POST("/", services.AdicionarComponenteEstoque)
		estq.PUT("/", services.AtualizarEstoque)
		estq.DELETE("/:codComp", services.DeletarComponenteEstoque)
	}

	sessao := r.Group("/api/sessao")
	{
		sessao.GET("/:codSessao", services.GetSessao)
		sessao.POST("/", services.CriarSessao)
		sessao.DELETE("/", services.FecharSessao)
	}

	subcat := r.Group("/api/subcategorias")
	{
		subcat.GET("/categoria/:codCat", services.MostrarSubcategoriasDeCategoria)
		subcat.GET("/subcategoria/:codSubcat", services.MostrarComponentesSubcategoria) // 
		subcat.POST("/", services.CriarSubcategoria)
		subcat.PUT("/", services.AtualizarSubcategoria)
		subcat.DELETE("/:codSubcat", services.DeletarSubcategoria)
	}

	usu := r.Group("/api/usuarios") 
	{
		usu.GET("/", services.MostrarTodosUsuarios)
		usu.GET("/:codUsu", services.MostrarUsuario)
		usu.POST("/login", services.ValidarDadosLogin)
		usu.POST("/", services.AdicionarUsuario)
		usu.PUT("/", services.AtualizarUsuario)
		usu.DELETE("/:codUsu", services.DeletarUsuario)
	}

	r.GET("/api/ping", pong)
		
	return r
}

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{ "message": "pong!" })
}
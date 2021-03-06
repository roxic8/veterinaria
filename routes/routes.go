package routes

import (
	"github.com/Julianrt/veterinaria/handlers"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

//StartHandleRoutes f
func StartHandleRoutes(app *fiber.App) {
	app.Settings.Templates = html.New("./templates", ".html")

	app.Get("/", handlers.AgendarCita)
	app.Static("/public/", "./assets")
	api(app)
	webApp(app)
}

func api(app *fiber.App) {
	empleadoRoutes(app)
	serviciosRoutes(app)

	app.Get("/fechas_disponibles/", handlers.FechasDisponibles)
	app.Get("/fechas_ocupadas/", handlers.FechasOcupadas)
	app.Post("/agendar_consulta/", handlers.AgendarConsulta)
}

func serviciosRoutes(app *fiber.App) {
	app.Get("/servicios/", handlers.GetServicios)
}

func empleadoRoutes(app *fiber.App) {
	app.Post("/empleados/", handlers.CreateEmpleado)
	app.Get("/empleados/", handlers.GetEmpleados)
	app.Get("/empleados/:id/", handlers.GetEmpleado)
	app.Put("/empleados/:id/", handlers.UpdateEmpleado)
	app.Delete("/empleados/:id/", handlers.DeleteEmpleado)
}

func webApp(app *fiber.App) {
	app.Get("/historial/", handlers.Historial)
	app.Get("/agendar/", handlers.AgendarCita)
	app.Post("/agendar/", handlers.AgendarCita)
	app.Get("/agenda/", handlers.Agenda)
	app.Get("/consulta/", handlers.Consulta)
	app.Post("/consulta/", handlers.Consulta)
	app.Get("/registrar/", handlers.Registrar)
}

// GENERATED CODE - DO NOT EDIT
// This file is the run file for Revel.
// It registers all the controllers and provides details for the Revel server engine to
// properly inject parameters directly into the action endpoints.
package run

import (
	_ "Web/app"
	controllers "Web/app/controllers"
	tests "Web/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel"
	_ "github.com/revel/revel"
	_ "github.com/revel/revel/cache"
	"github.com/revel/revel/testing"
	"reflect"
)

var (
	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

// Register and run the application
func Run(port int) {
	Register()
	revel.Run(port)
}

// Register all the controllers
func Register() {
	revel.AppLog.Info("Running revel server")

	revel.RegisterController((*controllers.Authors)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Authors",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					69: []string{
						"auth",
					},
				},
			},
			&revel.MethodType{
				Name: "Author",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{
					90: []string{
						"id",
						"firstName",
						"lastName",
						"desc",
					},
				},
			},
			&revel.MethodType{
				Name: "Remove",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Change",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
					&revel.MethodArg{Name: "FirstName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "LastName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "Description", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					151: []string{
						"id",
						"firstName",
						"lastName",
						"desc",
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "FirstName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "LastName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "Description", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					176: []string{},
				},
			},
		})

	revel.RegisterController((*controllers.Books)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Books",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					108: []string{
						"bks",
					},
				},
			},
			&revel.MethodType{
				Name: "Book",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{
					185: []string{
						"name",
						"genre",
						"author",
						"publisher",
						"added_User",
						"added_Time",
						"description",
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "publishers", Type: reflect.TypeOf((*[]controllers.BookPublishers)(nil))},
					&revel.MethodArg{Name: "users", Type: reflect.TypeOf((*[]controllers.BookUsers)(nil))},
					&revel.MethodArg{Name: "authors", Type: reflect.TypeOf((*[]controllers.BookAuthors)(nil))},
				},
				RenderArgNames: map[int][]string{
					258: []string{},
				},
			},
		})

	revel.RegisterController((*controllers.Error)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Error",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					12: []string{},
				},
			},
		})

	revel.RegisterController((*controllers.Login)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					13: []string{},
				},
			},
		})

	revel.RegisterController((*controllers.Publishers)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Publishers",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					65: []string{
						"pubs",
					},
				},
			},
			&revel.MethodType{
				Name: "Publisher",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{
					85: []string{
						"id",
						"name",
						"desc",
					},
				},
			},
			&revel.MethodType{
				Name: "Remove",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Change",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
					&revel.MethodArg{Name: "Name", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "Description", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					144: []string{
						"id",
						"name",
						"desc",
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "Name", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "Description", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					167: []string{},
				},
			},
		})

	revel.RegisterController((*controllers.Registration)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Registration",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					13: []string{},
				},
			},
		})

	revel.RegisterController((*controllers.Users)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Users",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					99: []string{
						"usrs",
					},
				},
			},
			&revel.MethodType{
				Name: "User",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{
					174: []string{
						"Name",
						"Login",
						"Level",
					},
				},
			},
		})

	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					76: []string{
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					125: []string{},
				},
			},
			&revel.MethodType{
				Name:           "List",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.DefaultValidationKeys = map[string]map[int]string{}
	testing.TestSuites = []interface{}{
		(*tests.AppTest)(nil),
	}
}

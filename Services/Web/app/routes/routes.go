// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"

type tApp struct{}

var App tApp

func (_ tApp) Index() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("App.Index", args).URL
}

type tAuthors struct{}

var Authors tAuthors

func (_ tAuthors) Authors() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Authors.Authors", args).URL
}

type tBooks struct{}

var Books tBooks

func (_ tBooks) Books() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Books.Books", args).URL
}

type tLogin struct{}

var Login tLogin

func (_ tLogin) Login() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Login.Login", args).URL
}

type tPublishers struct{}

var Publishers tPublishers

func (_ tPublishers) Publishers() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Publishers.Publishers", args).URL
}

type tRegistration struct{}

var Registration tRegistration

func (_ tRegistration) Registration() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Registration.Registration", args).URL
}

type tUsers struct{}

var Users tUsers

func (_ tUsers) Users() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Users.Users", args).URL
}

type tStatic struct{}

var Static tStatic

func (_ tStatic) Serve(
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
	moduleName string,
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
	moduleName string,
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}

type tTestRunner struct{}

var TestRunner tTestRunner

func (_ tTestRunner) Index() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
	suite string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
	suite string,
	test string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}

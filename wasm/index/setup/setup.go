package setup

import (
	"noegotribes/ui/pages"
	"noegotribes/ui/state"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

// OVERALL APPLICATION WIRING IN vuguSetup
func VuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	tl := state.LoadTacoListAPI()
	ca := state.LoadLoginAPI()
	// CREATE A NEW ROUTER INSTANCE
	router := vgrouter.New(eventEnv)

	// MAKE OUR WIRE FUNCTION POPULATE ANYTHING THAT WANTS A "NAVIGATOR".
	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
		if s, ok := b.(state.TacoListAPISetter); ok {
			s.TacoListAPISet(tl)
		}
		if s, ok := b.(state.LoginAPISetter); ok {
			s.LoginAPISet(ca)
		}
	})

	// CREATE THE ROOT COMPONENT
	root := &pages.Root{}
	buildEnv.WireComponent(root) // WIRE IT
	router.MustAddRouteExact("/",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
			root.Body = &pages.Index{}
		}))
	router.MustAddRouteExact("/login",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
			root.Body = &pages.Login{}
		}))
	router.SetNotFound(vgrouter.RouteHandlerFunc(
		func(rm *vgrouter.RouteMatch) {
			root.Body = &pages.PageNotFound{} // A PAGE FOR THE NOT-FOUND CASE
		}))

	// TELL THE ROUTER TO LISTEN FOR THE BROWSER CHANGING URLS
	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	// GRAB THE CURRENT BROWSER URL AND PROCESS IT AS A ROUTE
	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return root
}

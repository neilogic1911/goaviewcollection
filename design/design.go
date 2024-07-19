package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("goaviewcollection", func() {
	Title("View Collection")
	Server("goaviewcollection", func() {
		Services("viewcollection")
	})
})

var _ = Service("viewcollection", func() {
	Method("list_element", func() {
		Result(CollectionOf(Element), func() {
			View("tiny")
		})

		HTTP(func() {
			GET("/viewcollection")
			Response(StatusOK)
		})
	})

	Method("list_element_in_attribute", func() {
		Result(func() {
			Attribute("data", CollectionOf(Element), func() {
				View("tiny")
			})
		})

		HTTP(func() {
			GET("/viewcollection-attr")
			Response(StatusOK)
		})
	})
})

var Element = ResultType("application/vnd.goaviewcollection.element", "ElementResult", func() {
	Attributes(func() {
		Attribute("id", Int, "ID of the element")
		Attribute("name", String, "Name of the element")
		Attribute("description", String, "Description of the element")
		Attribute("created_at", String, "Creation time of the element")
		Required("id", "name", "description", "created_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("created_at")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})
})

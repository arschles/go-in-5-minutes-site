package screencasts

package actions

import (
	"fmt"

	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/views"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Screencast)
// DB Table: Plural (screencasts)
// Resource: Plural (Screencasts)
// Path: Plural (/screencasts)
// View Template Folder: Plural (/templates/screencasts/)

// BaseResource has the routes in it for read-only
type ReadOnlyResource struct {
}

// List gets all Screencasts. This function is mapped to the path
// GET /screencasts
func (v ScreencastsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	screencasts := &models.Screencasts{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Screencasts from the DB
	if err := q.All(screencasts); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	view, err := views.Screencasts(parsedManifest, screencasts)
	if err != nil {
		return err
	}

	return c.Render(200, render.EltToRenderer(view))
}

// Show gets the data for one Screencast. This function is mapped to
// the path GET /screencasts/{screencast_id}
func (v ScreencastsResource) Show(c buffalo.Context) error {
	// TODO: implement
	return c.Error(404, fmt.Errorf("Not found"))
	// // Get the DB connection from the context
	// tx, ok := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return fmt.Errorf("no transaction found")
	// }

	// // Allocate an empty Screencast
	// screencast := &models.Screencast{}

	// // To find the Screencast the parameter screencast_id is used.
	// if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
	// 	return c.Error(404, err)
	// }

	// return c.Render(200, r.Auto(c, screencast))
}

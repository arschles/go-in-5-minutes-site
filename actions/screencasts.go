package actions

import (
	"fmt"

	"github.com/arschles/go-in-5-minutes-site/models"
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

// ScreencastsResource is the resource for the Screencast model
type ScreencastsResource struct {
	buffalo.Resource
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

	return c.Render(200, r.Auto(c, screencasts))
}

// Show gets the data for one Screencast. This function is mapped to
// the path GET /screencasts/{screencast_id}
func (v ScreencastsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// To find the Screencast the parameter screencast_id is used.
	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, screencast))
}

// New renders the form for creating a new Screencast.
// This function is mapped to the path GET /screencasts/new
func (v ScreencastsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Screencast{}))
}

// Create adds a Screencast to the DB. This function is mapped to the
// path POST /screencasts
func (v ScreencastsResource) Create(c buffalo.Context) error {
	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// Bind screencast to the html form elements
	if err := c.Bind(screencast); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(screencast)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, screencast))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "screencast.created.success"))
	// and redirect to the screencasts index page
	return c.Render(201, r.Auto(c, screencast))
}

// Edit renders a edit form for a Screencast. This function is
// mapped to the path GET /screencasts/{screencast_id}/edit
func (v ScreencastsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, screencast))
}

// Update changes a Screencast in the DB. This function is mapped to
// the path PUT /screencasts/{screencast_id}
func (v ScreencastsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Screencast to the html form elements
	if err := c.Bind(screencast); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(screencast)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, screencast))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "screencast.updated.success"))
	// and redirect to the screencasts index page
	return c.Render(200, r.Auto(c, screencast))
}

// Destroy deletes a Screencast from the DB. This function is mapped
// to the path DELETE /screencasts/{screencast_id}
func (v ScreencastsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// To find the Screencast the parameter screencast_id is used.
	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(screencast); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "screencast.destroyed.success"))
	// Redirect to the screencasts index page
	return c.Render(200, r.Auto(c, screencast))
}

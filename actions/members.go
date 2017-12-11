package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/vuerecipe/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Member)
// DB Table: Plural (members)
// Resource: Plural (Members)
// Path: Plural (/members)
// View Template Folder: Plural (/templates/members/)

// MembersResource is the resource for the Member model
type MembersResource struct {
	buffalo.Resource
}

// List gets all Members. This function is mapped to the path
// GET /members
func (v MembersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transacation found"))
	}

	members := &models.Members{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Members from the DB
	if err := q.Where("band_id = ?", c.Param("band_id")).All(members); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the headers so clients know how to paginate.
	c.Response().Header().Set("X-Pagination", q.Paginator.String())

	return c.Render(200, r.JSON(members))
}

// Show gets the data for one Member. This function is mapped to
// the path GET /members/{member_id}
func (v MembersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Member
	member := &models.Member{}

	// To find the Member the parameter member_id is used.
	if err := tx.Find(member, c.Param("member_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(member))
}

// New default implementation. Returns a 404
func (v MembersResource) New(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// Create adds a Member to the DB. This function is mapped to the
// path POST /members
func (v MembersResource) Create(c buffalo.Context) error {
	// Allocate an empty Member
	member := &models.Member{}

	// Bind member to the html form elements
	if err := c.Bind(member); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(member)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(201, r.JSON(member))
}

// Edit default implementation. Returns a 404
func (v MembersResource) Edit(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// Update changes a Member in the DB. This function is mapped to
// the path PUT /members/{member_id}
func (v MembersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Member
	member := &models.Member{}

	if err := tx.Where("band_id = ?", c.Param("band_id")).Find(member, c.Param("member_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Member to the html form elements
	if err := c.Bind(member); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(member)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(200, r.JSON(member))
}

// Destroy deletes a Member from the DB. This function is mapped
// to the path DELETE /members/{member_id}
func (v MembersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Member
	member := &models.Member{}

	// To find the Member the parameter member_id is used.
	if err := tx.Where("band_id = ?", "band_id").Find(member, c.Param("member_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(member); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(member))
}

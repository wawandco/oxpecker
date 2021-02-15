package resource

var actionTmpl string = `package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v5"

	"{{ .ModelPkg }}"
)

// {{ .Name.Resource }}Resource is the resource for the {{.Name.Proper}} model
type {{ .Name.Resource }}Resource struct{
	buffalo.Resource
}

// List gets all {{.Name.Group }}. This function is mapped to the path
// GET /{{ .Name.URL }}
func (v {{ .Name.Resource }}Resource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	{{.Name.VarCasePlural}} := &models.{{ .Name.Group }}{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all {{ .Name.Group }} from the DB
	if err := q.All({{ .Name.VarCasePlural }}); err != nil {
		return err
	}

    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)
    c.Set("{{ .Name.VarCasePlural }}", {{ .Name.VarCasePlural }})

    return c.Render(http.StatusOK, r.HTML("/{{ .Name.File.Pluralize }}/index.plush.html"))
}

// Show gets the data for one {{.Name.Proper}}. This function is mapped to
// the path GET /{{ .Name.URL }}/{{"{"}}{{ .Name.ParamID }}}
func (v {{ .Name.Resource }}Resource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty {{ .Name.Proper }}
	{{ .Name.VarCaseSingle }} := &models.{{ .Name.Proper }}{}

	// To find the {{ .Name.Proper }} the parameter {{ .Name.ParamID }} is used.
	if err := tx.Find({{ .Name.VarCaseSingle }}, c.Param("{{ .Name.ParamID }}")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

    c.Set("{{ .Name.VarCaseSingle }}", {{ .Name.VarCaseSingle }})

    return c.Render(http.StatusOK, r.HTML("/{{ .Name.File.Pluralize }}/show.plush.html"))
}

// New renders the form for creating a new {{ .Name.Proper }}.
// This function is mapped to the path GET /{{ .Name.URL }}/new
func (v {{ .Name.Resource }}Resource) New(c buffalo.Context) error {
	c.Set("{{.Name.VarCaseSingle}}", &models.{{ .Name.Proper }}{})

	return c.Render(http.StatusOK, r.HTML("/{{ .Name.File.Pluralize }}/new.plush.html"))
}

// Create adds a {{ .Name.Proper }} to the DB. This function is mapped to the
// path POST /{{ .Name.URL }}
func (v {{ .Name.Resource }}Resource) Create(c buffalo.Context) error {
	// Allocate an empty {{.Name.Proper}}
	{{ .Name.VarCaseSingle }} := &models.{{ .Name.Proper }}{}

	// Bind {{ .Name.VarCaseSingle }} to the html form elements
	if err := c.Bind({{ .Name.VarCaseSingle }}); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate({{ .Name.VarCaseSingle }})
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("{{ .Name.VarCaseSingle }}", {{ .Name.VarCaseSingle }})

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/{{ .Name.File.Pluralize }}/new.plush.html"))
	}

    // If there are no errors set a success message
    c.Flash().Add("success","{{ .Name.VarCaseSingle }}.created.success")

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "{{ .Name.VarCaseSingle }}Path()", render.Data{"{{ .Name.ParamID }}": {{ .Name.VarCaseSingle }}.ID})
}

// Edit renders a edit form for a {{ .Name.Proper }}. This function is
// mapped to the path GET /{{ .Name.URL }}/{{"{"}}{{ .Name.ParamID }}}/edit
func (v {{ .Name.Resource }}Resource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty {{ .Name.Proper }}
	{{ .Name.VarCaseSingle }} := &models.{{ .Name.Proper }}{}

	if err := tx.Find({{ .Name.VarCaseSingle }}, c.Param("{{ .Name.ParamID }}")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("{{.Name.VarCaseSingle}}", {{.Name.VarCaseSingle}})

	return c.Render(http.StatusOK, r.HTML("/{{.Name.File.Pluralize}}/edit.plush.html"))
}

// Update changes a {{ .Name.Proper }} in the DB. This function is mapped to
// the path PUT /{{ .Name.URL}}/{{"{"}}{{ .Name.ParamID }}}
func (v {{ .Name.Resource }}Resource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty {{ .Name.Proper }}
	{{ .Name.VarCaseSingle }} := &models.{{ .Name.Proper }}{}

	if err := tx.Find({{ .Name.VarCaseSingle }}, c.Param("{{ .Name.ParamID }}")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind {{ .Name.Proper }} to the html form elements
	if err := c.Bind({{ .Name.VarCaseSingle }}); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate({{ .Name.VarCaseSingle }})
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("{{ .Name.VarCaseSingle }}", {{ .Name.VarCaseSingle }})

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/{{ .Name.File.Pluralize }}/edit.plush.html"))
	}

    // If there are no errors set a success message
    c.Flash().Add("success", "{{ .Name.VarCaseSingle }}.updated.success")

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "{{ .Name.VarCaseSingle }}Path()", render.Data{"{{ .Name.ParamID }}": {{ .Name.VarCaseSingle }}.ID})
}

// Destroy deletes a {{ .Name.Proper }} from the DB. This function is mapped
// to the path DELETE /{{ .Name.URL }}/{{"{"}}{{ .Name.ParamID }}}
func (v {{ .Name.Resource }}Resource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty {{ .Name.Proper }}
	{{ .Name.VarCaseSingle }} := &models.{{ .Name.Proper }}{}

	// To find the {{ .Name.Proper }} the parameter {{ .Name.ParamID }} is used.
	if err := tx.Find({{ .Name.VarCaseSingle }}, c.Param("{{ .Name.ParamID }}")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy({{ .Name.VarCaseSingle }}); err != nil {
		return err
	}

    // If there are no errors set a flash message
    c.Flash().Add("success", "{{ .Name.VarCaseSingle }}.destroyed.success")

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "{{ .Name.VarCasePlural }}Path()")
}
`

var actionTestTmpl string = `package actions

{{ range $a := .Actions }}
func (as *ActionSuite) Test_{{$.Name.Resource}}Resource_{{ $a.Pascalize }}() {
  	as.Fail("Not Implemented!")
}
{{ end }}
`

// templateIndexTmpl is a variable for the [template_name]/index.plush.html
var templateIndexTmpl string = `<div class="py-4 mb-2">
	<h3 class="d-inline-block">{{ .Name.Group }}</h3>
	<div class="float-right">
		<%= linkTo(new{{ .Name.Resource }}Path(), {class: "btn btn-primary"}) { %>
		Create New {{ .Name.Proper }}
		<% } %>
	</div>
</div>

<table class="table table-hover table-bordered">
	<thead class="thead-light">
		{{ range $p := .Model.Attrs -}}
		{{ if ne $p.CommonType "text" -}}
		<th>{{ $p.Name.Pascalize }}</th>
		{{ end }}
		{{- end -}}
		<th>&nbsp;</th>
	</thead>
	<tbody>
		<%= for ({{ .Name.VarCaseSingle }}) in {{ .Name.VarCasePlural }} { %>
		<tr>
			{{ range $mp := .Model.Attrs -}}
			{{- if ne $mp.CommonType "text" -}}
			<td class="align-middle"><%= {{ $.Name.VarCaseSingle }}.{{ $mp.Name.Pascalize }} %></td>
			{{ end }}
			{{- end -}}
			<td>
				<div class="float-right">
					<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-info", body: "View"}) %>
					<%= linkTo(edit{{ .Name.Proper }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-warning", body: "Edit"}) %>
					<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>
				</div>
			</td>
		</tr>
		<% } %>
	</tbody>
</table>

<div class="text-center">
	<%= paginator(pagination) %>
</div>`

// templateNewTmpl is a variable for the [template_name]/new.plush.html
var templateNewTmpl string = `<div class="py-4 mb-2">
  	<h3 class="d-inline-block">New {{ .Name.Proper }}</h3>
</div>

<%= formFor({{.Name.VarCaseSingle}}, {action: {{ .Name.VarCasePlural }}Path(), method: "POST"}) { %>
	<%= partial("{{ .Name.Folder.Pluralize }}/form.html") %>
	<%= linkTo({{ .Name.VarCasePlural }}Path(), {class: "btn btn-warning", "data-confirm": "Are you sure?", body: "Cancel"}) %>
<% } %>`

// templateEditTmpl is a variable for the [template_name]/edit.plush.html
var templateEditTmpl string = `<div class="py-4 mb-2">
<h3 class="d-inline-block">Edit {{.Name.Proper}}</h3>
</div>

<%= formFor({{.Name.VarCaseSingle}}, {action: {{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), method: "PUT"}) { %>
	<%= partial("{{ .Name.Folder.Pluralize }}/form.html") %>
<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-warning", "data-confirm": "Are you sure?", body: "Cancel"}) %>
<% } %>`

// templateShowTmpl is a variable for the [template_name]/show.plush.html
var templateShowTmpl string = `<div class="py-4 mb-2">
	<h3 class="d-inline-block">{{ .Name.Proper }} Details</h3>
	<div class="float-right">
		<%= linkTo({{ .Name.VarCasePlural }}Path(), {class: "btn btn-info"}) { %>
		Back to all {{ .Name.Group }}
		<% } %>
		<%= linkTo(edit{{ .Name.Proper }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-warning", body: "Edit"}) %>
		<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>
	</div>
</div>

<ul class="list-group mb-2 ">
{{- range $p := .Model.Attrs }}
	<li class="list-group-item pb-1">
		<label class="small d-block">{{ $p.Name.Pascalize }}</label>
		<p class="d-inline-block"><%= {{$.Name.VarCaseSingle}}.{{$p.Name.Pascalize}} %></p>
	</li>
{{- end }}
</ul>`

// templateFormTmpl is a variable for the [template_name]/form.plush.html
var templateFormTmpl string = `{{ range $p := .Model.Attrs -}}
{{ if eq $p.CommonType "text" -}}
<%= f.TextAreaTag("{{$p.Name.Pascalize}}", {rows: 10}) %>
{{ else -}}
{{ if eq $p.CommonType "bool" -}}
<%= f.CheckboxTag("{{$p.Name.Pascalize}}", {unchecked: false}) %>
{{ else -}}
<%= f.InputTag("{{$p.Name.Pascalize}}") %>
{{ end -}}
{{ end -}}
{{ end -}}

<button class="btn btn-success" role="submit">Save</button>`

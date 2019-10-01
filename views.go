// In an MVC appliction, the views are the HTML that is
// sent to the user's browser. (In a native app, we'd
// say instead that the views are the structure of what
// is displayed on the screen.) Views are often also called
// "templates".
//
package main

import (
	"html/template"
)

var layoutMarkup = `
	{{define "layout"}}
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body {
					max-width: 800px;
					margin: 2rem auto;
				}

				header {
					background-color: #8ac5c3;
				}

				footer {
					background-color: #edece8;
				}
			</style>
		</head>
		<body>
			<header>This is the header</header>
			{{block "content" .}}
				This is the default of the block called "content"
				in our templates/views.
			{{end}}
			<footer>This is the footer</footer>
		</body>
	{{end}}
`

var attendeesMarkup = `
{{define "content"}}
	<h2>People attending the party:</h2>
	<ul>
	{{range $person := .Attendees}}
		<li>{{$person}}</li>
	{{end}}
	</ul>
	<form method="get">
		<label for="q">Search</label>
		<input type="text" name="q" value="">
	</form>
{{end}}
`

var indexMarkup = `
{{define "content"}}
	<p>
		This is the body of the index. See <a href="/attendees">attendees</a>.
	</p>
	<p>
	   PS. What do I love?
	   <ul>
		   <li><a href="/love?things=kittens">kittens</a></li>
		   <li><a href="/love?things=chocolate&things=peanut+butter">chocolate and peanut butter</a></li>
		   <li><a href="/love">nothing...</a></li>
	   </ul>
    </p>
{{end}}
`

var loveMarkup = `
{{define "content"}}
	{{.}}
{{end}}
`

// The below lines are a little confusing; please bear with me.
//
// Compile the layout template. The "Must" function tells our
// program to crash if the template does not compile.
var layoutTemplate = template.Must(template.New("layout").Parse(layoutMarkup))

// Now, layer the attendees template on top of the layout
var attendeesTemplate = template.Must(template.Must(layoutTemplate.Clone()).Parse(attendeesMarkup))

// Layer the index template on top of the layout
var indexTemplate = template.Must(template.Must(layoutTemplate.Clone()).Parse(indexMarkup))

// Layer the love template on top of the layout
var loveTemplate = template.Must(template.Must(layoutTemplate.Clone()).Parse(loveMarkup))

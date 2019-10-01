// In an MVC application, "models" hold our data.
// In this case, our data are names of people, which we're
// storing in an array of fixed size. In a "real"
// app, your data would likely be stored in a 
// database (like PostgreSQL or MySQL) and your
// "models" would contain code for talking to that
// database.
//
package main

var people = []string{
    "Taly Reich",
    "Kyle Jensen",
    "Anjani Jain",
    "Kerwin Charles",
    "Sharon Oster",
    "Sherri Scully",
}

type party struct {
	Attendees []string
}

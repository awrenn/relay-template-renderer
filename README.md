# Relay Go Templater

This is a Relay for applying an Object to a string template, which uses Go's template syntax + some extra helper functions. There are three parameters:

`template` - a string that represents the template  
`parameters` - a JSON object that gets passed in to the template as `.`.  
`output` - the variable to output the rendered template as.

## Helper functions

There are a number of helper functions added outside regular go templates:

Self-explanatory binomial functions:  
`add(i, j int) int`
`sub(i, j int) int`
`mul(i, j int) int`
`div(i, j int) int` 

Convert a date to a prettier format than default time:    
`date(t time.Time) string`

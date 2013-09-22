Go-scrape
=========

A simple HTML scraper.

###How to use
```go run scrape.go (config) (url) (output file)```
The configs are written in toml. They're very simple and straight forward. Tag is the HTML tag, such as <a> (but you'd only put in a) and AttributeName is the atribute you want out of the tag.

After running the program, you should end up with the output file you gave it in the arguement. Then you'd run something like wget -i outputfile to download all of them

###Todo
* Make a script to go through each url (the output doesn't always have to be a url as well) in the output file and complete the task again (we need to go deeper!) 
* Think of other things to make this program better.

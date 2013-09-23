Go-scrape
=========

A simple HTML scraper.

###How to use
```./scrape (config) (url) ``` Example: ```./scrape configs/ahref.toml http://golang.org/doc/  ```


Say you wanted to download all the URLs you just scraped, you'd cat the output into a file, lets call that file ouputfile, then you'd run something like ```wget -i outputfile``` to download all of them. 

#####How the config files work
The configs are written in toml. They're very simple and straight forward. It comes with some examples, here is an commented version of the imgsrc example included.
```
Tag = "img" #This is the name of the html tag, <img>, so you'll just put in img
AttributeName = "src" #This is the name of the atribute you want to scrape out of the tag you specifed above
IsURL = true # This is done so that if you're scraping a URL it'll attempt to fix them up so you can download them
```
With ``IsURL``:  /images/image.gif gets replaced to http://example.com/images/image.gif 
But if you're dealing with say the width of images, you don't want something like http://exmaple.com/64 so you'd set it to false.


###Todo
* Make some helper scripts to do things like remove duplicate entries and go through each url that you've put into your output file and go through each one again, to go deeper.


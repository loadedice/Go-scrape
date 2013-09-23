Go-scrape
=========

A simple HTML scraper.

###How to use
```./scrape (config) (url) (output file)``` Example: ```./scrape configs/ahref.toml http://golang.org/doc/ golinks ```



After running the program, you should end up with the output file you gave it in the arguement. Then you'd run something like ```wget -i outputfile``` to download all of them. 

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
* Make a script to go through each url in the output file and complete the task again (we need to go deeper!) 
* Come up with more "configs" in examples.
* Remove duplicate entries if it's a URL.
* Room for optimisation.
* Add in warnings when overwriting existing file.
* Just make it print out the lines rather than put it into a file. You can just cat it in to a file if you want to. This allows you to use it with other programs easily and so you can also do some sed trickery with pipes with ease.

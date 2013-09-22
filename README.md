Go-scrape
=========

A simple HTML scraper.

###How to use
```./scrape (config) (url) (output file)``` Example: ```./scrape configs/ahref.toml http://golang.org/doc/ golinks ```


The configs are written in toml. They're very simple and straight forward. Tag is the HTML tag, such as ```<a>``` (but you'd only put in ``a``) and AttributeName is the atribute you want out of the tag.

After running the program, you should end up with the output file you gave it in the arguement. Then you'd run something like ```wget -i outputfile``` to download all of them. 

#####How the config files work
They're faily simple like said above, they're toml
```
Tag = "img" #This is the name of the html tag, <img>, so you'll just put in img
AttributeName = "src" #This is the name of the atribute you want to scrape out of the tag you specifed above
IsURL = true # This is done so that if you're scraping a URL it'll attempt to fix them up so you can download them
```
With ``IsURL``:  /images/image.gif gets replaced to http://example.com/images/image.gif 
But if you're dealing with say the width of images, you don't want something like http://exmaple.com/64px so you'd set it to false.


###Todo
* Make a script to go through each url (the output doesn't always have to be a url as well) in the output file and complete the task again (we need to go deeper!) 
* Come up with more "configs" in examples.

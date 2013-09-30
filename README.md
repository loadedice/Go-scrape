Go-scrape
=========

A simple HTML scraper that can be used to scrape html attributes out of html and other things.

##How to use
```./scrape (config) (url) ``` Example: ```./scrape configs/ahref.toml http://golang.org/doc/  ```


Say you wanted to download all the URLs you just scraped, you'd cat the output into a file, lets call that file ouputfile, then you'd run something like ```wget -i outputfile``` to download all of them. 

#####How the config files work
The configs are written in toml. They're very simple and straight forward. It comes with some examples, here is an commented version of the imgsrc example included.
```Toml
Tag = "img" #This is the name of the html tag, <img>, so you'll just put in img
AttributeName = "src" #This is the name of the attribute you want to scrape out of the tag you specified above
IsURL = true # This is done so that if you're scraping a URL it'll attempt to fix them up so you can download them
```
With ``IsURL``:  /images/image.gif gets replaced to http://example.com/images/image.gif 
But if you're dealing with say the width of images, you don't want something like http://exmaple.com/64 so you'd set it to false.
#####Using go-deeper.py or go-deeper.sh
(for the bellow, go-deeper.py should be the same as go-deeper.sh, go-deeper.py is deprecated) 
```./go-deeper.py (config) (file with urls) ``` Example of usage in conjunction with scrape ```./scrape configs/ahref.toml https://github.com/ | sort -u > test && ./go-deeper.py configs/imgsrc.toml test ```
That command above will scrape content as per ahref.toml (it'll scrape the links) and then it'll remove any duplicates and then put it into a file called test. Then it'll run go-deeper.py and then for each line in test it'll run ./scrape again with the specified config (this time it'll get the images out of each of those links). the file ``test`` in this case is the file with all the urls that need to be scraped. This is useful for batch jobs as well.

Also go-deeper.py is tested in python 3, so it might not work on older versions of python. 

###Tips & tricks
* After running scrape, pipe it into ``sort -u`` to remove any duplicate lines.
* After you've got your list of URL's to download, run ``wget -i download`` where download is the file with all the urls
* If you want to download all the links to images on a site you can run something like ```./scrape configs/ahref.toml  (URL) | sort -u | grep -iE '.*?\.(jp?g|png|gif|bmp)' > download && wget -i download```

##Dependancies
If you want to build it on your machine, you'll need to install go, and these dependancies. They're easily installed with the following commands.
```
go get code.google.com/p/go.net/html
go get github.com/BurntSushi/toml
```
And to run the program it's as simple as ``go run scrape.go`` or if you want to make a binary, ``go build scrape.go`` and then run it like you would any other program.

##Todo
* Add in some more tips to the tips section
* Test for bugs

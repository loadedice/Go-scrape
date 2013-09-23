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
AttributeName = "src" #This is the name of the attribute you want to scrape out of the tag you specified above
IsURL = true # This is done so that if you're scraping a URL it'll attempt to fix them up so you can download them
```
With ``IsURL``:  /images/image.gif gets replaced to http://example.com/images/image.gif 
But if you're dealing with say the width of images, you don't want something like http://exmaple.com/64 so you'd set it to false.
####Using go-deeper.py
```./go-deeper.py (config) (file with urls) ``` Example of usage in conjunction with scrape ```./scrape configs/ahref.toml https://github.com/ | sort -u > test && ./go-deeper.py configs/imgsrc.toml test ```
That command above will scrape content as per ahref.toml (it'll scrape the links) and then it'll remove any duplicates and then put it into a file called test. Then it'll run go-deeper.py and then for each line in test it'll run ./scrape again with the specified config (this time it'll get the images out of each of those links). the file ``test`` in this case is the file with all the urls that need to be scraped. This is useful for batch jobs as well.

Also go-deeper.py is tested in python 3, so it might not work on older versions of python.

###Tips & tricks
* After running scrape, pipe it into ``sort -u`` to remove any duplicate lines.
* After you've got your list of URL's to download, run ``wget -i download`` where download is the file with all the urls

###Todo
* Add in some more tips to the tips section
* Test for bugs
* See if it's easier to have go-deeper.py as a go file or maybe just a shell script.
* Test the binaries portability and give proper build instructions.

#! /usr/bin/env python3
import sys
import os
import time

with open(sys.argv[2], "rt") as file:
    for line in file:
        os.system("./scrape"+sys.argv[1]+" "+line)
        time.sleep(1) #sleep between so you can quit and so it doesn't spam the site with requests. You can delete this line if you know the site won't mind, but beware you'll need to hold down control c for each of the go programs because of the way that they run

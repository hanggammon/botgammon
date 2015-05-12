# botgammon
Backgammon web service for AI bots

Stuff to set up:

* Heroku: https://toolbelt.heroku.com
* Godep: https://github.com/tools/godep
* GOPATH environment variable: https://golang.org/doc/code.html 

Stuff to read:
* https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html

Deploy:
* (already done?) heroku create -b https://github.com/heroku/heroku-buildpack-go.git
* heroku git:remote -a botgammon 
* git push heroku master

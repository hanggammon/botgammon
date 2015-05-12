# botgammon
Backgammon web service for AI bots

Stuff to set up:

* Heroku: https://toolbelt.heroku.com
* Godep: https://github.com/tools/godep
* GOPATH environment variable: https://golang.org/doc/code.html 

Stuff to read:
* https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html

Deploy:
* heroku login
* heroku git:remote -a botgammon

Correct git remote setup:

`$ git remote -v
heroku	https://git.heroku.com/botgammon.git (fetch)
heroku	https://git.heroku.com/botgammon.git (push)
origin	https://github.com/hanggammon/botgammon.git (fetch)
origin	https://github.com/hanggammon/botgammon.git (push)`

* git push heroku master


Initial heroku app create (already done)
* heroku create -b https://github.com/heroku/heroku-buildpack-go.git

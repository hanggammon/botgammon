# botgammon
Backgammon web service for AI bots

Stuff to set up:

* Heroku: https://toolbelt.heroku.com
* Godep: https://github.com/tools/godep
* GOPATH environment variable: https://golang.org/doc/code.html 

Stuff to read:
* https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html

Initial Setup:
* mkdir $HOME/go
* export GOPATH=$HOME/go
* export PATH=$PATH:$GOPATH/bin
* mkdir $HOME/go/src/
* cd $HOME/go/src/ && git clone --botgammon location--
* cd $HOME/go/src/botgammon/
* go get github.com/tools/godep
* godep restore
* heroku login
* heroku git:remote -a botgammon

Correct git remote setup:

<pre>
$ git remote -v
heroku	https://git.heroku.com/botgammon.git (fetch)
heroku	https://git.heroku.com/botgammon.git (push)
origin	https://github.com/hanggammon/botgammon.git (fetch)
origin	https://github.com/hanggammon/botgammon.git (push)
</pre>

Local Compile / Test:
* go get
* PORT=5000 botgammon
* Use browser to visit http://localhost:5000

Normal Deploy:
* git push heroku master

Force Deploy:
* git push heroku +master

Typical Directory Layout Of In-Progress Dev Environment:
<pre>
$HOME/go/bin/botgammon
$HOME/go/bin/godep
$HOME/go/pkg/darwin_amd64/...
$HOME/go/src/botgammon
$HOME/go/src/github.com/...
$HOME/go/src/golang.org/...
</pre>

Initial heroku app creation (already done, should no longer be needed):
* heroku create -b https://github.com/heroku/heroku-buildpack-go.git

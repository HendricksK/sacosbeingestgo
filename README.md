"# sacosbeingestgo" 

# go get all 
# SET PORT=6666 or whatever port you would like


# go modules setting to use local instead for testing https://golang.org/doc/tutorial/call-module-code
# form bindings https://echo.labstack.com/guide/binding/

Better golang code runner 

# https://github.com/canthefason/go-watcher going to give this a shot going forward

call watcher to run watcher 

need to better understand the documentation from github on this,
but essentially we will smb share with Oracle drive, upload images from the site to that smb share
run a cron to get all new images, drop those on a queue
run the queue one by one to upload images to github, under the untagged section
we then need to build a tagging interface, to add meta data etc etc

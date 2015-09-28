# dopepope :fire::fire::fire::fire::fire::fire::fire:
the pope wants to drop a :fire: mixtape but doesn't have the time here's where I come in.

# Requirements

1. Go version 1.5 or greater - you can get this from [downloads page] (https://golang.org/dl/) on the Golang website
2. ffmpeg 2.5.2 `brew install ffmpeg` if you have OSX
3. mongoDB v2.4.9 or greater...probably
4. python versoin 2.7 
5. [gTTs] (https://github.com/faiq/dopepope) `pip install gTTS`

# How to Drop the :fire: tape
This application takes a bunch of speeches that the pope made, sanitized by using sed, populates them into mongoDB in an object which takes the last word of each sentence and the sentence itself.

If you want to do this (you only need to do this once)

1. run `mongod`
2. (hack) change the package name of `populate/text-to-mongo.go` from `populate` to `main`
3. from the project root `go run populate/text-to-mongo.go`

Once you have your database (this application uses DB name "dopepope" and collection name "sentencestest") all set up with the popes words you can run the `make-fire.go` file - which makes the :fire:.

To do this

1. run `mongod` if you closed it  
2. change package name of `populate/text-to-mongo.go` from `main` back to `populate` 
3. `go run make-fire.go`

make-fire takes in a -fire flag which gives the pope a topic to start dropping :fire: on. It will create an `output.txt` file where you can view all the words that rhyme with that topic

If you want to put some :fire: over a beat run (we use in da club as an example here)

1. run `gtts-cli.py -f output.txt -l 'en' output.mp3`
2. run `fmpeg -i output.mp3 -i indaclub.mp3 -filter_complex amerge -c:a libmp3lame -q:a 4 ~/Desktop/dopepope.mp3`

Its a long serious of steps, but if everything worked right you should have a :fire: song called `dopepope.mp3` on your desktop!

An example of one can be found in this directory called `dopepope.mp3`

Open a new issue or pull request if you want to help! 

# LICENSE

The Dope Pope License (DP)

Copyright (c) The Pope Himself 2015

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Scripture"), to deal
in the Scripture without restriction, including without limitation the rights
to pray with, use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Scripture, and to permit persons to whom the Scripture is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Scripture.

THE SCRIPTURE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
POPE OR CARDINALS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
LIABILITY IN CONNECTION WITH THE SCRIPTURE OR THE USE OR OTHER
DEALINGS IN THE SCRIPTURE.

:fire:

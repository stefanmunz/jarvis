# Jarvis.io

Welcome to Jarvis.io. This is our project at the 2014 Protonet Hackathon in Hamburg. It is a smart assistant that you can talk to!

# What you will need
 
 - a Raspberry Pi. To setup the Pi it is helpful to have an usb keyboard and a usb mouse, also a WiFi-Dongle is nice
 - a microfone: we used a good one from Zoom

# Setup

 - Setup an operating system on your Pi, we used NOOBS and followed the install instructions at (http://jankarres.de/2014/07/raspberry-pi-mit-noobs-einfach-installieren/)
 - it makes sense to enable SSH on the Pi, also enabling bonjour proved to be helpful
 - There was no setup involved when connecting the microfone. It just worked. Your mileage may vary
 - install go on the raspberry, we followed this instruction: http://dave.cheney.net/2012/09/25/installing-go-on-the-raspberry-pi
 - install voicecommand on the Pi, just follow the instructions: https://github.com/StevenHickson/PiAUISuite
 - voicecommand has config file, it lies in the home directory of your user on the pi, in `.commands.conf`. The config we used is part of this repository
 - clone the repository to your raspberry and do a `go get` and `go install`
 - then run `voicecommand -c`
 
 

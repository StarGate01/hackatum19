#!/usr/bin/python3

import sys, os, time, requests
from signal import *
from os import listdir
from os.path import isfile, join

imgs = []
cidx = 0

def trigger(img):
    print("Taking and sending image: " + img)
    imgfile = {"image": open("/data/mock/" + img, 'rb')}
    res = requests.post("http://core:3000/images", files=imgfile)
    print("Done sending image: " + str(res.status_code) + " " + res.text)

def main():
    global cidx, imgs
    imgs = [f for f in listdir("/data/mock") if isfile(join("/data/mock", f))]
    print(str(len(imgs)) + " images found")
    print("Camera mock started\n")
    cidx = 0
    while(True):
        trigger(imgs[cidx])
        cidx += 1
        if(cidx >= len(imgs)):
            cidx = 0
        time.sleep(20)

def sigusr(sig, frame):
    global cidx, imgs
    print("SIGUSR: cidx=" + str(cidx))
    trigger(imgs[cidx])

if __name__ == '__main__':
    signal(SIGUSR1, sigusr)
    main()
    
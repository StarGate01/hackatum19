#!/usr/bin/python3

import sys, os
from signal import *

def trigger():
    print("Taking and sending picture: " + img)

def main():
    print("Camera mock started")


def sigusr(sig, frame):
    trigger()

if __name__ == '__main__':
    signal(SIGUSR, core.siginth)
    if(len(sys.argv) > 1 and sys.argv[1] == "trigger"):
        trigger()
    main()
    
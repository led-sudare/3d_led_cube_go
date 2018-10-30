#!/bin/sh

cd /home/pi/go/src/3d_led_cube_go/
exec ./3d_led_cube_go -r 192.168.0.25:5501 -s show:{\"orders\":[{\"id\":\"object-realsense\",\"lifetime\":0}]}


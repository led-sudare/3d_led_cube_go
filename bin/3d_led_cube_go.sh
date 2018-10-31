#!/bin/sh

cd /home/pi/go/src/3d_led_cube_go/
exec ./3d_led_cube_go -d 192.168.0.10:9001 -r 192.168.0.20:5501 -s \{\"orders\":\[\{\"id\":\"object-realsense\",\"lifetime\":0\}\]\}

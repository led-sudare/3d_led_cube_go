#!/bin/bash -eu
cp bin/3d_led_cube_go.service /etc/systemd/system/.

INSTALL_DIR=/opt/3d_led_cube_go/bin 

mkdir -p $INSTALL_DIR
chmod 755 $INSTALL_DIR
cp bin/3d_led_cube_go.sh $INSTALL_DIR/.
chown root:root ${INSTALL_DIR}/3d_led_cube_go.sh
chmod 755 ${INSTALL_DIR}/3d_led_cube_go.sh



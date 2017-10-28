#!/bin/bash
cd /home/bandit24
touch pass
chmod 644 /home/bandit24/pass
cat /etc/bandit_pass/bandit24 &> /home/bandit24/pass

#!/bin/bash
sudo rm /app -rf
sudo mkdir /app
sudo chown -R ec2-user:ec2-user /app

sudo yum update -y
sudo yum install -y golang
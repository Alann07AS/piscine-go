#!/bin/bash
curl https://zone01normandie.org/assets/superhero/all.json | jq --arg city "$HERO_ID" ' .[] | select(.id == $HERO_ID)  | .connections | .relatives  ' -r
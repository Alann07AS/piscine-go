curl https://zone01normandie.org/assets/superhero/all.json | jq --arg HERO_ID "$HERO_ID" ' .[] | select(.id == contains('$HERO_ID'))  | .connections | .relatives  ' -r
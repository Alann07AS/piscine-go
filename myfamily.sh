import HERO_ID
curl https://zone01normandie.org/assets/superhero/all.json | jq ' .[] | select(.id == HERO_ID)  | .connections | .relatives  ' -r
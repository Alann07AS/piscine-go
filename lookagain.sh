find . -name '*.sh' | sed 's/.sh//g'  |sort -r |  cut -f 2 --delimiter="/"
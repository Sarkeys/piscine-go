find . \( -type d,f \) -name "*.sh" | cut -d "." -f 2| sort -r | cut --complement -d "/" -f 1 | cut --complement -d "/" -f 1 | cut --complement -d "/" -f 1 | cut --complement -d "/" -f 1 



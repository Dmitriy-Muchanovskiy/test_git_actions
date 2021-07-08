#/bin/bash

color=""
cat $COVER_TOTAL_FILE
#clear_cover=$(echo $cover | awk -F . '{print $1}')
clear_cover=$(cat $COVER_TOTAL_FILE | awk -F . '{print $1}')
echo $clear_cover
if [[ $clear_cover -gt 94 ]]
then
    color=brightgreen
elif [[ $clear_cover -gt 80 ]]
then
    color=green
elif [[ $clear_cover -gt 60 ]]
then
    color=yellow
elif [[ $clear_cover -gt 40 ]]
then
    color=orange
elif [[ $clear_cover -gt 20 ]]
then
    color=red 
elif [[ $clear_cover -ge 0 ]]
then
    color=lightgrey
fi
echo -e "$color \n 0-20 = lightgrey, 21-40 = red, 41-60 = orange, 61-80, = yellow, 81-95 = green, 95-100 brightgreen"
#echo $color
echo "::set-output name=cover_color::$color"
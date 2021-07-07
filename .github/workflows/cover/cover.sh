#/bin/bash

#cover=80.8%
cover=${{ needs.go_test.outputs.output1 }}
color=""
clear_cover=$(echo $cover | awk -F . '{print $1}')
echo $clear_cover
case $clear_cover in
    [1-9] | [1][0-9] | 20 ) color=lightgrey ;;
    [2][1-9] | [3][0-9] | 40 ) color=red ;;
    [4][1-9] | [5][0-9] | 60 ) color=orange ;;
    [6][1-9] | [7][0-9] | 80 ) color=yellow ;;
    [8][1-9] | [9][0-4] ) color=green ;;
    [9][5-9] | 100 ) color=brightgreen ;;
esac
echo "0-20 = lightgrey, 21-40 = red, 41-60 = orange, 61-80, = yellow, 81-95 = green, 95-100 brightgreen"
echo $color
echo "::set-output name=cover_color::$color"
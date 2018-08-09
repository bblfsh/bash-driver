#!/bin/bash

case $casevar in 
    'somestr') $action1='value1'
        ;;
    $othervar) $action2='value2'; $action3=4
        ;;
    [1-6]) echo 'foo'
        ;;
    100) echo 'bar'
        ;;
    *)  echo 'default'
        ;;
esac

#!/bin/sh

react-scripts start & 
npx tailwindcss -i ./src/App.css -o ./dist/output.css --watch 


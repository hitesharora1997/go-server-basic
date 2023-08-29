package main

import (
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

var rnd *renderer.Render
var db *mgo.Database // mongo db

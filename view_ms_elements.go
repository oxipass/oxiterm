package main

import "github.com/rivo/tview"

// Main vertical flex which is wrapping all othe flexes
var wrapperFlex *tview.Flex

// Main horizontal flex panels: items list and fields list
var itemsList *tview.List
var fieldsList *tview.List

// Horizontal buttons flex
var buttonAddItem *tview.Button
var buttonAddField *tview.Button
var buttonLock *tview.Button
var buttonQuit *tview.Button
var buttonAbout *tview.Button
var buttonSettings *tview.Button

// Tags filter and search input
var searchInput *tview.InputField
var tagsFilter *tview.DropDown

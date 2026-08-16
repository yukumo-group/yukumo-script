---
title: "Characters"
permalink: /script/characters
---

Character system is one of the most important concepts in this program. By defining characters, the user can get rid of the need for defining the phont file again and again in the script file. The following content will introduce the basic concepts in the characters system. 

## Characters

The definition of character struct is: 
```go
type Character struct {
	Name             string  `json:"name"`
	PhontName        string  `json:"phontName"`
	Description      string  `json:"description"`
	ProfileImagePath *string `json:"profileImagePath"`
}
```
In here, `Name` refers to the name of the character, which is also the ID of character. PhontName refers to the name of the phont of the character. Description describes the character and 

## Use Pre-defined Characters

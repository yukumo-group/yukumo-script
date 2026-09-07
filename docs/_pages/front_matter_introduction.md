---
title: "Front Matter Introduction"
permalink: /script/frontMatter/introduction
---

This part contains basic settings for the script, for instance the **default speed** for the generated audio and characters. The users are also allowed to define their own characters in the application and use them in the script.

This part follows the YAML file style so you can just provide the script file with body matter and the front matter YAML file to the application. 

Here is an example template of the front matter yaml file:

```yaml
# use_predefined_characters decide whether to use the predefined characters stored in the application. 
# It is recommended not to set it if you want the script to be used by others
# If this is set to false, the characters defined in the raw_characters will not be loaded and you can left it empty. 
# If this is set to true, the raw_characters cannot be empty. Otherwise the program will throws an error. 
use_predefined_characters: false
# config_name defines the name of this configuration. 
# If you write your config file from yaml file or script file, you can left it empty as config_name will be set to the name of the file.
config_name: test1
# raw_characters contains definition of characters. 
# It is a sequence. 
# If you set use_predefined_characters to true, this will not be loaded. Otherwise you can only use the characters defined here. 
raw_characters:
- character_name: Remilia Scarlet # character_name refers to the name that identifies a unique character. 
  phont_name: aq_f1c # phont_name refers to the name of the phont file, no .phont suffix needed
  description: Remilia Scarlet (レミリア・スカーレット Remiria Sukāretto) is a vampire who is the head of the Scarlet Devil Mansion. She is the sister of Flandre Scarlet, the mistress of Hong Meiling and Sakuya Izayoi as well as the other fairy maids and the friend of Patchouli Knowledge. She first appears as the Stage 6 boss and main antagonist of the Embodiment of Scarlet Devil and has been a playable character in multiple games since Imperishable Night, in a team with Sakuya, and Immaterial and Missing Power as well.
- character_name: Sakuya Izayoi
  phont_name: aq_f1b
  description: Easily one of the most enigmatic characters in the series despite appearing in so many games, Sakuya Izayoi is the Chief Maid at the Scarlet Devil Mansion. She works for her mistress, Remilia Scarlet, and is apparently the only human working and living within the mansion. She has the power to manipulate time, and is known to place knives in midair and resume time to allow these knives to fly towards her targets. Stopping time is also a handy way of doing large amounts of maid work in a short time. Because nearly everyone living and working at the mansion are maids, being the chief of them all means there are almost no people with more authority within the mansion than Sakuya.
```

---
title: "Introduction"
permalink: /script/introduction
---

Yukumo-Script is a scripting language for generating yukumo audios or subtitle files. It is intended to be designed like drama scripts to make it easier to use. 

Each script file is devided into two different parts, the **front matter** and the **body matter**. Their syntax is not the same. 

# Front Matter

Front matter contains basic settings for the script, for instance the **default speed** for the generated audio and characters. The users are also allowed to define their own characters in the application and use them in the script.

# Body Matter

Body matter contains the texts of the script, including the waiting part. The basic unit for the body matter part is **sentence**. One valid sentence must follow the basic structure of dialogue in drama script (*Character: Word*) or be a empty waiting sentence. Although the user can alter the speed of the whole sentence, the speed inside a sentence must be consistent. If you want to change the speed during a continuous speech of a character, please devide this sentence into multiple sentences. 
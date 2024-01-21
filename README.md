# gogrep
This is an attempt to recreate a part of ripgrep using go
## Disclaimer
⚠️ I'm making gogrep for learning purposes, both fun and challenging. It's still on a really early stage but we will see what it will become ⚠️

### Why not just use ripgrep
Well there's only one reason to use gogrep over ripgrep (I use ripgrep).
With gogrep you understand the how ripgrep/grep basic functionality works

### Features
- 🚀 Blazingly Fast
- 🛠️ Easy Understand and use
- 📦 Maintainable Code (Small Code Base)
- 🎨 Randomly assigns colors to your results

### Usage
Look for the input in the file content inside a directory
![image](https://github.com/MohamedBenMassouda/gogrep/assets/55658633/13cf13ac-66fd-4466-b206-13d38713c207) <br>
- <b>go run .</b>: is used to run the go code in the current directory
- <b>"go"</b>: is the input string to look for
- <b>.</b>: indicating the current directory

### Installation
```
git clone https://github.com/MohamedBenMassouda/gogrep && cd gogrep
```
```
go run . "🚀" .
```

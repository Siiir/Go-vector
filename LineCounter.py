import os

# os.chdir(r"G:\My Drive\Open to public\Around learning\IT\full-fledged\Go\modules\vectors")
cou= 0
for fName in os.listdir():
    if os.path.isfile(fName):
        try:
            if not fName.startswith(".git") and not fName.startswith("go."):
                for line in open(fName):
                    cou+=1
        except UnicodeDecodeError: pass
print(cou)